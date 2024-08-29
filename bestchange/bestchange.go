package bestchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
	"time"
)

type Client struct {
	token string
	cache map[string]Response
}

type Response struct {
	Rates []byte
	Time  time.Time
}

// Create bestchange client.
func New(token string) *Client {
	return &Client{token: token}
}

// Exchanger with specific rate.
type Rate struct {
	Rate    string   `json:"rate"`
	RateRev string   `json:"raterev"`
	Inmin   string   `json:"inmin"`
	Inmax   string   `json:"inmax"`
	Reserve string   `json:"reserve"`
	Marks   []string `json:"marks"`
	Changer int      `json:"changer"`
}

// Receive exchanger rates from bestchange for 2 currency ID's.
func (c *Client) Rates(first, second uint16) ([]Rate, error) {
	var body Response

	url := fmt.Sprintf("https://www.bestchange.app/v2/%s/rates/%d-%d", c.token, first, second)

	value, ok := c.cache[url]
	if ok && time.Since(value.Time) < time.Minute {
		body = value
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to create request: %v", err)
		}

		req.Header.Add("User-Agent", "Thunder Client (https://www.thunderclient.com)")
		req.Header.Add("accept", "application/json")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, fmt.Errorf("unable to execute request: %v", err)
		}

		defer res.Body.Close()
		cacheBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("unable to read response: %v", err)
		}
		c.cache[url] = Response{
			Rates: cacheBody,
			Time:  time.Now(),
		}
	}

	wrapmap := map[string]map[string][]Rate{}
	err := json.Unmarshal(body.Rates, &wrapmap)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal request 1: %v", err)
	}
	rates := wrapmap["rates"][fmt.Sprintf("%d-%d", first, second)]

	sort.Slice(rates, func(i, j int) bool {
		return rates[i].Rate < rates[j].Rate
	})
	for i := range rates {
		flt, _ := strconv.ParseFloat(rates[i].Rate, 8)
		rates[i].RateRev = fmt.Sprintf("%f", 1/flt)
	}

	return rates, err
}

// Give a table representation for received views, can be used for debugging.
func PrintTable(r []Rate) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, "Rate\tRateRev\tInmin\tInmax\tReserve\tId\t")
	for _, rate := range r {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\t\n", rate.Rate, rate.RateRev, rate.Inmin, rate.Inmax, rate.Reserve, rate.Changer)
	}
	w.Flush()
}

// Returns good forward selling rate, good backward selling rate and average
func (c *Client) EstimateRates(forward, backward []Rate) (float64, float64, float64) {
	var forwardRate, backwardRate, average float64
	for i := range forward {
		if i == 3 {
			break
		}
		flt, err := strconv.ParseFloat(forward[i].Rate, 32)
		if err != nil {
			panic(err)
		}
		average += flt
		forwardRate += flt
	}
	for i := range backward {
		if i == 3 {
			break
		}
		flt, err := strconv.ParseFloat(backward[i].RateRev, 32)
		if err != nil {
			panic(err)
		}
		average += flt
		backwardRate += flt
	}
	return forwardRate / 3, backwardRate / 3, average / 6
}

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
	Resp []byte
	Time time.Time
}

// Create bestchange client.
func New(token string) *Client {
	return &Client{
		token: token,
		cache: map[string]Response{},
	}
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
func (c *Client) Rates(first, second string) ([]Rate, error) {
	var body Response

	url := fmt.Sprintf("https://www.bestchange.app/v2/%s/rates/%s-%s", c.token, first, second)

	value, ok := c.cache[url]
	if ok && time.Since(value.Time) < time.Minute {
		body = value
	} else {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, fmt.Errorf("unable to create request: %v", err)
		}

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
			Resp: cacheBody,
			Time: time.Now(),
		}
		body = Response{Resp: cacheBody}
	}

	wrapmap := map[string]map[string][]Rate{}
	err := json.Unmarshal(body.Resp, &wrapmap)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal response: %v", err)
	}
	rates := wrapmap["rates"][fmt.Sprintf("%s-%s", first, second)]

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
func (c *Client) PrintTable(r []Rate) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, "Rate\tRateRev\tInmin\tInmax\tReserve\tId\t")
	for _, rate := range r {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\t\n", rate.Rate, rate.RateRev, rate.Inmin, rate.Inmax, rate.Reserve, rate.Changer)
	}
	w.Flush()
}

type Currencies struct {
	Currencies []Currency `json:"currencies"`
}

type Currency struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
}

// Create currency id to currency card mapper.
func (c *Client) CreateCurrencyMapper() (map[string]string, error) {
	var body Response

	url := fmt.Sprintf("https://www.bestchange.app/v2/%s/currencies/ru", c.token)

	value, ok := c.cache[url]
	if ok && time.Since(value.Time) < time.Minute {
		body = value
	} else {
		req, _ := http.NewRequest("GET", url, nil)

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
			Resp: cacheBody,
			Time: time.Now(),
		}
		body = Response{Resp: cacheBody}
	}

	var currencies Currencies
	err := json.Unmarshal(body.Resp, &currencies)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal response: %v", err)
	}

	result := map[string]string{}
	for _, curr := range currencies.Currencies {
		result[curr.Code] = fmt.Sprintf("%d", curr.Id)
	}
	return result, nil
}

// Function which takes rates and based on this rates estimates a good price
// for exchange operations.
func (c *Client) EstimateOver(r []Rate) (float64, float64) {
	var result float64
	var resultrev float64
	for i := range r {
		if i == 4 {
			break
		}
		flt, err := strconv.ParseFloat(r[i].Rate, 32)
		if err != nil {
			panic(err)
		}
		result += flt
		fltrev, err := strconv.ParseFloat(r[i].RateRev, 32)
		if err != nil {
			panic(err)
		}
		resultrev += fltrev
	}
	return result / 4, resultrev / 4
}

// Estimate exchane pair to get actual and reversed exchange rate for pair
func (c *Client) EstimateOperation(first, second string) (float64, error) {
	m, err := c.CreateCurrencyMapper()
	if err != nil {
		return 0, err
	}

	rates, err := c.Rates(m[first], m[second])
	if err != nil {
		return 0, err
	}
	rate, _ := c.EstimateOver(rates)
	return rate, nil
}

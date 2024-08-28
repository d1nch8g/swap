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
)

// Usage example:

// client := bestchange.NewClient("2dcd269e14d6bf5060e9df0fc7ab16a7")

// sbpton, err := client.Rates(21, 209) // SBP to TON
// if err != nil {
// 	panic(err)
// }
// fmt.Println(" ===== SBT-TON ===== Give SBT, receive TON")
// bestchange.PrintTable(sbpton)

// tonsbp, err := client.Rates(209, 21) // TON to SBP
// if err != nil {
// 	panic(err)
// }
// fmt.Println(" ===== TON-SBP ===== Give TON, receive SBT")
// bestchange.PrintTable(tonsbp)

// avg := bestchange.EstimateAverageRate(sbpton, tonsbp)
// fmt.Printf("Average estimated price is: %f\n", avg)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

type Rate struct {
	Rate    string   `json:"rate"`
	RateRev string   `json:"raterev"`
	Inmin   string   `json:"inmin"`
	Inmax   string   `json:"inmax"`
	Reserve string   `json:"reserve"`
	Marks   []string `json:"marks"`
	Changer int      `json:"changer"`
}

func (c *Client) Rates(first, second uint16) ([]Rate, error) {

	url := fmt.Sprintf("https://www.bestchange.app/v2/%s/rates/%d-%d", c.token, first, second)

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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}

	wrapmap := map[string]map[string][]Rate{}
	err = json.Unmarshal(body, &wrapmap)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal request 1: %v", err)
	}
	slice := wrapmap["rates"][fmt.Sprintf("%d-%d", first, second)]

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].Rate < slice[j].Rate
	})
	for i := range slice {
		flt, _ := strconv.ParseFloat(slice[i].Rate, 8)
		slice[i].RateRev = fmt.Sprintf("%f", 1/flt)
	}

	return slice, err
}

func PrintTable(r []Rate) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, "Rate\tRateRev\tInmin\tInmax\tReserve\tId\t")
	for _, rate := range r {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%d\t\n", rate.Rate, rate.RateRev, rate.Inmin, rate.Inmax, rate.Reserve, rate.Changer)
	}
	w.Flush()
}

func EstimateAverageRate(forward, backward []Rate) float64 {
	var result float64
	for i := range forward {
		if i == 3 {
			break
		}
		flt, _ := strconv.ParseFloat(forward[i].Rate, 8)
		result += flt
	}
	for i := range backward {
		if i == 3 {
			break
		}
		flt, _ := strconv.ParseFloat(backward[i].RateRev, 8)
		result += flt
	}
	return result / 6
}

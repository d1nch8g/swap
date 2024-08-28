package bestchange

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"text/tabwriter"
)

type Client struct {
	token string
}

func NewClient(token string) *Client {
	return &Client{token: token}
}

type Rate struct {
	Rate    string   `json:"rate"`
	Inmin   string   `json:"inmin"`
	Inmax   string   `json:"inmax"`
	Reserve string   `json:"reserve"`
	Marks   []string `json:"marks"`
	Changer int      `json:"changer"`
}

func (c *Client) Rates(first, second, limit uint16, reverse, order bool) ([]Rate, error) {

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

	if limit != 0 {
		slice = slice[:limit]
	}

	if order {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].Rate < slice[j].Rate
		})
	}

	return slice, err
}

func PrintTable(r []Rate) {
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	fmt.Fprintln(w, "Rate\tInmin\tInmax\tReserve\tId\t")
	for _, rate := range r {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%d\t\n", rate.Rate, rate.Inmin, rate.Inmax, rate.Reserve, rate.Changer)
	}
	w.Flush()

}

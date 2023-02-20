package binance

import (
	"encoding/json"
)

func (c *Client) GetOrderBooks(params []string) (interface{}, error) {
	req := Request{
		Path:   "/api/v3/depth",
		Method: "GET",
		Param: GetBooksParam{
			Symbol: params[0],
			Limit:  params[1],
		},
	}
	raw_data, date, err := c.do(req)
	data := &BooksData{
		Date: *date,
	}
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(raw_data, data); err != nil {
		return nil, err
	}

	return data, nil
}

type GetBooksParam struct {
	Symbol string `url:"symbol"`
	Limit  string `url:"limit,omitempty"`
}

type BooksData struct {
	Date         string     `json:"date"`
	LastUpdateID int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

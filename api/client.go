package binance

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

type Client struct {
	Host string
	C    *http.Client
}

// new *Client
func New(c *http.Client) *Client {
	return &Client{
		Host: "api.binance.com",
		C:    c,
	}
}

// Use package net/http
func (c *Client) do(r IRequest) ([]byte, *string, error) {
	var body []byte
	var date *string
	var err error
	req := c.newRequest(r)
	resp, err := c.C.Do(req)
	if err != nil { // ======================
		return nil, nil, err
	}
	defer resp.Body.Close()
	date = &resp.Header["date"][0]
	if len(resp.Header["date"]) > 1 {
		panic("len(resp.Header[\"date\"]) > 1")
	}
	if resp.StatusCode != http.StatusOK {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		return nil, date, fmt.Errorf("http status code:%d, desc:%s", resp.StatusCode, string(body))
	}
	body, _ = ioutil.ReadAll(resp.Body)
	return body, date, nil
}

// new Request
// for package net/http
func (c *Client) newRequest(r IRequest) *http.Request {
	var raw_body []byte
	var body io.Reader = nil
	path := r.GetPath()
	if r.IsPost() {
		raw_body, _ = json.Marshal(r.GetParam())
	} else if values, _ := query.Values(r.GetParam()); len(values) > 0 {
		path += "?" + values.Encode()
	}
	if string(raw_body) != "" {
		body = strings.NewReader(string(raw_body))
		// body = ioutil.NopCloser(strings.NewReader(sign.Body))
	}
	req, err := http.NewRequest(r.GetMethod(), c.Host+path, body)
	if err != nil {
		panic(err)
	}

	// to solve to EOF problem
	// Refer to https://stackoverflow.com/questions/17714494/golang-http-request-results-in-eof-errors-when-making-multiple-requests-successi
	req.Close = true

	return req
}

package binance

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestGetOrderBooks(t *testing.T) {
	proxyUrl, err := url.Parse("http://127.0.0.1:8443")
	if err != nil {
		panic(err)
	}
	httpClient := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   30 * time.Second,
	}
	client := New(httpClient)
	result, err := client.GetOrderBooks([]string{"BTCUSDT", "100"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", result.(*BooksData))
}

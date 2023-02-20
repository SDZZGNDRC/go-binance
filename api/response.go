package binance

type IResponse interface {
	GetCode() string
	GetMessage() string
	IsOk() bool
}

type Response struct {
	Date string      `json:"date"`
	Data interface{} `json:"data"`
}

type Response2 struct {
	Date string      `json:"date"`
	Data interface{} `json:"data"`
}

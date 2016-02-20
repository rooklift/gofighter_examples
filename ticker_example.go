package main

import (
	"fmt"
	"encoding/json"

	"github.com/fohristiwhirl/gofighter"
)

var INFO gofighter.TradingInfo = gofighter.TradingInfo{
	BaseURL: "https://api.stockfighter.io/ob/api",
	WebSocketURL: "wss://api.stockfighter.io/ob/api/ws",
	Account: "EXB123456",
	Venue: "TESTEX",
	Symbol: "FOOBAR",
}

func main() {

	ticker_results := make(chan gofighter.Quote, 64)
	go gofighter.Ticker(INFO, ticker_results)

	for n := 0 ; ; n++ {
		msg := <- ticker_results
		fmt.Printf("%d ", n)
		s, _ := json.MarshalIndent(msg, "", "  ")
		fmt.Println(string(s))
	}

	return
}

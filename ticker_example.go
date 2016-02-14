package main

import (
	"fmt"
	"encoding/json"

	"github.com/fohristiwhirl/gofighter"
)

const (
	BASE_URL = "https://api.stockfighter.io/ob/api"
	WS_URL   = "wss://api.stockfighter.io/ob/api/ws"
	// BASE_URL = "http://127.0.0.1:8000/ob/api"
	// WS_URL   = "ws://127.0.0.1:8000/ob/api/ws"

	ACCOUNT = "EXB123456"
	VENUE = "TESTEX"
	SYMBOL = "FOOBAR"
)

func main() {

	ticker_results := make(chan gofighter.Quote, 64)
	go gofighter.Ticker(WS_URL, ACCOUNT, VENUE, SYMBOL, ticker_results)

	for n := 0 ; ; n++ {
		msg := <- ticker_results
		fmt.Printf("%d ", n)
		s, _ := json.MarshalIndent(msg, "", "  ")
		fmt.Println(string(s))
	}

	return
}

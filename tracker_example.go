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

	tracker_results := make(chan gofighter.Execution, 64)
	go gofighter.Tracker(WS_URL, ACCOUNT, VENUE, SYMBOL, tracker_results)

	for n := 0 ; ; n++ {
		msg := <- tracker_results
		fmt.Printf("%d ", n)
		s, _ := json.MarshalIndent(msg, "", "  ")
		fmt.Println(string(s))
	}

	return
}

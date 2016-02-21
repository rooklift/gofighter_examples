package main

// This example opens an executions websocket and prints events from it...

import (
	"fmt"

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

	ex_chan := make(chan gofighter.Execution)

	go gofighter.Tracker(INFO, ex_chan)

	for n := 0 ; ; n++ {
		msg := <- ex_chan
		fmt.Printf("%s %s %d @ %d\n", msg.Account, msg.Order.Direction, msg.Filled, msg.Price)
	}

	return
}

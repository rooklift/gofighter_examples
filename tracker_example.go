package main

// This example opens an executions websocket and prints the current position any time it changes
// (note that on the test exchange, there is only one account and the net position always becomes 0)

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

	position_channel := make(chan gofighter.Position, 64)
	go gofighter.PositionUpdater(INFO, nil, nil, position_channel)

	for n := 0 ; ; n++ {
		newpos := <- position_channel
		fmt.Printf("%d: Shares: %d, Cents: %d\n", n, newpos.Shares, newpos.Cents)
	}

	return
}

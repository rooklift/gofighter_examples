package main

// This example opens an executions websocket and prints the current position any time it changes
// (note that on the test exchange, there is only one account and the net position always becomes 0)

import (
	"fmt"

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

	position_channel := make(chan gofighter.Position, 64)
	go gofighter.PositionUpdater(WS_URL, ACCOUNT, VENUE, SYMBOL, nil, nil, position_channel)

	for n := 0 ; ; n++ {
		newpos := <- position_channel
		fmt.Printf("%d: Shares: %d, Cents: %d\n", n, newpos.Shares, newpos.Cents)
	}

	return
}

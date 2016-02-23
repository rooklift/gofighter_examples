package main

// This example opens an executions websocket and prints events from it...

import (
	"fmt"

	"github.com/fohristiwhirl/gofighter"
)

func main() {

	// The following assumes game_start.exe has already been called, and thus that
	// there exists a file in the /gm/ directory which contains info about the level.

    info := gofighter.GetUserSelection("known_levels.json")

	// The Ticker goroutine called below is the direct interface to the server WebSocket.
	// It simply sends executions down a channel at us.

	ex_chan := make(chan gofighter.Execution)
	go gofighter.Tracker(info, ex_chan)

	for n := 0 ; ; n++ {
		msg := <- ex_chan
		fmt.Printf("%s %s %d @ %d\n", msg.Account, msg.Order.Direction, msg.Filled, msg.Price)
	}

	return
}

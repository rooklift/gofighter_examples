package main

// This example opens an executions websocket and prints events from it...

import (
	"fmt"

	"github.com/fohristiwhirl/gofighter"
)

func main() {

    info := gofighter.GetUserSelection("known_levels.json")

	ex_chan := make(chan gofighter.Execution)

	go gofighter.Tracker(info, ex_chan)

	for n := 0 ; ; n++ {
		msg := <- ex_chan
		fmt.Printf("%s %s %d @ %d\n", msg.Account, msg.Order.Direction, msg.Filled, msg.Price)
	}

	return
}

package main

import (
	"fmt"
	"encoding/json"

	"github.com/fohristiwhirl/gofighter"
)

func main() {

	// The following assumes game_start.exe has already been called, and thus that
	// there exists a file in the /gm/ directory which contains info about the level.

	info := gofighter.GetUserSelection("known_levels.json")

	// The Ticker goroutine called below is the direct interface to the server WebSocket.
	// It simply sends quotes down a channel at us.

	ticker_results := make(chan gofighter.Quote, 64)
	go gofighter.Ticker(info, ticker_results)

	for n := 0 ; ; n++ {
		msg := <- ticker_results
		fmt.Printf("%d ", n)
		s, _ := json.MarshalIndent(msg, "", "  ")
		fmt.Println(string(s))
	}

	return
}

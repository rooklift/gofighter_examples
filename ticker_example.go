package main

import (
	"fmt"
	"encoding/json"

	"github.com/fohristiwhirl/gofighter"
)

func main() {

	info := gofighter.GetUserSelection("known_levels.json")

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

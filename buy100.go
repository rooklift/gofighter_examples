package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {

	// Assuming that game_start.exe has been run, this is all we need...

	levelname := gofighter.NameFromUser()
	info := gofighter.TradingInfoFromName(levelname)

	order := gofighter.ShortOrder{
		Direction: "buy",
		OrderType: "limit",
		Qty: 100,
		Price: 10000,
	}

	result, _ := gofighter.Execute(info, order, nil)
	gofighter.PrintJSON(result)
}

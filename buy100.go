package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {

	// The following assumes game_start.exe has already been called, and thus that
	// there exists a file in the /gm/ directory which contains info about the level.

	info := gofighter.GetUserSelection("known_levels.json")

	// The info var above now contains account, venue, and symbol. So we just need
	// the following pieces of information; the ShortOrder type is designed for this.

	order := gofighter.ShortOrder{
		Direction: "buy",
		OrderType: "limit",
		Qty: 100,
		Price: 10000,
	}

	result, _ := gofighter.Execute(info, order, nil)
	gofighter.PrintJSON(result)
}

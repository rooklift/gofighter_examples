package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {

	info := gofighter.GetUserSelection("known_levels.json")

	order := gofighter.ShortOrder{
		Direction: "buy",
		OrderType: "limit",
		Qty: 100,
		Price: 10000,
	}

	result, _ := gofighter.Execute(info, order, nil)
	gofighter.PrintJSON(result)
}

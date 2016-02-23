package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {

	levelname := gofighter.NameFromUser()

	start, _ := gofighter.GMstart("", levelname)

	// The response is saved in the gm/ folder so that we can load the essential info
	// (account, venue, symbol) in the trading program, or whatever.

	gofighter.SaveGMfile(levelname, start)

	info := gofighter.TradingInfoFromName(levelname)
	gofighter.PrintJSON(info)
}

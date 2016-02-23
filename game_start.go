package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {

	// First we use the list of known levels to get a level name from the user.
	// Then we call the gamemaster to start that level.

	levelname := gofighter.NameFromUser("known_levels.json")
	start, _ := gofighter.GMstart("", levelname)

	// The response is saved in the /gm/ folder so that we can load the essential
	// info (account, venue, symbol) in the trading program, or whatever.

	gofighter.SaveGMfile(levelname, start)

	info := gofighter.TradingInfoFromName(levelname)
	gofighter.PrintJSON(info)
}

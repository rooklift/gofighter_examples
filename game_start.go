package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {

	// The following lines could be at the start of a real
	// program to get the info required to play the level...

	levelname := gofighter.NameFromUser()

	start, _ := gofighter.GMstart("", levelname)
	gofighter.SaveGMfile(levelname, start)

	// If the GM response was already present, the following would
	// be all that was needed (with the levelname set as above)...

	info := gofighter.TradingInfoFromName(levelname)
	gofighter.PrintJSON(info)
}

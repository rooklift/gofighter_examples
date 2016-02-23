package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {
	
	levelname := gofighter.NameFromUser()

	// We load the original response we got when we started the game, this is
	// located in the gm/ directory and containts the instance ID.

	level, _ := gofighter.LoadGMfile(levelname)

	stop, _ := gofighter.GMstop("", level.InstanceId)
	gofighter.PrintJSON(stop)
}

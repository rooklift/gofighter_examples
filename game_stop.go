package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {

	// First we use the list of known levels to get a level name from the user.

	levelname := gofighter.NameFromUser("known_levels.json")

	// We load the original response we got when we started the game, this is
	// located in the /gm/ directory and contains the instance ID.

	level, _ := gofighter.LoadGMfile(levelname)

	stop, _ := gofighter.GMstop("", level.InstanceId)	// When given "", the function will use the API key from api_key.txt
	gofighter.PrintJSON(stop)
}

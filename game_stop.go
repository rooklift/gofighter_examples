package main

import (
	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

func main() {
	levelname := gofighter.NameFromUser()
	level, _ := gofighter.LoadGMfile(levelname)

	stop, _ := gofighter.GMstop("", level.InstanceId)
	gofighter.PrintJSON(stop)
}

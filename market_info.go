package main

// This program merely displays constantly-updated market info.

import (
    "fmt"
    "time"

    "github.com/fohristiwhirl/gofighter"        // go get -u github.com/fohristiwhirl/gofighter
)

func main()  {

    // We assume game_start.exe has been called

    levelname := gofighter.NameFromUser()
    info := gofighter.TradingInfoFromName(levelname)

    var market gofighter.Market
    market.Init(info, gofighter.Ticker)

    for {
        market.Update()
        t := time.Now()
        fmt.Printf(t.Format("2006-01-02 15:04:05 ... "))
        market.Print()
        time.Sleep(2 * time.Second)
    }
}

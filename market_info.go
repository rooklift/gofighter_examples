package main

// This program merely displays constantly-updated market info.

import (
    "fmt"
    "time"

    "github.com/fohristiwhirl/gofighter"        // go get -u github.com/fohristiwhirl/gofighter
)

func main()  {

    info := gofighter.GetUserSelection("known_levels.json")

    market_queries := make(chan chan gofighter.Market)
    go gofighter.MarketWatch(info, market_queries)

    // How this works: the goroutine MarketWatch uses a WebSocket to keep track of
    // the market, and when we send a channel to it via its input channel, it sends
    // a copy of the market back to us along the channel we provided.

    for {
        market := gofighter.GetMarket(market_queries)   // Behind the scenes, this sends a
                                                        // one-time channel to the goroutine

        fmt.Printf(time.Now().Format("2006-01-02 15:04:05 ... "))
        market.Print()

        time.Sleep(2 * time.Second)
    }
}

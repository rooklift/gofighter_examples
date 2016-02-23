package main

// This program merely displays constantly-updated position info.

import (
    "fmt"
    "time"

    "github.com/fohristiwhirl/gofighter"        // go get -u github.com/fohristiwhirl/gofighter
)

func main()  {

    info := gofighter.GetUserSelection("known_levels.json")

    position_queries := make(chan chan gofighter.Position)
    go gofighter.PositionWatch(info, position_queries)

    // How this works: the goroutine PositionWatch uses a WebSocket to keep track of
    // the position, and when we send a channel to it via its input channel, it sends
    // a copy of the position back to us along the channel we provided.

    for {
        position := gofighter.GetPosition(position_queries) // Behind the scenes, this sends a
                                                            // one-time channel to the goroutine

        fmt.Printf(time.Now().Format("2006-01-02 15:04:05 ... "))
        fmt.Printf("Shares: %d, Dollars: $%d\n", position.Shares, position.Cents / 100)

        time.Sleep(2 * time.Second)
    }
}

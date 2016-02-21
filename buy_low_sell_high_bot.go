package main

// This is about the dumbest thing that can conceivably make money. It's used
// in some of our unofficial little PvP games as a really-slow market maker.
//
// The strategy: if we have negative shares, try buying at current price - 50
//               if we have positive shares, try selling at current price + 50

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/fohristiwhirl/gofighter"        // go get -u github.com/fohristiwhirl/gofighter
)

const (
    ACCOUNT = "BLSHBOTS"
    VENUE = "TESTEX"
    SYMBOL = "FOOBAR"
    KEY = "blshkey"
    BASE_URL = "http://127.0.0.1:8000/ob/api"   // No trailing slashes please
    WS_URL = "ws://127.0.0.1:8000/ob/api/ws"
)

// -----------------------------------------------------------------------------------------------

func order_and_cancel(info gofighter.TradingInfo, order gofighter.ShortOrder) {
    res, err := gofighter.Execute(info, order, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    time.Sleep(5 * time.Second)
    id := res.Id
    gofighter.Cancel(info, id)
}

func main() {

    info := gofighter.TradingInfo{
        Account: ACCOUNT,
        Venue: VENUE,
        Symbol: SYMBOL,
        ApiKey: KEY,
        BaseURL: BASE_URL,
        WebSocketURL: WS_URL,
    }

    // The market and position will be watched by 2 goroutines. In order to get info
    // back from them, we create 2 channels that we can use to request the current
    // state from them. We request the state by sending... a channel, of course.

    market_queries := make(chan chan gofighter.Market)
    go gofighter.MarketWatch(info, market_queries)

    position_queries := make(chan chan gofighter.Position)
    go gofighter.PositionWatch(info, position_queries)

    for {
        market := gofighter.GetMarket(market_queries)   // Behind the scenes, this sends a channel
                                                        // and gets the response through it.
        if market.LastPrice < 0 {
            fmt.Printf("Waiting for market action to start...\n")
            time.Sleep(1 * time.Second)
            continue
        }

        pos := gofighter.GetPosition(position_queries)  // So does this.

        pos.Print(market.LastPrice)

        var order gofighter.ShortOrder
        order.OrderType = "limit"
        order.Qty = 50 + rand.Intn(50)
        if pos.Shares > 0 || (pos.Shares == 0 && rand.Intn(2) == 0) {
            order.Direction = "sell"
            order.Price = market.LastPrice + 50
        } else {
            order.Direction = "buy"
            order.Price = market.LastPrice - 50
        }
        if order.Price < 0 {
            order.Price = 0
        }

        go order_and_cancel(info, order)

        time.Sleep(500 * time.Millisecond)
    }
}

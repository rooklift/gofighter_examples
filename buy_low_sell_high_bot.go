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

    var unsafe_pos gofighter.Position
    var order gofighter.ShortOrder

    var market gofighter.Market
    market.Init(info, gofighter.Ticker)

    // Spin off a seperate thread to update the position via an executions WebSocket...
    go gofighter.PositionUpdater(info, &unsafe_pos, nil, nil)

    order.OrderType = "limit"

    for {
        market.Update()

        if market.LastPrice < 0 {
            fmt.Printf("Waiting for market action to start...\n")
            time.Sleep(1 * time.Second)
            continue
        }

        unsafe_pos.Lock.Lock()
        pos := unsafe_pos
        unsafe_pos.Lock.Unlock()

        nav := pos.Cents + (pos.Shares * market.LastPrice)

        fmt.Printf("Shares: %d, Dollars: $%d, NAV: $%d\n", pos.Shares, pos.Cents / 100, nav / 100)

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

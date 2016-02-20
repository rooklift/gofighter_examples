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
)

type Market struct {
    Info            gofighter.TradingInfo
    LastPrice       int
    Bid             int
    Ask             int
    Ticker          chan gofighter.Quote
}

// -----------------------------------------------------------------------------------------------

func (m * Market) Init(info gofighter.TradingInfo)  {
    m.Ticker = make(chan gofighter.Quote, 256)
    go fake_ticker(info, m.Ticker)

    m.Info = info
    m.LastPrice = -1
    m.Bid = -1
    m.Ask = -1
}

func int_or_minus_one_from_ptr(ptr * int)  int {
    if ptr == nil {
        return -1
    }
    return *ptr
}

func (m * Market) Update()  int {

    // Update the market from the WebSocket results channel.
    // Return the number of WebSocket messages read.

    var count int

    loop:
    for {
        select {

            case q := <- m.Ticker:

                count++

                m.Bid = int_or_minus_one_from_ptr(q.Bid)
                m.Ask = int_or_minus_one_from_ptr(q.Ask)
                m.LastPrice = int_or_minus_one_from_ptr(q.Last)

            default:
                break loop
        }
    }

    return count
}

func fake_ticker(info gofighter.TradingInfo, quotes_chan chan gofighter.Quote)  {

    // Poor man's tickertape without WebSockets...

    for {
        res, err := gofighter.GetQuote(info)
        if err != nil {
            continue
        }
        quotes_chan <- res
        time.Sleep(500 * time.Millisecond)
    }
}

func order_and_cancel(info gofighter.TradingInfo, order gofighter.ShortOrder, moves_chan chan gofighter.Movement) {
    res, err := gofighter.Execute(info, order, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    time.Sleep(5 * time.Second)
    id := res.Id
    res, err = gofighter.Cancel(info, id)
    if err != nil {
        fmt.Println(err)
        return
    }
    move := gofighter.MoveFromOrder(res)
    moves_chan <- move
}

func pos_updater(unsafe_pos * gofighter.Position, moves_chan chan gofighter.Movement) {
    for {
        move := <- moves_chan
        unsafe_pos.Lock.Lock()
        unsafe_pos.Cents += move.Cents
        unsafe_pos.Shares += move.Shares
        unsafe_pos.Lock.Unlock()
    }
}

func main() {

    info := gofighter.TradingInfo{
        Account: ACCOUNT,
        Venue: VENUE,
        Symbol: SYMBOL,
        ApiKey: KEY,
        BaseURL: BASE_URL,
    }

    var unsafe_pos gofighter.Position
    var order gofighter.ShortOrder

    var market Market
    market.Init(info)

    moves_chan := make(chan gofighter.Movement)

    go pos_updater(&unsafe_pos, moves_chan)

    order.OrderType = "limit"

    for {
        market.Update()

        if market.LastPrice <= 0 {
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

        go order_and_cancel(info, order, moves_chan)

        time.Sleep(500 * time.Millisecond)
    }
}

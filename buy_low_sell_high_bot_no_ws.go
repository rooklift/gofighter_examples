package main

// This is about the dumbest thing that can conceivably make money. It's used
// in some of our unofficial little PvP games as a really-slow market maker.
//
// The strategy: if we have negative shares, try buying at current price - 50
//               if we have positive shares, try selling at current price + 50

import (
    "fmt"
    "math/rand"
    "sync"
    "time"

    "github.com/fohristiwhirl/gofighter"        // go get -u github.com/fohristiwhirl/gofighter
)

// This is set up to run on a server clone on localhost...

const (
    ACCOUNT = "BLSHBOTS"
    VENUE = "TESTEX"
    SYMBOL = "FOOBAR"
    KEY = "blshkey"
    BASE_URL = "http://127.0.0.1:8000/ob/api"   // No trailing slashes please
)

var UnsafeQuote gofighter.Quote
var Quote_MUTEX sync.Mutex

// -----------------------------------------------------------------------------------------------

func order_cancel_report(info gofighter.TradingInfo, order gofighter.ShortOrder, move_chan chan gofighter.Movement)  {
    res, err := gofighter.Execute(info, order, nil)
    if err != nil {
        fmt.Println(err)
        move_chan <- gofighter.Movement{}       // For consistency, send message even on failure
        return
    }
    time.Sleep(5 * time.Second)
    id := res.Id
    res, err = gofighter.Cancel(info, id)
    if err != nil {
        fmt.Println(err)
        move_chan <- gofighter.Movement{}       // For consistency, send message even on failure
        return
    }
    move_chan <- gofighter.MoveFromOrder(res)
    return
}

func quote_updater(info gofighter.TradingInfo)  {
    for {
        localquote, err := gofighter.GetQuote(info)   // this takes ages so can't lock before doing it
        if err != nil {
            time.Sleep(500 * time.Millisecond)
            continue
        }
        Quote_MUTEX.Lock()
        UnsafeQuote = localquote
        Quote_MUTEX.Unlock()
        time.Sleep(500 * time.Millisecond)
    }
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

    go quote_updater(info)

    pos := gofighter.Position{}

    move_chan := make(chan gofighter.Movement)

    for {
        Quote_MUTEX.Lock()
        localquote := UnsafeQuote
        Quote_MUTEX.Unlock()

        if localquote.Last == -1 {
            fmt.Printf("Waiting for market action to start...\n")
            time.Sleep(1 * time.Second)
            continue
        }

        lastprice := localquote.Last

        update_position:
        for {
            select {
            case move := <- move_chan:
                pos.UpdateFromMovement(move)
            default:
                break update_position
            }
        }

        pos.Print(lastprice)

        var order gofighter.ShortOrder
        order.OrderType = "limit"
        order.Qty = 50 + rand.Intn(50)
        if pos.Shares > 0 || (pos.Shares == 0 && rand.Intn(2) == 0) {
            order.Direction = "sell"
            order.Price = lastprice + 50
        } else {
            order.Direction = "buy"
            order.Price = lastprice - 50
        }
        if order.Price < 0 {
            order.Price = 0
        }

        go order_cancel_report(info, order, move_chan)

        time.Sleep(500 * time.Millisecond)
    }
}

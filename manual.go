package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fohristiwhirl/gofighter"		// go get -u github.com/fohristiwhirl/gofighter
)

const OFFICIAL_URL = "https://api.stockfighter.io/ob/api"

var base_url string
var api_key string
var order gofighter.RawOrder
var functions map[string]func([]string)
var extrahelp map[string]string

func init() {
	functions = make(map[string]func([]string))
	functions["help"] = help
	functions["execute"] = execute
	functions["cancel"] = cancel
	functions["status"] = status
	functions["quote"] = quote
	functions["orderbook"] = orderbook
	functions["heartbeat"] = heartbeat
	functions["checkvenue"] = checkvenue
	functions["venues"] = venues
	functions["stocks"] = stocks
	functions["print"] = print
	functions["local"] = local
	functions["official"] = official
	functions["account"] = account
	functions["venue"] = venue
	functions["symbol"] = symbol
	functions["stock"] = symbol
	functions["direction"] = direction
	functions["buy"] = buy
	functions["sell"] = sell
	functions["ordertype"] = ordertype
	functions["limit"] = limit
	functions["market"] = market
	functions["ioc"] = ioc
	functions["fok"] = fok
	functions["qty"] = qty
	functions["price"] = price
	functions["key"] = key
	functions["url"] = url

	extrahelp = make(map[string]string)
	extrahelp["print"] = "     <---- print current settings"
	extrahelp["execute"] = "   <---- execute order with current settings"
}

func init() {
	base_url = OFFICIAL_URL

	var err error
	api_key, err = gofighter.LoadAPIKey("api_key.txt")
	if err != nil {
		fmt.Println(err)
	}

	order.Account = "EXB123456"
	order.Venue = "TESTEX"
	order.Symbol = "FOOBAR"
	order.Direction = "buy"
	order.OrderType = "limit"
	order.Qty = 100
	order.Price = 5000
}

func getline()  string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func getlist()  []string {
	line := getline()
	return strings.Fields(line)
}

func print_error_or_json(in interface{}, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	gofighter.PrintJSON(in)
}

func print_url() {
	fmt.Printf("[URL : %s]\n", base_url)
}

func print_key() {
	fmt.Printf("[KEY : %s]\n", api_key)
}

// ----------------------------------------------------------------------------------------------

// User-callable functions.
// All must be of this form:
//
//    func whatever(args []string) {}

func help(args []string)  {
	var keys []string
    for k := range functions {
        keys = append(keys, k)
    }
    sort.Strings(keys)

	fmt.Println("Known commands:")
	for _, c := range(keys) {
		fmt.Printf("  %s %s\n", c, extrahelp[c])
	}
}

func execute(args []string)  {
	result, err := gofighter.Execute(base_url, api_key, order, nil)
	print_error_or_json(result, err)
}

func cancel(args []string)  {
	if len(args) != 2 {
		fmt.Println("Wrong number of args for cancel")
		return
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := gofighter.Cancel(base_url, api_key, order.Venue, order.Symbol, id)
	print_error_or_json(result, err)
}

func status(args []string)  {
	if len(args) != 2 {
		fmt.Println("Wrong number of args for status")
		return
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := gofighter.GetStatus(base_url, api_key, order.Venue, order.Symbol, id)
	print_error_or_json(result, err)
}

func quote(args []string)  {
	result, err := gofighter.GetQuote(base_url, api_key, order.Venue, order.Symbol)
	print_error_or_json(result, err)
}

func orderbook(args []string)  {
	result, err := gofighter.GetOrderbook(base_url, api_key, order.Venue, order.Symbol)
	print_error_or_json(result, err)
}

func heartbeat(args []string)  {
	result, err := gofighter.CheckAPI(base_url, api_key)
	print_error_or_json(result, err)
}

func checkvenue(args []string)  {
	var venue string
	if len(args) == 2 {
		venue = string(args[1])
	} else {
		venue = order.Venue
	}
	result, err := gofighter.CheckVenue(base_url, api_key, venue)
	print_error_or_json(result, err)
}

func venues(args []string)  {
	result, err := gofighter.GetVenueList(base_url, api_key)
	print_error_or_json(result, err)
}

func stocks(args []string)  {
	var venue string
	if len(args) == 2 {
		venue = string(args[1])
	} else {
		venue = order.Venue
	}
	result, err := gofighter.GetStockList(base_url, api_key, venue)
	print_error_or_json(result, err)
}

func print(args []string)  {
	gofighter.PrintJSON(order)
	print_url()
	print_key()
}

func local(args []string)  {
	base_url = "http://127.0.0.1:8000/ob/api"
	print_url()
}

func official(args []string)  {
	base_url = OFFICIAL_URL
	print_url()
}

func account(args []string)  {
	if len(args) == 2 {
		order.Account = string(args[1])
	}
	fmt.Printf("Account: %s\n", order.Account)
}

func venue(args []string)  {
	if len(args) == 2 {
		order.Venue = string(args[1])
	}
	fmt.Printf("Venue: %s\n", order.Venue)
}

func symbol(args []string)  {
	if len(args) == 2 {
		order.Symbol = string(args[1])
	}
	fmt.Printf("Symbol: %s\n", order.Symbol)
}

func direction(args []string)  {
	if len(args) == 2 {
		order.Direction = string(args[1])
	}
	fmt.Printf("Direction: %s\n", order.Direction)
}

func buy(args []string)  {
	order.Direction = "buy"
	fmt.Printf("Direction: %s\n", order.Direction)
}

func sell(args []string)  {
	order.Direction = "sell"
	fmt.Printf("Direction: %s\n", order.Direction)
}

func ordertype(args []string)  {
	if len(args) == 2 {
		order.OrderType = string(args[1])
	}
	fmt.Printf("OrderType: %s\n", order.OrderType)
}

func limit(args []string)  {
	order.OrderType = "limit"
	fmt.Printf("OrderType: %s\n", order.OrderType)
}

func market(args []string)  {
	order.OrderType = "market"
	fmt.Printf("OrderType: %s\n", order.OrderType)
}

func ioc(args []string)  {
	order.OrderType = "immediate-or-cancel"
	fmt.Printf("OrderType: %s\n", order.OrderType)
}

func fok(args []string)  {
	order.OrderType = "fill-or-kill"
	fmt.Printf("OrderType: %s\n", order.OrderType)
}

func qty(args []string)  {
	if len(args) == 2 {
		order.Qty, _ = strconv.Atoi(string(args[1]))
	}
	fmt.Printf("Qty: %d\n", order.Qty)
}

func price(args []string)  {
	if len(args) == 2 {
		order.Price, _ = strconv.Atoi(string(args[1]))
	}
	fmt.Printf("Price: %d\n", order.Price)
}

func key(args []string)  {
	if len(args) == 2 {
		api_key = string(args[1])
	}
	print_key()
}

func url(args []string)  {
	if len(args) == 2 {
		base_url = string(args[1])
		for {							// Remove trailing slashes
			if len(base_url) > 0 && base_url[len(base_url) - 1] == '/' {
				base_url = base_url[:len(base_url) - 1]
			} else {
				break
			}
		}								// Insert "http://"
		if strings.Index(base_url, "http://") == -1 && strings.Index(base_url, "https://") == -1 {
			fmt.Println("WARNING: No protocol specified, inserting http:// (not https, set that manually if needed)")
			base_url = "http://" + base_url
		}
	}
	print_url()
}

// ----------------------------------------------------------------------------------------------

func main() {
	for {
		fmt.Printf("> ")
		ls := getlist()

		command := strings.ToLower(ls[0])

		fn, ok := functions[command]
		if ok {
			fn(ls)
		}
	}
}

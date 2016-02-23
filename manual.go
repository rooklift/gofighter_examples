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

var order gofighter.ShortOrder
var info gofighter.TradingInfo

var functions map[string]func([]string)
var extrahelp map[string]string

func init()  {
	functions = make(map[string]func([]string))
	functions["help"] = help
	functions["execute"] = execute
	functions["cancel"] = cancel
	functions["status"] = status
	functions["statusall"] = statusall
	functions["statusstock"] = statusstock
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
	functions["bid"] = buy
	functions["sell"] = sell
	functions["ask"] = sell
	functions["ordertype"] = ordertype
	functions["limit"] = limit
	functions["market"] = market
	functions["ioc"] = ioc
	functions["fok"] = fok
	functions["qty"] = qty
	functions["price"] = price
	functions["key"] = key
	functions["url"] = url
	functions["load"] = load
	functions["reset"] = reset

	extrahelp = make(map[string]string)
	extrahelp["print"] = "     <---- print current settings"
	extrahelp["execute"] = "   <---- execute order with current settings"
	extrahelp["load"] = "      <---- load info from a saved GM result from game_start.exe"

	var args []string
	reset(args)
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

func print_error_or_json(in interface{}, err error)  {
	if err != nil {
		fmt.Println(err)
		return
	}
	gofighter.PrintJSON(in)
}

func print_url()  {
	fmt.Printf("[URL : %s]\n", info.BaseURL)
}

func print_key()  {
	fmt.Printf("[KEY : %s]\n", info.ApiKey)
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
	result, err := gofighter.Execute(info, order, nil)
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

	result, err := gofighter.Cancel(info, id)
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

	result, err := gofighter.GetStatus(info, id)
	print_error_or_json(result, err)
}

func statusall(args []string)  {
	result, err := gofighter.StatusAllOrders(info)
	print_error_or_json(result, err)
}

func statusstock(args []string)  {
	result, err := gofighter.StatusAllOrdersOneStock(info)
	print_error_or_json(result, err)
}

func quote(args []string)  {
	result, err := gofighter.GetQuote(info)
	print_error_or_json(result, err)
}

func orderbook(args []string)  {
	result, err := gofighter.GetOrderbook(info)
	print_error_or_json(result, err)
}

func heartbeat(args []string)  {
	result, err := gofighter.CheckAPI(info)
	print_error_or_json(result, err)
}

func checkvenue(args []string)  {
	result, err := gofighter.CheckVenue(info)
	print_error_or_json(result, err)
}

func venues(args []string)  {
	result, err := gofighter.GetVenueList(info)
	print_error_or_json(result, err)
}

func stocks(args []string)  {
	result, err := gofighter.GetStockList(info)
	print_error_or_json(result, err)
}

func print(args []string)  {
	gofighter.PrintJSON(info)
	gofighter.PrintJSON(order)
}

func local(args []string)  {
	var port int = 8000
	var err error

	if len(args) == 2 {
		port, err = strconv.Atoi(args[1])
		if err != nil {
			port = 8000
		}
	}
	info.BaseURL = fmt.Sprintf("http://127.0.0.1:%d/ob/api", port)
	print_url()
}

func official(args []string)  {
	info.BaseURL = OFFICIAL_URL
	print_url()
}

func account(args []string)  {
	if len(args) == 2 {
		info.Account = string(args[1])
	}
	fmt.Printf("Account: %s\n", info.Account)
}

func venue(args []string)  {
	if len(args) == 2 {
		info.Venue = string(args[1])
	}
	fmt.Printf("Venue: %s\n", info.Venue)
}

func symbol(args []string)  {
	if len(args) == 2 {
		info.Symbol = string(args[1])
	}
	fmt.Printf("Symbol: %s\n", info.Symbol)
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
		info.ApiKey = string(args[1])
	}
	print_key()
}

func url(args []string)  {
	if len(args) == 2 {
		info.BaseURL = string(args[1])
		for {							// Remove trailing slashes
			if len(info.BaseURL) > 0 && info.BaseURL[len(info.BaseURL) - 1] == '/' {
				info.BaseURL = info.BaseURL[:len(info.BaseURL) - 1]
			} else {
				break
			}
		}								// Insert "http://"
		if strings.Index(info.BaseURL, "http://") == -1 && strings.Index(info.BaseURL, "https://") == -1 {
			fmt.Println("WARNING: No protocol specified, inserting http:// (not https, set that manually if needed)")
			info.BaseURL = "http://" + info.BaseURL
		}
	}
	print_url()
}

func load(args []string)  {
	info := gofighter.GetUserSelection("known_levels.json")
	gofighter.PrintJSON(info)
}

func reset(args []string)  {
	var err error
	info.BaseURL = OFFICIAL_URL
	info.ApiKey, err = gofighter.LoadAPIKey("api_key.txt")
	if err != nil {
		fmt.Println(err)
	}
	info.Account = "EXB123456"
	info.Venue = "TESTEX"
	info.Symbol = "FOOBAR"
	order.Direction = "buy"
	order.OrderType = "limit"
	order.Qty = 100
	order.Price = 5000

	gofighter.PrintJSON(info)
	gofighter.PrintJSON(order)
}

// ----------------------------------------------------------------------------------------------

func set_price_directly(s string)  {		// Called if the user types "$43.20" or similar
	var dollars int
	var cents int

	n := s[1:]
	dollars_and_cents := strings.Split(string(n), ".")

	dollars, _ = strconv.Atoi(dollars_and_cents[0])

	if len(dollars_and_cents) > 1 {
		cent_str := dollars_and_cents[1]
		if len(cent_str) > 2 {
			cent_str = cent_str[:2]
		}
		if len(cent_str) == 1 {
			cent_str += "0"
		}
		cents, _ = strconv.Atoi(cent_str)
	}

	order.Price = dollars * 100 + cents
	fmt.Printf("Price: %d\n", order.Price)
}

func set_qty_directly(s string)  {		// Called if the user types "200" or similar
	order.Qty, _ = strconv.Atoi(s)
	fmt.Printf("Qty: %d\n", order.Qty)
}

func main()  {
	for {
		fmt.Printf("> ")
		ls := getlist()

		if len(ls) == 0 {
			continue
		}

		command := strings.ToLower(ls[0])

		fn, ok := functions[command]
		if ok {
			fn(ls)
			continue
		}

		if command[0] == '$' {
			set_price_directly(string(command))
			continue
		}

		if command[0] >= '0' && command[0] <= '9' {
			set_qty_directly(string(command))
			continue
		}
	}
}

package lib

import (
  "os"
  "fmt"
  "strconv"

  tm "github.com/buger/goterm"
  bf "github.com/kecbigmt/ccyutils/bitflyer"
  bb"github.com/kecbigmt/ccyutils/bitbank"
)

func MakeOrder(marketFlag bool, side string, args []string){
  service := args[0]
  currency_pair := args[1]
  var order_type string
  var size float64
  var price float64
  // define order_type
  switch marketFlag{
  case true:
    order_type = "MARKET"
    size, _ = strconv.ParseFloat(args[2], 64)
  default:
    order_type = "LIMIT"
    size, _ = strconv.ParseFloat(args[2], 64)
    price, _ = strconv.ParseFloat(args[3], 64)
  }
  switch service{
  case "bitflyer":
    oid, err := bf.Order(currency_pair, order_type, side, size, price, 43200, "GTC")
    output(oid, err, service, currency_pair, order_type, side, size, price)
  case "bitbank":
    oid, err := bb.Order(currency_pair, order_type, side, size, price)
    output(oid, err, service, currency_pair, order_type, side, size, price)
  default:
    fmt.Printf("Service '%s' not found.\n", service)
  }
}

func output(oid string, err error, service string, currency_pair string, order_type string, side string, size float64, price float64){
  if err != nil{
    fmt.Println(err)
  }else{
    //ターミナルに出力
    fmt.Printf("[Info]%s done.\n", side)
    fmt.Println("[Detail]")
    w := tm.NewTable(0, 8, 1, '\t', 0)
    fmt.Fprintf(w, "Order ID\t%s\n", oid)
    fmt.Fprintf(w, "Service\t%s\n", service)
    fmt.Fprintf(w, "Currency Pair\t%s\n", currency_pair)
    fmt.Fprintf(w, "Side\t%s\n", side)
    fmt.Fprintf(w, "Order Type\t%s\n", order_type)
    fmt.Fprintf(w, "Size\t%f\n", size)
    fmt.Fprintf(w, "Price\t%f\n", price)
    tm.Println(w)
    tm.Flush()
    //環境変数に一時保存
    last_order_info := fmt.Sprintf(`
    {
      "sercice":"%s",
      "currency_pair":"%s",
      "oid":"%s"
    }`, service, currency_pair, oid)
    err := os.Setenv("LAST_ORDER_INFO", last_order_info)
    if err != nil{
      panic(err)
    }
  }
}

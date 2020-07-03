package lib

import (
  "fmt"

  bf "github.com/kecbigmt/ccyutils/bitflyer"
  bb"github.com/kecbigmt/ccyutils/bitbank"
)

func CancelOrder(service string, currency_pair string, oid string){
  var success bool
  var err error
  switch service{
  case "bitflyer":
    success, err = bf.Cancel(currency_pair, oid)
  case "bitbank":
    success, err = bb.Cancel(currency_pair, oid)
  default:
    err = fmt.Errorf("Service '%s' not found.\n", service)
  }
  switch success{
  case true:
    fmt.Printf("[Info]Cancel done(oid=%s).\n", oid)
  case false:
    fmt.Printf("[Error]Failed to cancel(oid=%s).\n", oid)
    fmt.Println(err)
  }
}

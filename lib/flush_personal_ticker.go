package lib

import (
  "fmt"
  "time"

  tm "github.com/buger/goterm"
  tls "github.com/kecbigmt/ccyutils/tools"
)

func Flush(p tls.PersonalTickArray, p0 tls.PersonalTickArray) {
  // Based on http://golang.org/pkg/text/tabwriter
  w := tm.NewTable(0, 8, 1, '\t', 0)
  group := "[Info]\t \t[Ticker]\t \t \t \t[Balance]\t \t \t"
  header := fmt.Sprintf("Service\tCode\tLastPrice(%v)\tLastPrice(BTC)\tSpread\tVolume\tAmount\tAmount(%v)\tAmount(BTC)\t ", p[0].KeyCurrencyCode, p[0].KeyCurrencyCode)
  fmt.Fprintln(w, group)
  fmt.Fprintln(w, header)
  for i, t := range p{
    row := fmt.Sprintf("%s\t%s\t%.2f\t%f\t%.2f%%\t%d\t%f\t%.2f\t%f\t%s", t.ServiceName, t.CurrencyCode, t.LastPrice_key, t.LastPrice_BTC, t.Spread*100.0, int(t.Volume), t.Amount, t.Amount_key, t.Amount_BTC, updown(t.Amount_BTC, p0[i].Amount_BTC))
    fmt.Fprintln(w, row)
  }
  sum_row := fmt.Sprintf(" \t \t \t \t \t \t-TOTAL-\t%.2f\t%f\t%s", p.SumKey(), p.SumBTC(), updown(p.SumBTC(), p0.SumBTC()))
  fmt.Fprintln(w, sum_row)
  tm.Println(w)

  tm.Flush()
}
func updown(f float64, f0 float64) string{
  switch {
  case f > f0:
    return "\x1b[31m▲\x1b[0m"
  case f < f0:
    return "\x1b[34m▼\x1b[0m"
  default:
    return " "
  }
}

func FlushPersonalTicker(key_currency_code string, autoFlag bool) {
  switch {
  case autoFlag:
    tm.Clear() // Clear current screen

    //initial
    tm.MoveCursor(1, 1)
    pt := tls.PersonalTicker(key_currency_code)
    pt0 := pt
    Flush(pt, pt0)
    time.Sleep(time.Second)

    // loop
    for {
		tm.MoveCursor(1, 1)
    pt = tls.PersonalTicker(key_currency_code)
		Flush(pt, pt0)
    pt0 = pt
		time.Sleep(time.Second)
	}
  default:
    pt := tls.PersonalTicker(key_currency_code)
  	Flush(pt, pt)
  }
}

// Copyright © 2017 Kecy <keishiro.oym@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
  "os"
  "encoding/json"
	"github.com/spf13/cobra"
	"github.com/kecbigmt/ccycli/lib"
)
var lastFlag bool

// BUY command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel your order",
	Long: `cancel
  your
  order`,
	Run: func(cmd *cobra.Command, args []string) {
    var service string
    var currency_pair string
    var oid string
    switch lastFlag{
    case true:
      bytes := []byte(os.Getenv("LAST_ORDER_INFO"))
      var last_order_info map[string]string
      if err := json.Unmarshal(bytes, &last_order_info); err != nil {
          panic(err)
      }
      service = last_order_info["service"]
      currency_pair = last_order_info["currency_pair"]
      oid = last_order_info["oid"]
    default:
      service = args[0]
      currency_pair = args[1]
      oid = args[2]
    }
    lib.CancelOrder(service, currency_pair, oid)
	},
}

func init() {
  cancelCmd.Flags().BoolVarP(&lastFlag, "last", "l", false, "cancel your LAST order")
	RootCmd.AddCommand(cancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ptickerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ptickerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

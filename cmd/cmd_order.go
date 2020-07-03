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
	"github.com/spf13/cobra"
	"github.com/kecbigmt/ccycli/lib"
)
var marketFlag bool

// BUY command
var buyCmd = &cobra.Command{
	Use:   "buy",
	Short: "limit or market",
	Long: `limit
  or
  market`,
	Run: func(cmd *cobra.Command, args []string) {
    lib.MakeOrder(marketFlag, "BUY", args)
	},
}

// SELL command
var sellCmd = &cobra.Command{
	Use:   "sell",
	Short: "limit or market",
	Long: `limit
  or
  market`,
	Run: func(cmd *cobra.Command, args []string) {
    lib.MakeOrder(marketFlag, "SELL", args)
	},
}

func init() {
  buyCmd.Flags().BoolVarP(&marketFlag, "market", "m", false, "make MARKET order")
  sellCmd.Flags().BoolVarP(&marketFlag, "market", "m", false, "make MARKET order")
	RootCmd.AddCommand(buyCmd)
  RootCmd.AddCommand(sellCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ptickerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ptickerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

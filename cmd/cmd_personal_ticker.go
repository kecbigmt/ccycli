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
var autoFlag bool

// ptickerCmd represents the collect command
var ptickerCmd = &cobra.Command{
	Use:   "pticker",
	Short: "your personal ticker",
	Long: `your
  personal
  ticker`,
	Run: func(cmd *cobra.Command, args []string) {
    key_currency_code := "JPY" //default
    if len(args) > 0{
      key_currency_code = args[0]
    }
    lib.FlushPersonalTicker(key_currency_code, autoFlag)
	},
}

func init() {
  ptickerCmd.Flags().BoolVarP(&autoFlag, "auto", "a", false, "updates info automatically")
	RootCmd.AddCommand(ptickerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ptickerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ptickerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

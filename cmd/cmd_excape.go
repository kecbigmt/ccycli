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
  "strconv"
	"github.com/spf13/cobra"
	tls "github.com/kecbigmt/ccyutils/tools"
)

// command
var escapeCmd = &cobra.Command{
	Use:   "escape",
	Short: "for maintainance",
	Long: `for maintainance`,
	Run: func(cmd *cobra.Command, args []string) {
    service := args[0]
    currency_pair := args[1]
    size, err := strconv.ParseFloat(args[2], 64)
    if err != nil{
      panic(err)
    }
    tls.Escape(service, currency_pair, size)
	},
}

func init() {
	RootCmd.AddCommand(escapeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ptickerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ptickerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

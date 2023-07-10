/*
Copyright © 2022 Polygon <engineering@polygon.technology>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Lesser General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Lesser General Public License for more details.

You should have received a copy of the GNU Lesser General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package rpc

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	_ "embed"

	ethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/spf13/cobra"
)

//go:embed usage.md
var usage string

// rpcCmd represents the rpc command
var RpcCmd = &cobra.Command{
	Use:   "rpc URL method param_0 param_1 ... param_n",
	Short: "Wrapper for making RPC requests.",
	Long:  usage,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		ec, err := ethrpc.DialContext(ctx, args[0])
		if err != nil {
			return err
		}

		params := toGenericParams(args[2:])
		var res = new(json.RawMessage)
		err = ec.Call(res, args[1], params...)
		if err != nil {
			return err
		}
		body, err := res.MarshalJSON()
		if err != nil {
			fmt.Println("gyahhhhhhhh1")
			return err
		}
		fmt.Println(string(body))
		return nil
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("rpc needs at least two arguments. A URL and a method")
		}

		_, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		return nil
	},
}

func toGenericParams(args []string) []any {
	retArgs := make([]any, len(args))
	for k := range args {
		retArgs[k] = convertStringTypes(args[k])
	}
	return retArgs
}

func convertStringTypes(param string) any {
	lowerParam := strings.ToLower(param)

	if lowerParam == "true" {
		return true
	}
	if lowerParam == "false" {
		return false
	}
	if lowerParam == "null" {
		return nil
	}

	return param
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

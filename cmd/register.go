// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yulibaozi/kubectl-switch/server"
)

/*
// Execute input
func Execute() {
	args := os.Args[1:]
	argsLen := len(args)
	if argsLen <= 0 {
		fmt.Fprintln(os.Stderr, "please input corrently command")
		return
	}
	cmd := server.CmdShim{
		Run: server.Exec,
	}
	cmd.SubCmd = args[0]
	if argsLen > 1 {
		cmd.Args = args[1:]
	}
	cmd.Run(cmd)
}

*/

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:     "register",
	Short:   "Register cluster in switch plugin",
	Example: "kubectl switch register clusterName ./qa.config",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Fprintln(os.Stderr, "Error: ", "Names can only be composed of lowercase args")
			return
		}
		shim := &server.CmdShim{
			SubCmd: "register",
			Args:   args,
			Run:    server.Exec,
		}
		shim.Run(shim)
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
}

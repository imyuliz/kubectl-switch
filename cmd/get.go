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

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "View cluster details",
	Example: "kubectl switch get clusterName",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
	},
}

// https://blog.csdn.net/wangfenglin995/article/details/80113243
func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags()
}

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
	// rootCmd.AddCommand(getCmd)
	// getCmd.Flags()
}

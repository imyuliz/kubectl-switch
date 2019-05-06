package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yulibaozi/kubectl-switch/server"
)

// nowCmd represents the get command
var nowCmd = &cobra.Command{
	Use:     "now",
	Short:   "View cluster of currently in use",
	Example: "kubectl switch now",
	Run: func(cmd *cobra.Command, args []string) {
		shim := &server.CmdShim{
			SubCmd: "now",
			Args:   args,
			Run:    server.Exec,
		}
		shim.Run(shim)
	},
}

// https://blog.csdn.net/wangfenglin995/article/details/80113243
func init() {
	rootCmd.AddCommand(nowCmd)
	// nowCmd.Flags()
}

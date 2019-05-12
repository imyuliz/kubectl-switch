package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yulibaozi/kubectl-switch/server"
)

// nowCmd represents the get command
var whereCmd = &cobra.Command{
	Use:     "where",
	Short:   "View cluster of currently in use",
	Example: "kubectl switch now",
	Run: func(cmd *cobra.Command, args []string) {
		shim := &server.CmdShim{
			SubCmd: "where",
			Args:   args,
			Run:    server.Exec,
		}
		shim.Run(shim)
	},
}

// https://blog.csdn.net/wangfenglin995/article/details/80113243
func init() {
	rootCmd.AddCommand(whereCmd)
	// nowCmd.Flags()
}

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yulibaozi/kubectl-switch/server"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all cluster message",
	Example: "kubectl switch list",
	Run: func(cmd *cobra.Command, args []string) {
		shim := &server.CmdShim{
			SubCmd: "list",
			Args:   args,
			Run:    server.Exec,
		}
		shim.Run(shim)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

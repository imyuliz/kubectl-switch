package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yulibaozi/kubectl-switch/server"
)

// removeallCmd represents the removeall command
var removeallCmd = &cobra.Command{
	Use:     "removeall",
	Short:   "Removeall cluster config",
	Example: "kubectl switch removeall",
	Run: func(cmd *cobra.Command, args []string) {
		shim := &server.CmdShim{
			SubCmd: "removeall",
			Args:   args,
			Run:    server.Exec,
		}
		shim.Run(shim)
	},
}

func init() {
	rootCmd.AddCommand(removeallCmd)
}

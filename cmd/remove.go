package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yulibaozi/kubectl-switch/server"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove the specified cluster name",
	Example: "kubectl switch remove clusterName",
	Run: func(cmd *cobra.Command, args []string) {
		shim := &server.CmdShim{
			SubCmd: "remove",
			Args:   args,
			Run:    server.Exec,
		}
		shim.Run(shim)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}

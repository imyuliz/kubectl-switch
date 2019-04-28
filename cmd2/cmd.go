package cmd

import (
	"fmt"
	"os"

	"github.com/yulibaozi/kubectl-switch/server"
)

// Execute input
func Execute() {
	args := os.Args[1:]
	argsLen := len(args)
	if argsLen <= 0 {
		fmt.Fprintln(os.Stderr, "please input corrently command")
		return
	}
	cmd := &server.Command{
		Run: server.Exec,
	}
	cmd.SubCmd = args[0]
	if argsLen > 1 {
		cmd.Args = args[1:]
	}
	cmd.Run(cmd)
}
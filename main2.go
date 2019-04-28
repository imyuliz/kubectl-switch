package main

import (
	"github.com/yulibaozi/kubectl-switch/cmd"
	_ "github.com/yulibaozi/kubectl-switch/server/externalcmds"
	_ "github.com/yulibaozi/kubectl-switch/server/subcmds"
)

func main2() {
	cmd.Execute()

}

package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/vars"
)

func init() {
	CheckAllConfig()
}

// hictl  register           qa /root/admin.config

// hictl    qa               get pod -n yulibaozi
//   |       |                 |
//   |	     |                 |
//  cmd     subcmd            args                 flags

// Command 组成结构体
type Command struct {
	SubCmd      string
	ExternalCmd string
	Args        []string
	Flags       []string
	Run         func(c *Command)
}

// Exec 入口操作
var Exec = func(c *Command) {
	c.SubCmd = strings.TrimSpace(c.SubCmd)
	if strings.EqualFold(c.SubCmd, "") {
		fmt.Fprintln(os.Stderr, "SUBCMD is not allowed to be empty")
		return
	}
	c.SubCmd = strings.ToLower(c.SubCmd)
	if err := Trace(c); err != nil {
		fmt.Fprintln(os.Stderr, "exec:", err.Error())
	}

}

// Trace 判断应该执行什么命令并转发
// 先判断是否是命令,如果不是就是kubectl的命令,已为包含其他命令留下接口
func Trace(c *Command) error {
	if IsSubCmd(c.SubCmd) {
		// 做默认命令操作
		return Run(c)
	}
	clusters := GetClusterNames()
	if clusters[c.SubCmd] {
		//TODO 需要执行kubectl 命令
		c.ExternalCmd = vars.KUBECTL
		return RunExternalCmd(c)
	}
	return fmt.Errorf("not find command:%s", c.SubCmd)
}

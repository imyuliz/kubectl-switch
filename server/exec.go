package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/vars"
)

// func init() {
// 	CheckAllConfig()
// }

// hictl  register           qa /root/admin.config

// hictl    qa               get pod -n yulibaozi
//   |       |                 |
//   |	     |                 |
//  cmd     subcmd            args                 flags

// CmdShim fields
type CmdShim struct {
	SubCmd      string
	ExternalCmd string
	Args        []string
	Flags       []string
	Run         func(c *CmdShim)
}

// Exec cmd
var Exec = func(c *CmdShim) {
	c.SubCmd = strings.TrimSpace(c.SubCmd)
	if strings.EqualFold(c.SubCmd, "") {
		fmt.Fprintln(os.Stderr, "subcmd is not allowed to be empty")
		return
	}
	c.SubCmd = strings.ToLower(c.SubCmd)
	if err := Trace(c); err != nil {
		fmt.Fprintln(os.Stderr, "exec:", err.Error())
	}

}

// Trace sub cmd forward
// If it is not an internal command or an external command(kubectl).
func Trace(c *CmdShim) error {
	if IsSubCmd(c.SubCmd) {
		// exec default cmd
		return Run(c)
	}
	clusters := GetClusterNames()
	if clusters[c.SubCmd] {
		// exec kubectl cmd
		c.ExternalCmd = vars.KUBECTL
		return RunExternalCmd(c)
	}
	return fmt.Errorf("not find command %s", c.SubCmd)
}

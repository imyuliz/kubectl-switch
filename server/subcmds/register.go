package subcmds

import (
	"fmt"
	"regexp"
	"unicode"

	"github.com/yulibaozi/kubectl-switch/server"
)

var _ server.SubCommand = &Register{}

// Register register cmd
type Register struct{}

// register cmd
func init() {
	server.RegisterSubCmd((*Register)(nil))

}

// Validation Verify that the parameters match
func (r *Register) Validation(cmd *server.CmdShim) bool {
	if cmd.SubCmd == "" {
		return false
	}
	if len(cmd.Args) > 0 && cmd.Args[0] == "" {
		return false
	}
	return true
}

func isWord(str string) {
	for _, v := range str {
		unicode.IsLetter(v)

	}
}

// Exec func
// hictl  register           qa /root/admin.config
// hictl    qa               get pod -n yulibaozi
//   |       |                 |
//   |	     |                 |
//  cmd     subcmd            args                 flags
func (r *Register) Exec(cmd *server.CmdShim) error {
	if r.Validation(cmd) {
		clusterName := cmd.Args[0]
		if match, _ := regexp.MatchString("^[a-z]+$", clusterName); !match {
			return fmt.Errorf("names can only be composed of lowercase letters")
		}
		clusterNames := server.GetClusterNames()
		clusterPath := server.GetConfigDir(clusterName)
		if !clusterNames[clusterName] {
			if err := server.MKDir(clusterPath); err != nil {
				return fmt.Errorf("register failed. mkdir:%s err:%v", clusterPath, err)
			}
			// check config path
			if len(cmd.Args) > 1 && cmd.Args[1] != "" {
				if err := server.CopyConfig(cmd.Args[1], clusterPath); err != nil {
					return err
				}
				return fmt.Errorf("congratulate! %s cluster register succeed", clusterName)
			}
			return fmt.Errorf("%s cluster register succeed. but you must mv cluster config mv to:%s", clusterName, clusterPath)
		}
		return fmt.Errorf("%s cluster already registered, no need to register again. if need update. please update config:%s", clusterName, clusterPath)
	}
	return fmt.Errorf("please input the correct subcommand")
}

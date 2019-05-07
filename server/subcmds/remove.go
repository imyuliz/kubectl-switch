package subcmds

import (
	"errors"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/fileutil"

	"github.com/yulibaozi/kubectl-switch/server"
)

// Remove remove cmd
type Remove struct{}

// register cmd
func init() {
	server.RegisterSubCmd((*Remove)(nil))

}

var _ server.SubCommand = &Remove{}

// Exec exec remove
func (r *Remove) Exec(cmd *server.CmdShim) error {
	if len(cmd.Args) <= 0 {
		return errors.New("not found cluster name")
	}
	for i := range cmd.Args {
		clusterName := strings.ToLower(cmd.Args[i])
		if clusterName == "" {
			break
		}
		clusterNames := server.GetClusterNames()
		if clusterNames[clusterName] {
			clusterPath := server.GetConfigDir(clusterName)
			// del dir
			return fileutil.DelDir(clusterPath)
		}
	}
	return nil
}

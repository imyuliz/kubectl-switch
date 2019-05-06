package subcmds

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server"

	"github.com/yulibaozi/kubectl-switch/server/fileutil"
)

// Now show use cluster
type Now struct{}

func init() {
	server.RegisterSubCmd((*Now)(nil))
}

// Exec now cmd
func (n *Now) Exec(cmd *server.CmdShim) error {
	kubePath, err := server.GetKubeConfigPath()
	if err != nil {
		return err
	}
	exist, dir := fileutil.PathStatus(kubePath)
	if !exist || (exist && dir) {
		return fmt.Errorf("not found kube config")
	}

	kubeClient, err := server.GetClient(kubePath)
	if err != nil {
		return err
	}
	kubeHost := kubeClient.Host
	clusterNames := server.GetClusterNames()
	for k := range clusterNames {
		configPath, err := server.GetConfigNameByClusterName(k)
		if err != nil {
			return err
		}
		client, err := server.GetClient(configPath)
		if err != nil {
			return err
		}
		if strings.EqualFold(client.Host, kubeHost) {
			fmtNow(k, nil)
			return nil
		}
	}
	return errors.New("not found cluster")
}

func fmtNow(name string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "get info err:%v \n", err)
		return
	}
	fmt.Fprintf(os.Stdout, "The cluster currently in use is: %s \n", name)

}

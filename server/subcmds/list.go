package subcmds

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/yulibaozi/kubectl-switch/server"
	"github.com/yulibaozi/kubectl-switch/server/fileutil"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// List cmd
type List struct{}

var _ server.SubCommand = &List{}

// list cmd
func init() {
	server.RegisterSubCmd((*List)(nil))
}

// Validation validation cmd
func (l *List) Validation(cmd *server.CmdShim) bool {
	return true
}

func fmtLine(name, host, md5 string) {
	fmt.Println(fmt.Sprintf("|---  %-7s---|--- %-27s ---|--- %-32s ---|", name, host, md5))
}

// Exec  list
func (l *List) Exec(cmd *server.CmdShim) error {
	clusterNames := server.GetClusterNames()
	fmt.Println("|---- Name -----|--------------- Host --------------|----------------- Md5 ------------------|")
	for name := range clusterNames {
		clusterPath := server.GetConfigDir(name)
		fileInfos, err := ioutil.ReadDir(clusterPath)
		if err != nil {
			fmtLine(name, "-", "-")
			return err
		}
		fileCount := server.FileCount(fileInfos)
		if fileCount != 1 {
			fmtLine(name, "config err", "-")
			continue
		}
		for i := range fileInfos {
			if !fileInfos[i].IsDir() {
				path := filepath.Join(clusterPath, fileInfos[i].Name())
				md5sum, err := fileutil.Md5Sum(path)
				if err != nil {
					return err
				}
				// server := getServer(path)
				config, err := configFromPath(path)
				if err != nil {
					return err
				}
				client, err := config.ClientConfig()
				if err != nil {
					return err
				}
				fmtLine(name, client.Host, md5sum)
				break
			}
		}
	}
	fmt.Println("|---------------|-----------------------------------|----------------------------------------|")
	return nil
}

func configFromPath(path string) (clientcmd.ClientConfig, error) {
	rules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: path}
	credentials, err := rules.Load()
	if err != nil {
		return nil, fmt.Errorf("the provided credentials %q could not be loaded: %v", path, err)
	}

	overrides := &clientcmd.ConfigOverrides{
		Context: clientcmdapi.Context{
			Namespace: os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_NAMESPACE"),
		},
	}

	var cfg clientcmd.ClientConfig
	context := os.Getenv("KUBECTL_PLUGINS_GLOBAL_FLAG_CONTEXT")
	if len(context) > 0 {
		rules := clientcmd.NewDefaultClientConfigLoadingRules()
		cfg = clientcmd.NewNonInteractiveClientConfig(*credentials, context, overrides, rules)
	} else {
		cfg = clientcmd.NewDefaultClientConfig(*credentials, overrides)
	}

	return cfg, nil
}

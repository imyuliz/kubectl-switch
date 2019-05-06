package subcmds

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/yulibaozi/kubectl-switch/server"
	"github.com/yulibaozi/kubectl-switch/server/fileutil"
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
				client, err := server.GetClient(path)
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

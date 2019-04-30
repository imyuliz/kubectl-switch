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

// list cmd
func init() {
	server.RegisterSubCmd((*List)(nil))
}

// Validation validation cmd
func (l *List) Validation(cmd *server.CmdShim) bool {
	return true
}

func fmtLine(name, md5, msg string) {
	fmt.Println(fmt.Sprintf("|--- %s ---|--- %s ---|--- %s ---|", name, md5, msg))
}

// Exec  list
func (l *List) Exec(cmd *server.CmdShim) error {
	clusterNames := server.GetClusterNames()
	fmt.Println("|--- name ---|--- md5sum ---|--- msg ---|")
	for name := range clusterNames {
		clusterPath := server.GetConfigDir(name)
		fileInfos, err := ioutil.ReadDir(clusterPath)
		if err != nil {
			fmtLine(name, "-", "can't find config")
			return err
		}
		fileCount := server.FileCount(fileInfos)
		if fileCount != 1 {
			fmtLine(name, "-", "config err")
			continue
		}
		for i := range fileInfos {
			if !fileInfos[i].IsDir() {
				path := filepath.Join(clusterPath, fileInfos[i].Name())
				md5sum, err := fileutil.Md5Sum(path)
				if err != nil {
					fmtLine(name, "-", "md5sum err")
					break
				}
				fmtLine(name, md5sum, "-")
				break
			}
		}
	}
	fmt.Println("|------------|--------------|-----------|")
	return nil
}

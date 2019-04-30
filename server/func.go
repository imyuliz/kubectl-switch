package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/yulibaozi/kubectl-switch/server/vars"

	env "github.com/yulibaozi/kubectl-switch/server/environment"
	"github.com/yulibaozi/kubectl-switch/server/fileutil"
)

// GetClusterNames list cluster
func GetClusterNames() map[string]bool {
	base := fileutil.GetBase()
	clusterNames, err := fileutil.SubDirs(base)
	if err != nil {
		return nil
	}
	return clusterNames
}

// AddCluster mkdir cluster config dir
func AddCluster(clusterName string) error {
	base := fileutil.GetBase()
	clusterDir := filepath.Join(base, clusterName)
	return fileutil.MkdirAll(clusterDir)
}

// MKDir mkdir path
func MKDir(path string) error {
	return fileutil.MkdirAll(path)
}

// GetConfigDir get cluster config path
func GetConfigDir(clusterName string) string {
	base := fileutil.GetBase()
	return filepath.Join(base, clusterName)
}

// CopyConfig  cp file
func CopyConfig(srcPath, desPath string) error {
	isExist, isDir := fileutil.PathStatus(srcPath)
	if !isExist {
		return fmt.Errorf("src path %s is not existd", srcPath)
	}
	if isDir {
		return fmt.Errorf("src path %s must end with file name", srcPath)
	}
	fileName := fileutil.GetFileName(srcPath)
	desPath = fileutil.Join(desPath, fileName)
	isExist, isDir = fileutil.PathStatus(desPath)
	if isExist && !isDir {
		return errors.New("the file already exists")
	}
	if err := fileutil.MkFile(desPath); err != nil {
		return err
	}
	return fileutil.Copy(srcPath, desPath)
}

// IsCluster check cluster env
// true cluster already registered
// false is not cluster
func IsCluster(subCmd string) bool {
	cluster := GetClusterNames()
	return cluster[subCmd]
}

// IsSubCmd check default cmd
func IsSubCmd(subCmd string) bool {
	options := vars.GetSubCmds()
	return options[subCmd]
}

// CheckAllConfig check if the configuration is correct
func CheckAllConfig() {
	base := fileutil.GetBase()
	dirs, err := fileutil.SubDirs(base)
	if err != nil {
		fmt.Fprintf(os.Stdout, "WARNNING: check cluster config filed. err:%v\n", err)
		return
	}
	if len(dirs) <= 0 {
		fmt.Fprintf(os.Stdout, "WARNNING: not found cluster info,please register first\n")
		return
	}
	for dir := range dirs {
		path := filepath.Join(base, dir)
		fileInfos, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Fprintf(os.Stdout, "WARNNING: check cluster %s config failed. err:%v\n", dir, err)
			return

		}
		fileCount := FileCount(fileInfos)
		WarnConfig(path, dir, fileCount)
	}
}

// WarnConfig config warnnging
func WarnConfig(path, clusteName string, fileCount int) {
	if fileCount < 1 {
		fmt.Fprintf(os.Stdout, "WARNNING: missing configuration file of %s cluster, please add config under path %s.\n", clusteName, path)
		return
	}
	if fileCount > 1 {
		fmt.Fprintf(os.Stdout, "WARNNING: too many configuration files of %s cluster, please remove under path %s.\n", clusteName, path)
		return
	}

}

// GetConfigNameByClusterName  get cluster config by cluster name
func GetConfigNameByClusterName(name string) (string, error) {
	clusterPath := GetConfigDir(name)
	fileInfos, err := ioutil.ReadDir(clusterPath)
	if err != nil {
		return "", err
	}
	for i := range fileInfos {
		// file
		if !fileInfos[i].IsDir() {
			return filepath.Join(clusterPath, fileInfos[i].Name()), nil
		}
	}
	return "", fmt.Errorf("cann't find config in: %s", clusterPath)
}

// CheckConfig check is correct
func CheckConfig(path string) bool {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARINNING: check path: %s config failed.", err)
		return false
	}
	fileCount := FileCount(fileInfos)
	if fileCount != 1 {
		fmt.Fprintf(os.Stdout, "WARNNING: %s cluster config file is abnormal", path)
		return false
	}
	return true

}

func FileCount(fileInfos []os.FileInfo) int {
	fileCount := 0
	for i := range fileInfos {
		// file
		if !fileInfos[i].IsDir() {
			fileCount = fileCount + 1
		}
	}
	return fileCount
}

// GetKubeConfigPath get kube config path
func GetKubeConfigPath() string {
	return filepath.Join(env.GetHome(), vars.KUBECONFIGPATH, vars.KUBECONFIGFILE)
}

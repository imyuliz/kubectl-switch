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

// GetClusterNames 获取集群名字列表
func GetClusterNames() map[string]bool {
	base := fileutil.GetBase()
	clusterNames, err := fileutil.SubDirs(base)
	if err != nil {
		return nil
	}
	return clusterNames
}

// AddCluster 给新注册的集群添加一个配置文件目录
func AddCluster(clusterName string) error {
	base := fileutil.GetBase()
	clusterDir := filepath.Join(base, clusterName)
	return fileutil.MkdirAll(clusterDir)
}

// MKDir 根据path创建文件夹
func MKDir(path string) error {
	return fileutil.MkdirAll(path)
}

// GetConfigDir 获取单个集群的配置目录
func GetConfigDir(clusterName string) string {
	base := fileutil.GetBase()
	return filepath.Join(base, clusterName)
}

// CopyConfig 判断是否是目录
func CopyConfig(srcPath, desPath string) error {
	isExist, isDir := fileutil.PathStatus(srcPath)
	if !isExist {
		return fmt.Errorf("源路径:%s 不存在,请检查", srcPath)
	}
	if isDir {
		return fmt.Errorf("源路径:%s 必须以文件名结尾", srcPath)
	}
	fileName := fileutil.GetFileName(srcPath)
	desPath = fileutil.Join(desPath, fileName)
	isExist, isDir = fileutil.PathStatus(desPath)
	if isExist && !isDir { //如果存在这个文件
		return errors.New("文件已经存在")
	}
	if err := fileutil.MkFile(desPath); err != nil {
		return err
	}
	return fileutil.Copy(srcPath, desPath)
}

// IsCluster 查看是否是集群环境
// true 是已经注册了的集群
// false 不是
func IsCluster(subCmd string) bool {
	cluster := GetClusterNames()
	return cluster[subCmd]
}

// IsSubCmd 是否是自带的命令
func IsSubCmd(subCmd string) bool {
	options := vars.GetSubCmds()
	return options[subCmd]
}

// CheckAllConfig 检查所有集群的配置是否正确
func CheckAllConfig() {
	base := fileutil.GetBase()
	dirs, err := fileutil.SubDirs(base)
	if err != nil {
		fmt.Fprintf(os.Stdout, "WARNNING: 检查配置失败:%v\n", err)
		return
	}
	if len(dirs) <= 0 {
		fmt.Fprintf(os.Stdout, "WARNNING: switch暂无集群注册,请先注册后再使用!\n")
		return
	}
	for dir := range dirs {
		path := filepath.Join(base, dir)
		fileInfos, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Fprintf(os.Stdout, "WARNNING: 检查集群:%s 配置失败:%v\n", dir, err)
			return

		}
		fileCount := fileCount(fileInfos)
		WarnConfig(dir, fileCount)
	}
}

// WarnConfig 配置文件警告
func WarnConfig(clusteName string, fileCount int) {
	if fileCount < 1 {
		fmt.Fprintf(os.Stdout, "WARNNING: 集群:%s 下的配置文件缺失,请添加才能正常使用.\n", clusteName)
		return
	}
	if fileCount > 1 {
		fmt.Fprintf(os.Stdout, "WARNNING: 集群:%s 下的配置文件过多,清理后才能正常使用.\n", clusteName)
		return
	}

}

// GetConfigNameByClusterName 通过集群名字获取配置文件
func GetConfigNameByClusterName(name string) (string, error) {
	clusterPath := GetConfigDir(name)
	fileInfos, err := ioutil.ReadDir(clusterPath)
	if err != nil {
		return "", err
	}
	for i := range fileInfos {
		// 如果是文件
		if !fileInfos[i].IsDir() {
			return filepath.Join(clusterPath, fileInfos[i].Name()), nil
		}
	}
	return "", fmt.Errorf("cann't find config in: %s", clusterPath)
}

// CheckConfig 检查某个集群下的配置是否正常
func CheckConfig(path string) bool {
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "WARINNING: 检查PATH: %s 的配置失败.", err)
		return false
	}
	fileCount := fileCount(fileInfos)
	if fileCount != 1 {
		fmt.Fprintf(os.Stdout, "WARNNING: 集群:%s 下的配置存在异常", path)
		return false
	}
	return true

}

func fileCount(fileInfos []os.FileInfo) int {
	fileCount := 0
	for i := range fileInfos {
		// 如果是文件
		if !fileInfos[i].IsDir() {
			fileCount = fileCount + 1
		}
	}
	return fileCount
}

// GetKubeConfigPath 获取使用的kube config文件名字
func GetKubeConfigPath() string {
	return filepath.Join(env.GetHome(), vars.KUBECONFIGPATH, vars.KUBECONFIGFILE)
}

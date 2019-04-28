package fileutil

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/vars"

	env "github.com/yulibaozi/kubectl-switch/server/environment"
)

// GetBase 获取项目的基本目录
func GetBase() string {
	return filepath.Join(env.GetHome(), vars.DEFAULTDIR)
}

// PathStatus 判断给定的路径是是否存在，且是否是目录
// exist: true 存在; false 不存在
// dir: true 是文件夹; false 是文件
func PathStatus(path string) (exist, dir bool) {
	s, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true, false
		}
		return false, false
	}
	return true, s.IsDir()

}

// MkdirAll 创建相关目录，父目录如果不存在，创建
func MkdirAll(path string) error {
	return os.MkdirAll(path, 0777)
}

// Touch 创建配置文件
func Touch(configPath string) error {
	//如果已经存在就不管
	if exist, dir := PathStatus(configPath); exist && !dir {
		return nil
	}
	dirPath, _ := filepath.Split(configPath)
	exist, dir := PathStatus(dirPath)
	// 如果这个path不是目录，或者不存在就创建
	if (exist && !dir) || !exist {
		err := MkdirAll(dirPath)
		if err != nil {
			return err
		}
	}
	return MkFile(configPath)
}

// MkFile 创建文件
func MkFile(path string) error {
	// file, err := os.OpenFile(path, os.O_RDWR|O_CREATE|O_TRUNC, 0777)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// GetFileName 获取目录中的文件名
func GetFileName(path string) string {
	return filepath.Base(path)
}

// Join 拼接地址和文件
func Join(path, fileName string) string {
	return filepath.Join(path, fileName)
}

// SubDirs 获取一个目录下的所有子目录
func SubDirs(path string) (map[string]bool, error) {
	dirs := map[string]bool{}
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for i, _ := range fileInfos {
		if fileInfos[i].IsDir() {
			dirs[fileInfos[i].Name()] = true
		}
	}
	return dirs, nil
}

// SubFiles 获取一个目录下的文件
func SubFiles(path string) ([]string, error) {
	var files []string
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for i := range fileInfos {
		// 如果是文件
		if !fileInfos[i].IsDir() {
			files = append(files, fileInfos[i].Name())
		}
	}
	return files, nil
}

// DelDir 删除文件夹
func DelDir(path string) error {
	return os.RemoveAll(path)
}

// Copy 文件复制 效率最高的一种方式
// https://www.jb51.net/article/148552.htm
func Copy(source, des string) error {
	// sourceFile, err := os.OpenFile(source, os.O_RDWR, 0666)
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()
	desFile, err := os.OpenFile(des, os.O_RDWR, 0666)
	// desFile, err := os.Open(des)
	if err != nil {
		return err
	}
	defer desFile.Close()
	buf := make([]byte, vars.BUFFERSIZE)
	for {
		n, err := sourceFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		if _, err := desFile.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// FileEmpty 判断一个文件是否为空
// true 为空
// false 不为空
func FileEmpty(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return true, err
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return true, err
	}
	return fileStat.Size() == 0, nil
}

// FileMd5Equal 判断两个文件是否相等
func FileMd5Equal(srcFile, desFile string) (bool, error) {
	srcMd5 := md5.New()
	src, err := os.Open(srcFile)
	if err != nil {
		return false, err
	}
	defer src.Close()
	io.Copy(srcMd5, src)
	srcMd5Str := hex.EncodeToString(srcMd5.Sum(nil))
	desMd5 := md5.New()
	des, err := os.Open(desFile)
	if err != nil {
		return false, err
	}
	defer des.Close()
	io.Copy(desMd5, src)
	desMd5Str := hex.EncodeToString(desMd5.Sum(nil))
	return strings.EqualFold(srcMd5Str, desMd5Str), nil
}

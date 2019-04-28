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

// GetBase get root path
func GetBase() string {
	return filepath.Join(env.GetHome(), vars.DEFAULTDIR)
}

// PathStatus  path is exist and  path is dir
// exist: true exist; false not exist
// dir: true is dir; false not
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

// MkdirAll make all dir,
// if parent dir is not exist,mkdir parent dir. like: cmd mkdir -p
func MkdirAll(path string) error {
	return os.MkdirAll(path, 0777)
}

// Touch  create file.
// if parent dir is not exist. mkdir parent dir. like cmd: touch
func Touch(configPath string) error {
	// Ignore configPath if it exists
	if exist, dir := PathStatus(configPath); exist && !dir {
		return nil
	}
	dirPath, _ := filepath.Split(configPath)
	exist, dir := PathStatus(dirPath)
	// create file when path is not dir or is not exist
	if (exist && !dir) || !exist {
		err := MkdirAll(dirPath)
		if err != nil {
			return err
		}
	}
	return MkFile(configPath)
}

// MkFile create file
func MkFile(path string) error {
	// file, err := os.OpenFile(path, os.O_RDWR|O_CREATE|O_TRUNC, 0777)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

// GetFileName Get the file name under the specified path
func GetFileName(path string) string {
	return filepath.Base(path)
}

// Join join path and file name
func Join(path, fileName string) string {
	return filepath.Join(path, fileName)
}

// SubDirs get all file names under the specified path map
func SubDirs(path string) (map[string]bool, error) {
	dirs := map[string]bool{}
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for i := range fileInfos {
		if fileInfos[i].IsDir() {
			dirs[fileInfos[i].Name()] = true
		}
	}
	return dirs, nil
}

// SubFiles get all file name under the specified path slice
func SubFiles(path string) ([]string, error) {
	var files []string
	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for i := range fileInfos {
		if !fileInfos[i].IsDir() {
			files = append(files, fileInfos[i].Name())
		}
	}
	return files, nil
}

// DelDir delete dir
func DelDir(path string) error {
	return os.RemoveAll(path)
}

// Copy copy file content
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

// FileEmpty Is the file content empty
// true: file content is empty
// false: file content is not empty
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

// FileMd5Equal Determine whether the contents of the two files are equal
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

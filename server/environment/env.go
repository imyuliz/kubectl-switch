package environment

import (
	"os"
	"runtime"
)

// GetHome 获取当前用户的家目录
func GetHome() string {
	home := os.Getenv("HOME")
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	}
	return home
}

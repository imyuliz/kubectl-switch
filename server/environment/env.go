package environment

import (
	"os"
	"runtime"
)

// GetHome get user home dir
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

package server

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/vars"
)

// SubCommand default sub cmd
type SubCommand interface {
	Exec(cmd *CmdShim) error
}

// DEFAULTMETHOD  sub achieve Exec func
const DEFAULTMETHOD = "Exec"

// subCmdPools sub cmd pool
var subCmdPools = make(map[string]reflect.Type)

// RegisterSubCmd register sub cmd
func RegisterSubCmd(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	cmd := strings.ToLower(t.Name())
	vars.AddSubCmd(cmd)
	subCmdPools[t.Name()] = t
}

// GetSubCmd  get sub cmd in sub cmd pool
func GetSubCmd(subCmdName string) reflect.Type {
	return subCmdPools[subCmdName]
}

func isExist(subCmd string) (reflect.Type, bool) {
	for k, v := range subCmdPools {
		if subCmd == strings.ToLower(k) {
			return v, true
		}
	}
	return nil, false
}

// Run sub cmd
func Run(cmd *CmdShim) error {
	elem, ok := isExist(cmd.SubCmd)
	if !ok {
		return fmt.Errorf("This method: %s is not implemented", cmd.SubCmd)
	}
	obj := reflect.New(elem)
	// set params for method
	in := []reflect.Value{reflect.ValueOf(cmd)}
	// run method
	result := obj.MethodByName(DEFAULTMETHOD).Call(in)
	if len(result) > 0 && result[0].Interface() != nil {
		return result[0].Interface().(error)
	}
	return nil
}

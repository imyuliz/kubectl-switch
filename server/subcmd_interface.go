package server

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/vars"
)

// SubCommand switch自带只命令的实现
type SubCommand interface {
	Exec(cmd *Command) error
}

// DEFAULTMETHOD 子命令实现了这个方法
const DEFAULTMETHOD = "Exec"

// subCmdPools 子命令池
var subCmdPools = make(map[string]reflect.Type)

// RegisterSubCmd 注册子命令
func RegisterSubCmd(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	cmd := strings.ToLower(t.Name())
	vars.AddSubCmd(cmd)
	subCmdPools[t.Name()] = t
}

// GetSubCmd 从子命令池中获取某一个
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

// Run 执行子命令
func Run(cmd *Command) error {
	elem, ok := isExist(cmd.SubCmd)
	if !ok {
		return fmt.Errorf("This method: %s is not implemented", cmd.SubCmd)
	}
	obj := reflect.New(elem)
	//为方法设置参数
	in := []reflect.Value{reflect.ValueOf(cmd)}
	// 执行方法
	result := obj.MethodByName(DEFAULTMETHOD).Call(in)
	if len(result) > 0 && result[0].Interface() != nil {
		return result[0].Interface().(error)
	}
	//执行成功
	return nil
}

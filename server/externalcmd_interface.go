package server

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/vars"
)

// externalCmdPools 外部命令池
var externalCmdPools = make(map[string]reflect.Type)

// RegisterExternalCmd 外部命令
func RegisterExternalCmd(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	cmdLower := strings.ToLower(t.Name())
	vars.AddExternalCmd(cmdLower, t.Name())
	externalCmdPools[t.Name()] = t
}

func getExternalCmdPools() map[string]reflect.Type {
	return externalCmdPools
}

// RunExternalCmd 执行外部命令
func RunExternalCmd(cmd *Command) error {
	pool := getExternalCmdPools()
	if len(pool) <= 0 {
		return errors.New("can't find externalCmd in externalCmdPools, please check")
	}
	cmdStr := strings.ToLower(cmd.ExternalCmd)
	externalCmds := vars.GetExternalCmds()
	externalCmd := externalCmds[cmdStr]
	if strings.EqualFold(externalCmd, "") {
		return fmt.Errorf("外部命令:%s 暂未实现", cmd.ExternalCmd)
	}
	obj := reflect.New(pool[externalCmd])
	in := []reflect.Value{reflect.ValueOf(cmd)}
	result := obj.MethodByName(DEFAULTMETHOD).Call(in)
	if len(result) > 0 && result[0].Interface() != nil {
		return result[0].Interface().(error)
	}
	//执行成功
	return nil
}

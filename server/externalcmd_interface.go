package server

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server/vars"
)

// externalCmdPools external cmd pool
var externalCmdPools = make(map[string]reflect.Type)

// RegisterExternalCmd register external cmd
func RegisterExternalCmd(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	cmdLower := strings.ToLower(t.Name())
	vars.AddExternalCmd(cmdLower, t.Name())
	externalCmdPools[t.Name()] = t
}

func getExternalCmdPools() map[string]reflect.Type {
	return externalCmdPools
}

// RunExternalCmd  run external cmd
func RunExternalCmd(cmd *CmdShim) error {
	pool := getExternalCmdPools()
	if len(pool) <= 0 {
		return errors.New("not found sub cmd")
	}
	cmdStr := strings.ToLower(cmd.ExternalCmd)
	externalCmds := vars.GetExternalCmds()
	externalCmd := externalCmds[cmdStr]
	if strings.EqualFold(externalCmd, "") {
		return fmt.Errorf("not found cmd %s", cmd.ExternalCmd)
	}
	obj := reflect.New(pool[externalCmd])
	in := []reflect.Value{reflect.ValueOf(cmd)}
	result := obj.MethodByName(DEFAULTMETHOD).Call(in)
	if len(result) > 0 && result[0].Interface() != nil {
		return result[0].Interface().(error)
	}
	return nil
}

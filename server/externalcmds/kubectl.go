package externalcmds

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/yulibaozi/kubectl-switch/server"
	"github.com/yulibaozi/kubectl-switch/server/fileutil"
)

func init() {
	server.RegisterExternalCmd((*Kubectl)(nil))
}

// Kubectl kubectl
type Kubectl struct{}

// Exec kubectl cmd exec
func (ctl *Kubectl) Exec(cmd *server.CmdShim) error {
	// config  path
	if !server.CheckConfig(server.GetConfigDir(cmd.SubCmd)) {
		return nil
	}
	configPath, err := server.GetConfigNameByClusterName(cmd.SubCmd)
	if err != nil {
		return err
	}
	configClient, err := server.GetClient(configPath)
	if err != nil {
		return err
	}
	if configClient.Host == "" {
		return fmt.Errorf("can't found %s cluster kubeconfig", cmd.SubCmd)
	}
	// kube config path
	kubePath, err := server.GetKubeConfigPath()
	if err != nil {
		return err
	}
	if exist, dir := fileutil.PathStatus(kubePath); exist && !dir {
		// fileutil.DelDir(path string)
		if err := os.Remove(kubePath); err != nil {
			return err
		}
	}
	if err := fileutil.Touch(kubePath); err != nil {
		return err
	}
	if err := fileutil.Copy(configPath, kubePath); err != nil {
		return fmt.Errorf("cp config file failed. err:%v, src:%s, des:%s", err, configPath, kubePath)
	}
	// if exist, dir := fileutil.PathStatus(kubePath); !exist || (exist && dir) {
	// 	err := fileutil.Touch(kubePath)
	// 	if err != nil {
	// 		return fmt.Errorf("touch path failed. err:%v, path:%s", err, kubePath)
	// 	}
	// }
	// kubeClient, err := server.GetClient(kubePath)
	// if err != nil || !strings.EqualFold(kubeClient.Host, configClient.Host) {
	// 	if err := fileutil.Copy(configPath, kubePath); err != nil {
	// 		return fmt.Errorf("cp config file failed. err:%v, src:%s, des:%s", err, configPath, kubePath)
	// 	}
	// 	return nil
	// }
	ctl.execKubectl(cmd)
	return nil

	// clusterPath := server.GetConfigDir(cmd.SubCmd)
	// if !server.CheckConfig(clusterPath) {
	// 	return nil
	// }
	// configPath, err := server.GetConfigNameByClusterName(cmd.SubCmd)
	// if err != nil {
	// 	return err
	// }
	// configClient, err := server.GetClient(configPath)
	// if err != nil {
	// 	empty, err := fileutil.FileEmpty(configPath)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if empty {

	// 	}
	// }
	// equal := strings.EqualFold(kubeClient.Host, configClient.Host)
	// if equal {
	// 	empty, err := fileutil.FileEmpty(configPath)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if empty {
	// 		return fmt.Errorf("%s cluster config file(path:%s) is not allowed to be empty. ", cmd.SubCmd, configPath)
	// 	}
	// }
	// if !equal {
	// 	if err := fileutil.Copy(configPath, kubePath); err != nil {
	// 		return fmt.Errorf("cp config file failed. err:%v, src:%s, des:%s", err, configPath, kubePath)
	// 	}
	// 	empty, err := fileutil.FileEmpty(kubePath)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if empty {
	// 		return fmt.Errorf("config file is not allowed to be empty. path:%s", kubePath)
	// 	}
	// }

}

func (ctl *Kubectl) execKubectl(cmd *server.CmdShim) {
	if len(cmd.Args) <= 0 {
		return
	}
	if len(cmd.Flags) > 0 {
		cmd.Args = append(cmd.Args, cmd.Flags...)
	}
	var execCmd = exec.Command("kubectl", cmd.Args...)
	out, err := execCmd.CombinedOutput()
	failed := false
	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintln(os.Stdout, text)
		if !strings.Contains(text, "WARNING") {
			failed = true
		}
	}

	if err != nil {
		fmt.Print("exec kubectl err: ", err)
		failed = true
	}
	if failed {
		os.Exit(1)
	}
}

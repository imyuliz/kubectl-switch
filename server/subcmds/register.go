package subcmds

import (
	"fmt"
	"unicode"

	"github.com/yulibaozi/kubectl-switch/server"
)

var _ server.SubCommand = &Register{}

// Register 注册方法
type Register struct{}

// 初始化向中心注册方法
func init() {
	server.RegisterSubCmd((*Register)(nil))

}

// Validation 注册方法的基本判断
func (r *Register) Validation(cmd *server.Command) bool {
	if cmd.SubCmd == "" {
		return false
	}
	if len(cmd.Args) > 0 && cmd.Args[0] == "" {
		return false
	}
	return true
}

func isWord(str string) {
	for _, v := range str {
		unicode.IsLetter(v)

	}
}

// Exec 注册方法的实现
// hictl  register           qa /root/admin.config
// hictl    qa               get pod -n yulibaozi
//   |       |                 |
//   |	     |                 |
//  cmd     subcmd            args                 flags
func (r *Register) Exec(cmd *server.Command) error {
	if r.Validation(cmd) {
		clusterName := cmd.Args[0]
		clusterNames := server.GetClusterNames()
		clusterPath := server.GetConfigDir(clusterName)
		if !clusterNames[clusterName] {
			if err := server.MKDir(clusterPath); err != nil {
				//创建数据中心目录失败
				return fmt.Errorf("注册失败，创建目录:%s  err:%v", clusterPath, err)
			}
			// 看看是否填入了config地址
			if len(cmd.Args) > 1 && cmd.Args[1] != "" {
				err := server.CopyConfig(cmd.Args[1], clusterPath)
				if err != nil {
					return err
				}
				return fmt.Errorf("集群: %s 注册成功", clusterName)
			}
			return fmt.Errorf("集群: %s 注册成功,但还需要把集群配置文件复制到PATH: %s", clusterName, clusterPath)
		}
		return fmt.Errorf("集群: %s 已经注册过了,如果需要更新,请更新path: %s 的配置文件", clusterName, clusterPath)
	}
	return fmt.Errorf("请输入正确的 register 子选项")
}

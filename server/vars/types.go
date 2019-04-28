package vars

// 定义一下基本常量
const (
	DEFAULTDIR = ".switch"
	BUFFERSIZE = 1024
	//使用的配置文件默认在 ~/.kube 目录下
	KUBECONFIGPATH = ".kube"
	KUBECONFIGFILE = "config"

	// 定义了外部命令
	KUBECTL = "kubectl"
)

// 定义一些默认操作符
var (
	// switch 自定义命令
	subCmds = map[string]bool{}

	externalCmds = map[string]string{}
)

// GetSubCmds 获取小写的子命令
func GetSubCmds() map[string]bool {
	return subCmds
}

// AddSubCmd 添加子命令
func AddSubCmd(cmd string) {
	subCmds[cmd] = true
}

// AddExternalCmd  外部命令注册
func AddExternalCmd(cmdlower, cmd string) {
	externalCmds[cmdlower] = cmd
}

// GetExternalCmds 获取现在已经实现的子命令
func GetExternalCmds() map[string]string {
	return externalCmds
}

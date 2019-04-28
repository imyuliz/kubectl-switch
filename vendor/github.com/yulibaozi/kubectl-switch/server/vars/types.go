package vars

// base const
const (
	DEFAULTDIR = ".switch"
	BUFFERSIZE = 1024
	//default config path ~/.kube
	KUBECONFIGPATH = ".kube"
	KUBECONFIGFILE = "config"

	// external command kubectl
	KUBECTL = "kubectl"
)

// default operator
var (
	subCmds      = map[string]bool{}
	externalCmds = map[string]string{} //external command pool
)

// GetSubCmds get lower sub-cmd
func GetSubCmds() map[string]bool {
	return subCmds
}

// AddSubCmd add sub cmd
func AddSubCmd(cmd string) {
	subCmds[cmd] = true
}

// AddExternalCmd  external command register
func AddExternalCmd(cmdlower, cmd string) {
	externalCmds[cmdlower] = cmd
}

// GetExternalCmds get all external sub cmds
func GetExternalCmds() map[string]string {
	return externalCmds
}

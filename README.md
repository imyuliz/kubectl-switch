# kubectl-switch

kubectl-switch是一个kubernetes命令行工具插件,它做的事情非常简单,即:你可以在一台计算机上可以操作多个kubernetes集群.

### 使用快照

这是一个小示例
 
### 安装


**`kubectl` 必须在1.12.0及以上**, 你可以使用`kubectl version`命令来查看是否满足前置条件

如果需要安装`kubectl`, 请查看:[Install and Set Up kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

```
➜  ~ kubectl version

GitVersion:"v1.14.1"
```

源码编译:

```
go get -u github.com/yulibaozi/kubectl-switch

cd $GOPATH/src/github.com/yulibaozi/kubectl-switch

go build .

# move PATH
mv  ./kubectl-switch /usr/local/bin/
```

Mac :

```
brew tap yulibaozi/tap && brew install kubectl-switch
```



### 使用流程

step 1. 当你安装完`kubectl-switch`后, 需要向`kubectl-switch`注册需要操作的集群
```
kubectl switch register      qa     /root/yulibaozi/admin.kubeconfig
#                         集群名字   集群操作所需要config文件
```

2. 查看向`kubectl-switch`注册成功的集群列表
```
kubectl switch list
```

3. 查看当前`kubectl`操作的集群
```
kubectl switch now
```
4. 切换`kubectl`操作的集群
```
kubectl switch qa
```
5. 正常执行其他kubectl命令

```
两种方式:
 
 如获取节点列表列表:

    1. kubectl get node 
    2. kubectl switch qa get node
```

### 使用惯例

1. 查看支持哪些命令
```

➜  ~ kubectl switch -h

Kubernetes multi-cluster command line management tool.

Usage:
  kubectl-switch [flags]
  kubectl-switch [command]

Available Commands:
  help        Help about any command
  list        List all cluster message
  now         View cluster of currently in use
  register    Register cluster in switch plugin
  remove      Remove the specified cluster name
  removeall   Removeall cluster config
  version     view switch plugin version

Flags:
  -h, --help     help for kubectl-switch
  -t, --toggle   Help message for toggle

Use "kubectl-switch [command] --help" for more information about a command.

```




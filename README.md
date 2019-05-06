# switch

# kubectl-switch

kubectl-switch是一个kubernetes命令行工具插件,它做的事情非常简单,即:你可以在一台计算机上可以操作多个kubernetes集群.

### 安装

源码编译:

```
go get -u github.com/yulibaozi/kubectl-switch

cd $GOPATH/src/github.com/yulibaozi/kubectl-switch

go build .

# move PATH
mv  ./kubectl-switch /usr/local/bin/
```

Mac:

```
brew install yulibaozi/tap/kubectl-switch
```

### 注意事项


### 使用示例




Kubernetes multi-cluster command line management tool

chmod -R 777 go 
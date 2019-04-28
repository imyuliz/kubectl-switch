# switch [春神] 开发步骤及细节

### 实现思路

1. 使用switch 包裹kubectl实现kubernetes多集群管理。
2. 需要提前向switch 注册集群，集群名字必须是英文字符且唯一，同时上传kubeconfig文件到switch指定目录。
3. 多集群的kubernetes的版本建议一致，因为kubectl指定的版本可能不能操作其他版本的kubernetes集权
4. 使用风格和kubectl类似，只是命令不一样，且第一个选项符是集群名字


### 优化细节

1. 如果用户不指定集群名字，默认沿用上一个对集权名字进行操作
2. 集群名字应该屏蔽kubectl已支持操作符，例如 get,delete,create,apply等
3. 可以通过执行命令，来查看我现在管理的集群列表,对正在操作的集群进行标识,这里可以使用md5来判断
4. 通过执行switch now 打印现在正在操作的kubernetes集群
5. 是否需要记录日志,什么时间，做了哪些操作，是否是危险命令，是否成功。(不记录获取命令)

### 执行约定

1. [家目录]项目使用的目录默认在HOME/switch目录下。
2.  在家目录下新建集群名字作为目录，如:HOME/kubectl-switch/gm/kube.config
3. 日志目录计划和其他日志目录一样，但是也可以指定
## 安装go-hulk工具
1. 拉取 go-hulk工具
```
go get -u git.mysre.cn/yunlian-golang/go-hulk-tool/cmd/hulk
```


2. 执行 `hulk --help`查看工具帮助文档：
```
$  hulk --help
hulk: 云链轻量级微服务框架

Usage:
   [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  new         创建应用模版。目前支持模版类型：grpc、demo
  upgrade     更新 hulk 工具

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

3. 执行 `hulk xxx --help` 查看子命令说明，比如：
```
$  hulk new --help
在当前目录创建应用模版。 示例：hulk new appname

Usage:
   new [flags]

Flags:
  -t, --appType string     可选，应用类型名，目前支持grpc、demo。默认为grpc
  -p, --apppath string     可选，应用存放路径，默认存放当前目录下
  -d, --debug              是否开启调试模式
  -u, --giturl string      可选，git地址前缀（加上后生成的项目包路径会带上git地址）
  -g, --groupname string   可选，分组名称（加上后生成的项目包路径会带上分组名）
  -h, --help               help for new
```


**注意：** 拉不下可以替换包源及私有仓库配置，详情请阅读[这个连接](https://note.youdao.com/s/FOu7YjjS)。

## 创建服务

1. 新建一个应用：
```
// 创建一个 grpc 应用
$  hulk new appname 
create a grpc app【appname】done!

// 创建一个 demo 示例应用
$  hulk new appname -t demo
create a demo app【appname】done!
```


2. 进入应用目录，执行初始化命令:
```
$  cd appname
$  make init
```
3. 在 api 目录下编写 protobuf 声明文件，执行如下命令：
```
$  make proto
protoc --proto_path=./api \
           --go_out=paths=source_relative:./api \
           --go-grpc_out=paths=source_relative:./api \
           --grpc-gateway_out=paths=source_relative:./api \
           api/biz/user.proto
```


4. 打开 `confg/app.yaml`更改配置：
```
app:
  name: yunlian.appname
  env: dev

grpc:
  port: 9090

http:
  port: 80

mysql: 
    host: "xxx.xxx.xxx.xx"
    port: "3306"
    user: "xxxx"
    password: "xxxx"
    database: "testdb"

log:
  path: ./runtime/logs/
```
**注意：**  示例项目的表结构已经放在 ./config/demo.sql 下

5. 执行以下命令运行服务：
```
make start
```

## 目录结构说明
go-hulk 应用目录结构参考了 DDD 的分层架构，提前阅读 DDD 及 DDD 分层架构相关文章有助于理解框架的设计思想。
```
yourService
|-- api // protobuf 等接口定义文件
    |-- google // google api 接口定义目录
    |-- biz // 业务接口定义目录
|-- cmd
    |-- server
        |-- main.go // 服务运行入口文件
|-- config
    |-- app.yaml 
|-- internal
    |-- application // 应用层
        |-- service.go // 应用服务，自身不包含业务逻辑，通过调度用例及领域服务间接实现业务
    |-- domain // 领域层
        |-- demo 
            |-- demo.go // 存放实体，值对象
            |-- repo.go // 定义 repo 接口
    |-- common // 基础层 
        |-- utils
        |-- logger
        |-- repo
            |-- demo.go // 实现领域层定义的repo 接口
    |-- interfaces // 用户接口层
        |-- assembler
        |-- dto
        |-- server.go // 实现 grpc、http 服务的注册
|-- test
    |-- README.md
|-- vendor
|-- go.mod
|-- go.sum
|-- README.md
|-- CHANGELOG.md

```
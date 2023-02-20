# TikDuck 一个极简版抖音的后端

[![codecov](https://codecov.io/gh/808-not-found/tik_duck/branch/master/graph/badge.svg?token=ZRGZZ0HMMG)](https://codecov.io/gh/808-not-found/tik_duck)

## 项目工具依赖

- [Go 1.18](https://go.dev/)
- [go-task](https://taskfile.dev/installation/)，使用 `task` 查看所有的构建目标。（部分 linux 发行版安装完成之后的命令是 `go-task`）
- [golangci-lint](https://golangci-lint.run/)，使用 `task lint` 运行 linter
- [mockery](https://github.com/vektra/mockery) 生成 mock，mock 测试的具体方案可以参考 [mock-simple-demo](https://github.com/808-not-found/mock-simple-demo)
- [kitex](https://www.cloudwego.io/docs/kitex/) 通过 IDL 生成 RPC 代码框架
- [prettier](https://prettier.io/) 非代码文件使用它进行格式化

## 开发

每次 commit 之前请确保以下命令正常运行：

- `task lint` 对代码进行静态检查
- `task test` 进行单元测试
- `task prettier` 对非代码文件进行格式化

`docker-compose up -d` 启动 MySQL 和 Etcd 服务。

分别在 `cmd/user` `cmd/useruser` `cmd/userplat` `cmd/web` 下运行 `go run .` 启动服务。

## 远程开发

本项目已经配置好 gitpod 与 github codespaces 的初始化文件，在你 fork 本仓库后，可以根据性能需求和连接稳定性选择其中之一进行远程开发，不需要在本地额外配置开发环境。**这两个远程开发平台有时间限制，退出页面之前请确保关闭工作区！**

在进入开发环境之后，请更改 go 插件设置中 `Lint Tool` 的选项为 `golangci-lint`。

### gitpod

使用 github 账号登录进 [gitpod](https://gitpod.io/workspaces) 后，点击 `New Workspace` **输入你 fork 的仓库地址**。

### github codespaces

在**你 fork 的仓库**中，点击右上角 `Code`，切换到 Codespaces 即可创建。

## 调用关系

> 当前为第一版，接口说明请参考飞书内部文档\
> 值得注意的是，尽管中文我们称他们为基础接口、互动接口、社交接口\
> 但是我们将他们分别表述为 user, user_platform, user_user

![call_relation.svg](./call_relation.svg)

## 目录结构

```
.
├── docker-compose.yaml     // 用于创建 mysql 和 Etcd 的 docker-compose
├── go.mod
├── IDLs                    // 存放 IDL 文件
├── README.md               // 本文档
├── Taskfile.yaml           // 自动命令
├── cmd
│   ├── user                // 用户接口
│   ├── userplat            // 互动接口
│   ├── useruser            // 社交接口
│   └── web                 // 网关
├── kitex_gen               // IDL 生成代码
└── pkg                     // 通用代码，如 logger
    └── initsql
```

## 数据图关系图

> 初始化 sql 语句在 `/pkg/initsql` 下

![database.png](./database.png)

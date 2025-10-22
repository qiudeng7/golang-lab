# Golang笔记

我个人对Golang很有兴趣，因为我感兴趣的devops、网络和云原生，正是golang的基本盘。

## 安装

无论是linux/mac/windows，都需要下载预编译的二进制包来安装，当然也可以源码编译，但是整个go生态的作风似乎都更倾向于直接提供二进制包，毕竟跨平台是go的一大特性，我很喜欢这个作风。go二进制包下载地址: https://go.dev/doc/manage-install

本项目的demo1在dockerfile中使用ubuntu镜像安装了golang，并且每个demo都会通过devcontainer的方式来运行go环境。具体的Dockerfile代码见demo中的Dockerfile文件。

go的安装参考 [golang官方文档: Download and install](https://go.dev/doc/install)

### 多版本管理

不像python和node需要其他第三方工具，go本身就可以管理多版本go:

1. 先执行`go install golang.org/dl/go1.10.7@latest` 这个命令会下载一个version wrapper(版本包装器)，是一个用来下载对应版本go的工具。

2. 再执行`go1.10.7 download`（需要先把`go install`的安装目录添加到PATH环境变量），这一步会使用包装器下载完整的go，包括编译器、标准库和文档。

多版本管理参考 [golang官方文档: Managing Go installations](https://go.dev/doc/manage-install)

> go常常会给我这种感觉 “这种功能居然是由官方来提供的吗”，从这个角度来看，go生态还是比较关注development experience（开发体验）的，我很欣赏。不过听说go的语法比较丑陋，我已经做好了心理准备。

---

1. 关于`go install`，这个命令用来下载一些go中的包
   1. 会下载到哪里？ 
      1. 默认安装到`GOBIN`环境变量指定的位置，这也是`GOBIN`变量的作用，如果没有设置`GOBIN`，则会安装到`$GOPATH/bin`或者`~/go/bin`
   2. 下载的是什么？
      1. `go install`命令后面的参数格式是`[build flags] [packages]`，这里`golang.org/dl/go1.10.7@latest`对应的是packages参数，packages要填是go中的包名，类似Docker镜像的命名方式，`域名/命名空间/包名@版本`
      2. `go install`命令会下载并编译这个包，然后放到指定的位置，如果那个位置被添加到了PATH变量，则可以直接执行`go1.10.7`命令来直接执行它的可执行文件。 
   3. 更多关于go install命令参考 [官方文档](https://pkg.go.dev/cmd/go#hdr-Compile_and_install_packages_and_dependencies)
2. 关于version warpper
   1. 除了下载 Go 版本，通过golang.org/dl安装的版本包装器本身就是一个功能完整的命令，它复制了标准 go命令的所有功能，但作用域仅限于其对应的特定版本。
   2. 每个版本包装器都是一个独立的可执行文件，安装在$GOPATH/bin下。它不依赖系统默认的go命令。
   3. 当你运行`go1.10.7 xxx`时，它会确保其相关的环境变量（如 GOROOT）指向自己对应的工具链目录（如 ~/sdk/go1.10.7/），从而实现完美的环境隔离
3. GO中的其他环境变量

> 下表由AI生成
> | 变量名 | 类别 | 功能描述 |
> |-------|------|----------|
> | GOROOT | 核心路径 | 指定Go语言的安装根目录 |
> | GOPATH | 核心路径 | 早期的工作区目录，其下的src、pkg、bin目录分别存放项目源码、编译的包文件和可执行文件。自Go 1.11起，模块（Module）成为标准后，其重要性下降 |
> | GOBIN | 核心路径 | go install命令编译后生成的可执行文件的存放目录 |
> | GOOS | 交叉编译 | 指定目标程序的操作系统，如linux、windows、darwin（macOS） |
> | GOARCH | 交叉编译 | 指定目标程序的处理器架构，如amd64、arm64 |
> | GOPROXY | 模块管理 | 设置模块代理镜像地址，以加速依赖下载。国内常用https://goproxy.cn |
> | GOSUMDB | 模块管理 | 指定校验和数据库，用于验证依赖包的完整性。可设置为off以关闭校验 |
> | GOPRIVATE | 模块管理 | 指示哪些模块路径是私有的（如公司内部仓库），对这些模块不进行代理拉取和校验和验证 |
> | GOMODULE | 模块管理 | 控制Go模块的开启与否。可设为on（开启）、off（关闭）或auto（自动）。在Go 1.16及以后版本中默认开启，此变量作用已减小 |
> | GOCACHE | 构建与缓存 | 指定Go编译构建过程中的缓存目录 |
> | GOMODCACHE | 构建与缓存 | 指定Go模块的缓存目录 |
> | GOTMPDIR | 构建与缓存 | 指定Go命令写入临时文件的目录 |
> | GODEBUG | 运行时与调试 | 启用运行时各种调试功能 |
> | GOMAXPROCS | 运行时与调试 | 设置应用程序可使用的CPU最大核心数 |
> | GOENV | 配置与信息 | 显示存储Go环境变量配置的文件路径。此文件通常为只读 |


# golang-lab

我个人对Golang很有兴趣，因为我感兴趣的devops、网络和云原生，正是golang的基本盘。


## 前言

我只是对这些领域感兴趣但是没有明确的目标，于是我找到了 [avelino/awesome-go](https://github.com/avelino/awesome-go) 这个仓库，像这样的awesome-xxx仓库里会有很多优秀的项目。

浏览之后发现，我有兴趣了解的项目类型有
1. 区块链
2. 聊天机器人
3. 分布式系统
4. zero trust
5. workflow

其中workflow我了解了一下代表性项目 [cadence-workflow/cadence](https://github.com/cadence-workflow/cadence)，有点类似我在py中用过的celery，是一个异步的任务编排工具。

我计划了解过go的基本机制之后着手阅读它们的源码。

### 读者说明

也许我需要向“可能的读者”说明我当前的知识基础，和[我的js-lab的读者说明](https://github.com/qiudeng7/js-lab?tab=readme-ov-file#%E8%AF%BB%E8%80%85%E8%AF%B4%E6%98%8E) 是一样的，下面是一个抄录，如果js-lab的读者说明更新了，可能这里不会及时更新。

> 我也许需要对“可能的读者”说明我当前的知识基础。这不是某种炫耀，而是方便你知道：
> 1. 如果我比你菜，那么我写的东西对你来说信息浓度偏低，你需要跳着看；
> 2. 如果你比我菜，那么我写的东西可能需要你自己去补充更多的前置知识。
> 
> 下面是一些我常用的技术栈,可能在文中不会进行额外说明。再次强调这不是炫耀，“谓余勉乡人以学者，余之志也；诋我夸际遇之盛而骄乡人者，岂知予者哉”。
> 
> 1. 容器化相关技术栈，括号内为前置知识
>     1. wsl(linux): windows中的linux子系统
>     2. devcontainer(docker): 容器化的开发环境
>     3. kind(k8s): 让容器成为k8s节点，开发环境秒级重建k8s。
> 2. 我比较了解的框架和库
>    1. vue(ts): 一个人尽皆知的前端开发框架。
>    2. uniapp和taro(ts): 二者都用于开发小程序，学习taro对我来说是个弯路。
>    3. django(python): 一个python的后端框架，文档质量高，学习曲线温和。
>    4. DrissionPage(python): 一个基于cdp(chrome devtools protocol)的浏览器自动化库。
> 3. 一些部署服务的经验
>    1. 云服务商的基本使用: 至少会用cdn，dns，对象存储
>    2. 简单的CICD: 至少会用 github actions，gitea
>    3. 反向代理: 至少会用 caddy和nginx
>    4. 内网穿透: 至少会用 frp
> 4. 一些很基础的js逆向经验
>    1. 如果你想学的话可以在b站搜索“志远逆向”
>    2. 老实说我从来没有真的成功逆向过什么网站，但是掌握chrome devtools相当重要。
> 5. 了解一些通用的知识
>    1. 比如某些语言的依赖管理、模块和导包机制、事件循环异步、基本的debug、基本的设计模式、基本的类型系统的知识、简单的编译原理知识
>    2. 了解自己正在用的IDE和操作系统
>    3. 简单的操作系统（计算机基础知识中的操作系统）和网络知识。

## 安装

无论是linux/mac/windows，都需要下载预编译的二进制包来安装，当然也可以源码编译，但是整个go生态的作风似乎都更倾向于直接提供二进制包，毕竟跨平台是go的一大特性，我很喜欢这个作风。go二进制包下载地址: https://go.dev/doc/manage-install

本项目的demo1在dockerfile中使用ubuntu镜像安装了golang，并且每个demo都会通过devcontainer的方式来运行go环境。具体的Dockerfile代码见demo中的Dockerfile文件。

go的安装参考 [golang官方文档: Download and install](https://go.dev/doc/install)

### 多版本管理

不像python和node需要其他第三方工具，go本身就可以管理多版本go:

1. 先执行`go install golang.org/dl/go1.10.7@latest` 这个命令会下载一个version wrapper(版本包装器)，是一个用来下载对应版本go的工具。

2. 再执行`go1.10.7 download`（需要先把`go install`的安装目录添加到PATH环境变量），这一步会使用包装器下载完整的go，包括编译器、标准库和文档。对于国内下载速度很慢的问题，设置环境变量`GOPROXY=https://goproxy.cn`

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

### GOPATH、GOROOT和其他环境变量

有时候执行go命令会遇到这个报错
```text
warning: both GOPATH and GOROOT are the same directory (/usr/local/go); see https://go.dev/wiki/InstallTroubleshooting
```

GOROOT​​指的是Go语言的​​安装目录​​,GOPATH用来放个人项目代码、第三方依赖包和编译后的可执行文件，一般指向	$HOME/go。

安装go的时候一般要像这样设置环境变量:
```dockerfile
ENV GOROOT=/usr/local/go 
ENV GOPATH=$HOME/go
ENV GOPROXY=https://goproxy.cn
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

下面是一些常见的GO环境变量

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
# golang module

## 创建模块、模块名、模块机制

使用这个命令新建模块:
```bash
go mod init [模块名]
```

1. mod是module的缩写；模块名需要遵循模块名规范，一个示例是`golang.org/x/repo/sub/v2`。

2. 模块名被设计为可以直接下载安装，大致就是直接在模块名前面加 http(s) 然后尝试下载。具体的寻找逻辑参考 [Finding a repository for a module path](https://go.dev/ref/mod#vcs-find)

3. 所以模块名基本就是类似URL的结构，如果不考虑会被别人下载的话，随便写就可以了。如果想被下载，比如你的代码在github上，你可以把模块名写成这样: `github.com/qiudeng7/golang-lab.git/demo2-module`，go install的时候实际上会使用git去下载仓库里的这个目录。

> 再一次感觉golang完全不像一个编程语言，而是像一个开发框架，很多设计都看起来非常的工程优先、体验有限。



所以我这里执行`mkdir module_1 && go mod init github.com/qiudeng7/golang-lab.git/demo2-module/module_1`，然后随便写个函数 就算是一个可使用的模块了。

## 调用模块

再新建一个[module_2](./module_2):
```
mkdir module_2 && go mod init github.com/qiudeng7/golang-lab.git/demo2-module/module_2
```

然后写个[main.go](./module_2/main.go)，导入module_1的时候要写全名

```go
package main

import (
	"fmt"

	"github.com/qiudeng7/golang-lab.git/demo2-module/module_1"
)
```

然后在go.mod中替换指引:
```bash
go mod edit -replace github.com/qiudeng7/golang-lab.git/demo2-module/module_1=../module_1
```

go.mod就会多一行
```go.mod
replace github.com/qiudeng7/golang-lab.git/demo2-module/module_1 => ../module_1
```

再执行`go mod tidy`，这个命令会收集当前模块内引用的全部其他模块，添加到go.mod中。详细参考 [go-mod-tidy](https://go.dev/ref/mod#go-mod-tidy)

go.mod就会多一行
```go.mod
require github.com/qiudeng7/golang-lab.git/demo2-module/module_1 v0.0.0-00010101000000-000000000000
```

然后执行`go run .`，会执行当前main包的main函数。

## 异常



## 注意点

import的是整个module，但是调用的时候可以直接调用module内具体的package名，package名是每个.go文件头部声明的名字，和文件名无关，当然最好是一致的。

也就是说和py、js不一样的是go的import语句不会直接出现具体要用的module名和函数名，而且它的module要比pkg大一级。
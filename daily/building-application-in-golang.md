# Building Applications in GoLang

为不同平台构建应用程序是目前最重要的事情之一。 Go 对如何完成有非常严格的政策。在这篇文章中，我们将学习如何构建一个 GoLang 应用程序。

## GOOS and GOARCH

GOOS 和 GOARCH 定义了当前程序的操作系统和架构。当我们输出 env 时，我们可以看到它们的设置。

```shell
$ go env GOOS GOARCH
windows
amd64
```
这些是 go build 命令适用的平台参数。

## GoLang Supported Platforms

要查看 Go 支持的所有可能平台和构建，这里有一个命令将列出所有这些。

```shell
$ go tool dist list
aix/ppc64
android/386
android/amd64
android/arm
android/arm64
darwin/amd64
darwin/arm64
dragonfly/amd64
freebsd/386
freebsd/amd64
freebsd/arm
freebsd/arm64
illumos/amd64
ios/amd64
ios/arm64
js/wasm
linux/386
linux/amd64
linux/arm
linux/arm64
linux/mips
linux/mips64
linux/mips64le
linux/mipsle
linux/ppc64
linux/ppc64le
linux/riscv64
linux/s390x
netbsd/386
netbsd/amd64
netbsd/arm
netbsd/arm64
openbsd/386
openbsd/amd64
openbsd/arm
openbsd/arm64
openbsd/mips64
plan9/386
plan9/amd64
plan9/arm
solaris/amd64
windows/386
windows/amd64
windows/arm
windows/arm64
```

## Building for a custom OS and Arch

要为自定义Os和ARCH 构建，我们需要使用构建标签。这里有些例子。比如说，我们想为 `linux/386` 创建一个构建，我们可以使用这样的构建标签来实现。

```shell
// + build linux,386
// ...file source
```

这只是文件前的注释。 

现在，如果我们运行 go build 它只会在构建标签（build tag）匹配时将此文件编译。

这是一个例子： 在一个文件中 include.go 包含一个 int ，只有在平台不是 windows 时才会被编译。现在，如果我们在 main.go 文件中使用它会发生什么。

```go
// main.go contents
package main
 
import (
    "fmt"
)
 
func main() {
    fmt.Println("Hello world")
    fmt.Println(AnInt)
}
 
// include.go contents
// +build !windows
package main
 
var AnInt = 42
```

现在，我们可以创建一个特定于 Windows 的文件来为其构建，包括该整数或其他任何内容。构建系统允许为不同平台和架构进行模块化构建，这些构建可以轻松完成，无需太多麻烦。这就是构建工具的亮点。

## Reference
[Building Applications in GoLang](https://golangdocs.com/building-applications-in-golang)


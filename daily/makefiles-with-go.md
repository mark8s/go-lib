# Makefiles with Go

当你经常在涉及 Golang 的大型项目上工作时，重复输入构建和测试命令 go build xx、go test xx 通常会很耗时。

Makefiles 提供了一种非常有效的方法来为我们的应用程序自动编写任务。

让我们了解如何使用 makefile 轻松设置工作流程！

## Using Makefiles in Go

makefile 的主要功能用例是使用“标签”轻松运行不同的任务。使用标签，make TAGNAME 将简单地运行与 TAGNAME 对应的任务。 所以让我们假设一个示例程序 main.go。如果你想构建这个文件，你通常会运行这个命令：

```shell
go build main.go
```

如果你想直接执行程序，你可以这样做：

```shell
go build -o main.out main.go
./main.out
```

但是，使用单个 makefile，您可以创建 2 个任务来构建和运行。

```shell
BINARY_NAME=main.out
 
build:
    go build -o ${BINARY_NAME} main.go
 
run:
    go build -o ${BINARY_NAME} main.go
    ./${BINARY_NAME}
 
clean:
    go clean
    rm ${BINARY_NAME}
```

现在，您可以像这样运行构建和运行任务：

```shell
make build
make run
```

如果你想添加一个默认标签来调用必要的任务，我们可以使用 all 标签：

```shell
BINARY_NAME=main.out
 
all: build test
 
build:
    go build -o ${BINARY_NAME} main.go
 
test:
    go test -v main.go
 
run:
    go build -o ${BINARY_NAME} main.go
    ./${BINARY_NAME}
 
clean:
    go clean
    rm ${BINARY_NAME}
```

这将在我们运行时构建和测试我们的 main.go 程序：

```shell
make
```

看看我们的工作现在变得多么容易？这就是为什么这些构建工具链非常有用！

请注意，您可以在 makefile 中使用 shell 变量。变量 BINARY_NAME 实际上是 main.out，因此美元符号将执行变量替换。 最好的部分是你不需要停在这里。您可以根据需要运行任意数量的复杂任务。

例如，如果您的项目有多个需要使用 go get package-name 安装的依赖项，您也可以自动执行该操作！

我们可以创建一个称为DEPS的单独任务，该任务将直接安装所有相关软件包。

例如，如果我的项目需要 Gorilla websocket 库，我可以有这样的任务：

```shell
deps:
    go get github.com/gorilla/websocket
```

当我运行时：

```shell
make deps
```

所有依赖项都已安装。




## Reference

[Makefiles with Go](https://golangdocs.com/makefiles-golang)
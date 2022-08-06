# Command Line arguments in Golang

让我们开始用 Go 编写一些程序，这些程序将允许我们在运行程序时传递命令行参数。这真的很简单，并且允许我们在运行 Go 程序时传递自定义参数。我们也将创建我们的自定义参数。

## Working with command-line arguments

如何传递参数？只需将它们以空格分隔的方式与程序名称放在一起。

```shell
./program-name -a b –c=d
```

## 1. The “os” package

os 包包含 Args 数组，这是一个字符串数组，其中包含传递的所有命令行参数。让我们看看我们的第一个参数，您根本不需要传递的程序名称。

以下代码将在命令行中打印程序名称。

```go
package main
 
import (
    "os"
    "fmt"
)
 
func main() {
        // the first argument is always program name
    programName := os.Args[0]
    fmt.Println(programName)
}
```
get ready:
```shell
mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ go mod init example
go: creating new go.mod: module example
go: to add module requirements and sums:
        go mod tidy

mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ ls
example.go  go.mod

mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ go build .

mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ ls
example.exe*  example.go  go.mod
```

output:
```shell
$ ./example.exe
D:\mark\project\go-lib\daily\command-line\example.exe
```

## 2. Number of CLI arguments
我们可以轻松获取通过命令行传递的参数数量。为此，我们使用 len() 函数。 这是打印传递的参数长度的代码。

```go
package main
 
import (
    "fmt"
    "os"
)
 
func main() {
    argLength := len(os.Args[1:])
        // use fmt.Printf() to format string
    fmt.Printf("Arg length is %d", argLength) 
}
```

我们正在使用切片来获取除第一个参数之外的所有参数。

## 3. Iterating over command-line arguments passed（遍历所有传递来的命令行函数）

要遍历通过命令行传递的参数，我们将遍历 os.Args 数组。 下面，您可以看到如何做到这一点。

```go
package main
 
import (
    "fmt"
    "os"
)
 
func main() {
        // the first argument i.e. program name is excluded
    argLength := len(os.Args[1:])  
    fmt.Printf("Arg length is %d\n", argLength)
 
    for i, a := range os.Args[1:] {
        fmt.Printf("Arg %d is %s\n", i+1, a) 
    }
}
```

output:
```shell
$ ./example.exe a=1 b=2 -c=3 -d=hello world
Arg length is 5
Arg 1 is a=1
Arg 2 is b=2
Arg 3 is -c=3
Arg 4 is -d=hello
Arg 5 is world
```

## Creating our custom arguments
现在，我们将创建flags来实现自定义 CLI 参数。

## 4. The “flag” package
为了实现标志，我们将使用一个名为 flag 的包。我们将创建一个虚假登录机制作为示例。 以下是虚假登录系统的代码。

以下是虚假登录系统的代码。

```shell
package main
 
import (
    "flag"
    "fmt"
)
 
func main() {
    // variables declaration  
    var uname string    
    var pass string      
 
    // flags declaration using flag package
    flag.StringVar(&uname, "u", "root", "Specify username. Default is root")
    flag.StringVar(&pass, "p", "password", "Specify pass. Default is password")
 
    flag.Parse()  // after declaring flags we need to call it
 
    // check if cli params match
    if uname == "root" && pass == "password" {
        fmt.Printf("Logging in")
    } else {
        fmt.Printf("Invalid credentials!")
    }
}
```
output:
```shell
mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ ./example.exe
Logging in
mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ ./example.exe -u root
Logging in
mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ ./example.exe -u root -p 123456
Invalid credentials!
mark@LAPTOP-VH57ARI1 MINGW64 /d/mark/project/go-lib/daily/command-line (master)
$ ./example.exe -u admin
Invalid credentials!
```

## 5. Inserting wrong values
插入错误的值会给您一个错误，并为您提供这些标志的正确用法。

```shell
$ ./example.exe -c
flag provided but not defined: -c
Usage of D:\mark\project\go-lib\daily\command-line\example.exe:
  -p string
        Specify pass. Default is password (default "password")
  -u string
        Specify username. Default is root (default "root")
```

## 6. Creating custom usage

我们可以使用 flag.Usage 定义自定义使用字符串，错误使用后将显示该字符串。因此，我们可以自定义它的显示方式。

``` go
package main
 
import (
    "flag"
    "fmt"
)
 
func main() {
    var name string
    flag.StringVar(&name, "n", "admin", "Specify name. Default is admin.")
 
    flag.Usage = func() {
        fmt.Printf("Usage of our Program: \n")
        fmt.Printf("./go-project -n username\n")
        // flag.PrintDefaults()  // prints default usage
    }
    flag.Parse()
}
```

## Reference

[Command Line arguments in GoLang](https://golangdocs.com/command-line-arguments-in-golang)

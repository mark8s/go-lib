# Goroutines

## Goroutines in GoLang
Goroutines 是 Go 编程语言最重要的方面之一。它是最小的执行单元。这篇文章探讨了 goroutine 的不同方面以及如何使用它。

### What is a goroutine?
goroutine 是 GoLang 中的轻量级线程。它可以继续与主 goroutine 一起工作，从而创建并发执行。

### Goroutine syntax
创建一个 goroutine 非常简单。我们只需要在要同时运行的函数前面添加关键字“go”，它就会起作用。

```go
go FunctionName()
```

### Goroutine example
```go
package main
 
import (
    "fmt"
    "time"
)
 
func f() {
    var i int
    for i = 0; i < 5; i++ {
        fmt.Print(i, " ")
    }
}
 
func main() {
    go f()
    f()
}
```

现在，该程序将运行。但是，它可能只会打印`f()`函数，goroutine 不会被执行。这是个问题。它发生的原因是主 goroutine 不等待另一个 goroutine 完成。也即main goroutine执行完了，但是goroutine还未执行。

为了解决这个问题，我们可以让主 goroutine 休眠一会儿。这样我们就为 goroutine 提供了足够的时间来执行和完成。这是我们如何做到这一点的。
```go
package main
 
import (
    "fmt"
    "time"
)
 
func f() {
    var i int
    for i = 0; i < 5; i++ {
        // Sleep pauses the current goroutine for at least the duration d.
    	time.Sleep(10 * time.Millisecond)
        fmt.Print(i, " ")
    }
}
 
func main() {
    go f()
    f()
}
```
输出
```go
0 0 1 1 2 2 3 3 4 4
```

time.Sleep 会停止当前的goroutine，所以不管是 go f() 还是 f(),一开始进入 loop中执行逻辑的时候，都会阻塞10毫米，然后才执行打印。

### Anonymous goroutines
Go 支持匿名函数。 Goroutines 也可以是匿名的。这是一个匿名 goroutine 的示例。
```go
package main

import (
	"fmt"
	"time"
)

func PrintName(f string, l string) {
	fmt.Println(f, l)
}

func main() {
	var i int
	go func() {
		for i = 0; i < 7; i++ {
			fmt.Print(i, " ")
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Second)
	PrintName("John", "Doe")
}
```
输出
```go
0 1 2 3 4 5 6 John Doe
```
这里 time.Sleep 会 pause main goroutine 一秒钟，再这个时候 goroutine 每次pause 100毫米，所以此时会先跑完goroutine。最后执行main。

### When to use Goroutines in GoLang

- 当一项任务可以分成不同的部分以更好地执行时。 
- 当向不同的 API 端点发出多个请求时。 
- 任何可以利用多核 CPU 的工作都应该使用 goroutine 进行优化。 
- 在程序中运行后台操作可能是 goroutine 的一个用例。

### Real-Life Use Cases of Goroutines or Concurrency

- 读取巨大的日志文件并处理异常或错误消息
- 当它们不相互依赖时，在不同的线程中发布多个 API 调用
- 落实“一劳永逸”的局面
- 处理一个巨大的 SQL 文件以将数据转储到表中。

## Channels in GoLang
通道是 goroutine 用来进行有效通信的媒介。这是理解 goroutines 是如何工作的之后要掌握的最重要的概念。这篇文章旨在详细解释通道的工作原理及其在 Go 中的用例。







# Reference
[Goroutines in GoLang](https://golangdocs.com/goroutines-in-golang)


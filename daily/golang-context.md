# Golang Context Package 

## What is a GoLang context package?

GoLang 上下文包定义了 Context 类型，它携带截止日期、取消信号和其他跨 API 边界和进程之间的请求范围值。对服务器的传入请求应该创建一个上下文，对服务器的传出调用应该接受一个上下文。
它们之间的函数调用链必须传播上下文。这里有两个重要的上下文是取消和传播（Cancellation & Propagation）。

这意味着你可以要求我给你买一个三明治（你需要为此支付额外的费用），然后我去市场买。取消（Cancellation）是一种能力，如果你在中途打电话告诉我你不再需要三明治，无论我离市场多近/远，我都可以停下来。

现在考虑一下我告诉你做一个三明治来教你 Golang 的场景。然后你命令你的五个朋友去买三明治的材料。现在，如果我告诉你取消它，你不仅应该能够停止制作它，而且你还应该能够打电话给你的朋友并告诉他们放弃这个想法。这称为传播（Propagation）。

因此，现在从我们将三明治比喻为服务器的类比回来，上下文使您能够在特定时间过去后使服务器超时，从而减少负载和资源使用。

## How to create values

So, how do we do it?

让我们导入必要的包：

```go
package main
 
import (
    "context"
    "fmt"
    "log"
    "time"
)
```

上下文包中的根方法称为Background。它就像你的上下文树的根。（是的......我们正在制作一棵树 - 一个上下文树）

所有其他上下文方法都是根的孩子。要定义Background，我们使用：

```go	
c := context.Background()
```

```go
// Background returns a non-nil, empty Context. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
return background
}
```

我们可以使用 context.Context 将此上下文添加到任何函数中。假设我们有一个函数 myfunc，那么我们可以添加上下文作为参数：

```go
func myfunc(c context.Context, arg string) string {
//do something for context
     time.Sleep(2*time.Second)
     fmt.Println("Hello %s",arg)
}
```

如果经过很多时间，我们有四种方法可以让我们的上下文停止程序执行：

- context.Background
- context.WithCancel
- context.WithTimeout
- context.WithDeadline

## Adding context to a function

好吧，既然我们知道如何定义上下文值，我们必须将它传递给一个函数。例如，如果您尝试运行以下命令：

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()
	c, cancelFunc := context.WithTimeout(c, time.Second)
	defer cancelFunc()
	myFunc(c, "Gopher")
}

func myFunc(c context.Context, arg string) {
	//do something for context
	time.Sleep(2 * time.Second)
	fmt.Printf("Hello %s\n", arg)
}

```

它将打印 Hello Gopher，忽略上下文。所以我们想要同时执行两个不同的通道操作，为此，我们可以使用 select。

我们需要在 time.Sleep 的同时添加该函数，而 time 包为我们提供了一种方便的方法来做到这一点：

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Second)
	defer cancel()
	myFunc(c, 2*time.Second, "Gopher")
}

func myFunc(c context.Context, d time.Duration, arg string) {
	select {
	case <-time.After(d):
		fmt.Printf("Hello %s", arg)
	case <-c.Done():
		log.Print(c.Err())
	}
}
```

time.After(d) 将等待 d 个时间单位，然后打印 Hello Gopher。但是如果上下文在那个时间内超时，那么它将打印

```shell
2022/08/07 11:59:42 context deadline exceeded
```

## Improving Efficiency and Minimizing Waste
如果你只是在每个 HTTP 处理程序上都这样做——如果在每个 HTTP 处理程序中你只是这样做——比如，“我要调用这个函数，但如果发生这种情况，只需取消它”，如果用户只是发送请求并取消请求，因为 TCP 连接断开，该上下文将被取消。因此，带宽和存储等宝贵资源可以分配到其他地方。


## Reference
[Golang Context Package](https://golangdocs.com/golang-context-package)

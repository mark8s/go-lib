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

### GoLang Channels syntax
为了使用通道，我们必须首先创建它。我们有一个非常方便的函数 make 可以用来创建通道。通道取决于它携带的数据类型。这意味着我们不能通过 int 通道发送字符串。因此，我们需要创建一个特定于其目的的通道。

以下是我们创建Channel的方式。 chan 是一个关键字，用于使用 make 函数声明通道。
```go
// a channel that only carries int
ic := make(chan int)
```

要使用通道发送和接收数据，我们将使用通道运算符 `<-`

```go
ic <- 42         // send 42 to the channel
v := <-ic        // get data from the channel
```

未初始化或零值的通道为零(nil)。
```go
var ch chan int
fmt.Println(ch)    // <nil>
```

### Working with channels

现在，我们将尝试使用通道发送和接收数据。让我们从创建一个基本的 goroutine 开始，它将向通道发送数据，主goroutine从此通道中取数据。

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

在这里，在这个例子中，我们通过 goroutine 发送数据并相应地接收数据。现在，我们将尝试发送自定义数据，例如结构。

### Sending custom data via channels

自定义数据可以像任何其他数据类型一样发送。创建和使用通道时，我们需要注意在创建通道时使用正确的数据类型。这是一个通过通道发送 Person 结构的示例。

```go
package main
 
import (
    "fmt"
    // "time"
)
 
type Person struct {
    Name string
    Age  int
}
 
func SendPerson(ch chan Person, p Person) {
    ch <- p
}
 
func main() {
 
    p := Person{"John", 23}
 
    ch := make(chan Person)
 
    go SendPerson(ch, p)
 
    name := (<-ch).Name
    fmt.Println(name)
}
```

### The send and receive operation

通道操作默认是阻塞的。这意味着当我们使用任何发送或接收操作时，通道会阻塞，除非工作完成。从而允许它们同步。

### Using directional channels（使用定向channel）

通道可以是单向的。这意味着可以声明通道，以便通道只能发送或接收数据。这是channel的一个重要属性。

语法如下：
```go
ch := make(chan<- data_type)        // The channel operator is after the chan keyword
                                    // 操作符在 chan 的后面
                                    // The channel is send-only
 
ch := make(<-chan data_type)        // The channel operator is before the chan keyword
                                    // 操作符在 chan 的前面
                                    // The channel is receive-only
```

```go
package main
 
func f(ch chan<- int, v int) {
    ch <- v
}
 
func main() {
        // send-only channel
    ch := make(chan<- int)
 
    go f(ch, 42)
    go f(ch, 41)
    go f(ch, 40)
 
}
```

在上面的代码中，我们使用了一个仅发送通道的通道。这意味着数据只能发送到其中，但是当我们尝试从通道接收任何数据时，它会产生错误。

### Closing a channel

通过通道发送值后，可以关闭通道。 close 函数会执行此操作并生成一个布尔输出，然后可以使用该输出来检查它是否已关闭。channel关闭，布尔值为 true。

```go
package main

import "fmt"

func SendDataToChannel(ch chan string, s string) {
	ch <- s
	close(ch)
}

func main() {

	ch := make(chan string)

	go SendDataToChannel(ch, "Hello World!")

	// receive the second value as ok
	// that determines if the channel is closed or not
	v, ok := <-ch

	// check if closed
	if ok {
		fmt.Println("Channel closed")
	}

	fmt.Println(v) // Hello World!
}
```

### Using a loop with a channel（使用带有通道的循环）
范围循环可用于遍历通过通道发送的所有值。这是一个例子。

```go
package main

import "fmt"

func f(ch chan int, v int) {
	ch <- v
	ch <- v * 2
	ch <- v * 3
	ch <- v * 7
	close(ch)
}

func main() {

	ch := make(chan int)

	go f(ch, 2)

	for v := range ch {
		fmt.Println(v)
	}
}
```

正如我们所看到的，循环是在通道发送的所有值上完成的。程序按预期输出。发送值后也应该关闭通道。


## Multiple goroutines in GoLang
在这篇文章中，我们将看到如何使用多个 goroutine。

### Multiple Goroutines Simple Example
这段代码展示了两个 goroutine 在并发运行时如何交互。输出顺序将不会被维护，并且每次运行该程序时它可能会产生一个全新的输出。

```go
package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for _, c := range s {
		fmt.Print(string(c), " ")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// run two different goroutine
	go f("Hello")
	go f("World")

	// sleep main goroutine
	time.Sleep(1 * time.Second)
}
```

输出符合预期，看起来像这样。
```shell
W H e o r l l l d o 
```

### Nested Goroutines（嵌套Goroutines）
现在我们将看看如果我们嵌套两个 goroutine 会发生什么。 Goroutine 可以嵌套到任意深度，并且在许多情况下输出会更加随机。这是一个例子。

```go
package main
 
import (
    "fmt"
    "time"
)
 
func g(v int) {
    fmt.Println(v*2, v*v, " ")
}
 
func SpawnGoroutines(n int) {
    for i := 0; i < n; i++ {
        go g(i)
    }
}
 
func main() {
 
    go SpawnGoroutines(10)
 
    // sleep main goroutine
    time.Sleep(1 * time.Second)
}
```

输出：
```go
2 1  
18 81  
4 4  
6 9  
8 16  
10 25  
0 0  
12 36  
14 49  
16 64
```

### Communication through Channels（Channel通信）
Goroutine 可以通过通道进行通信。它是一个帮助在两个 goroutine 之间传递信息的管道。这是一个例子。

```go
package main
 
import (
    "fmt"
    "time"
)
 
func f(ch chan int) {
    for i := 0; i < 10; i++ {
        // send data to the channel
        ch <- i
    }
 
    // close the channel
    close(ch)
}
 
func g(ch chan int) {
    // loop over the data from the channel
    for v := range ch {
        fmt.Print(v, " ")
    }
}
 
func main() {
    ch := make(chan int)
 
    // send data to the channel
    go f(ch)
 
    // receive from the channel
    go g(ch)
 
    // sleep main goroutine
    time.Sleep(1 * time.Second)
}

```
输出：
```go
0 1 2 3 4 5 6 7 8 9 
```

可以看出，通道是多个 goroutine 通信的方式。

## Select statement in GoLang
GoLang select 语句类似于 switch 语句，用于多通道操作。在提供的任何案例准备好之前，此语句会阻塞。这篇文章将探讨 Go 编程语言中的 select 语句。

### GoLang Select Statement Syntax
select 语句的语法类似于 switch 语句。这真的很容易使用。

```go
select {
    case case1:
        // case 1...
    case case2:
        // case 2...
    case case3:
        // case 3...
    case case4:
        // case 4...
    default:
                // default case...
}
```

让我们看一个 select 语句的例子。

### Select Statement Example

这是一个 select 语句的示例，它显示了它是如何工作的。 Select 语句的工作方式类似于 switch 语句，但它没有具体的情况，而是使用通道发送或接收操作。

```go
package main
 
import (
    "fmt"
)
 
func g1(ch chan int) {
    ch <- 12
}
 
func g2(ch chan int) {
    ch <- 32
}
 
func main() {
 
    ch1 := make(chan int)
    ch2 := make(chan int)
 
    go g1(ch1)
    go g2(ch2)
 
    select {
    case v1 := <-ch1:
        fmt.Println("Got: ", v1)
    case v2 := <-ch2:
        fmt.Println("Got: ", v2)
    }
}
```

输出：
```go
Got:  12
````
或
```go
Got:  32
```

我们得到的输出完全取决于当时执行的内容。它只是随机的。我们无法预测输出，因为 select 的工作方式非常不同。如果所有语句都准备好执行，它会选择任何输出。

准备好执行：channel中可以取出数据

上面没有设置 default case，所以在 ch1 和 ch2 塞入了value之前，main goroutine都会被阻塞。

### The default case in select statement

如果没有其他case准备好执行，则执行默认case（default case）。它可以防止 select 阻塞主 goroutine，因为默认情况下操作是阻塞的。

```go
package main
 
import (
    "fmt"
)
 
func g1(ch chan int) {
    ch <- 42
}
 
func g2(ch chan int) {
    ch <- 43
}
 
func main() {
 
    ch1 := make(chan int)
    ch2 := make(chan int)
 
    go g1(ch1)
    go g2(ch2)
 
    select {
    case v1 := <-ch1:
        fmt.Println("Got: ", v1)
    case v2 := <-ch2:
        fmt.Println("Got: ", v2)
    default:
        fmt.Println("The default case!")
    }
    
    // 
}

```

output:

```go
The default case!
```

在上面的程序中，由于 goroutine 没有足够的时间来产生输出，所以会打印默认情况。因此，将打印默认值。

现在，我们可以尝试让主线程休眠，输出将完全不同。 

```go
// add 
time.Sleep(100 * time.Second)
```
outout:
```go
Got:  42
```

这表明 select 完全根据先发生的事情或简单地首先得到的事情来选择执行用例。

### The empty select statement
当我们在程序中使用空的 select 语句时，select 语句将永远阻塞，因为没有 goroutine 可用于提供任何数据。因此，主 goroutine 引发了恐慌并停止了死锁。

```go
package main
 
func main() {
    select {}
}
```

The output becomes something like this:
```shell
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [select (no cases)]:
main.main()
	D:/mark/project/go-lib/goroutines/select/demo3/example.go:4 +0x17

Process finished with the exit code 2
```

### Uses of the select statement in Go（Go中select的使用）

当多个 goroutine 通过通道发送数据时使用 select 语句，然后 select 语句同时接收数据并在所有准备就绪时随机选择 case。
如果没有case准备好，那么如果之前已经提供了default case，它只会输出default case。

这显示了 select 语句的多功能性，该语句**用于有选择地从多个提供程序通道获取数据**。

## Waitgroups in GoLang
在 Go 中进行并发编程时，可以看到 goroutine 可能没有执行，主线程在执行完之前就停止了。这是做并发编程时经常出现的问题。 Waitgroups 以一种简单的方式解决了这个问题。

### How GoLang waitgroups work?
Waitgroup 是一种阻塞机制，当该组内的所有 goroutine 都没有执行时，它会阻塞。如果一个 goroutine 完成，它就会解除对组的阻塞。

### GoLang Waitgroup example
这是一个示例，说明如何将等待组与 goroutine 一起使用。

```go
package main

import (
	"fmt"
	"sync"
)

// pass waitgroup as a pointer
func f(wg *sync.WaitGroup) {
	// do work
	fmt.Println("Working...")

	// call done
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// add to the waitgroup counter
	wg.Add(1)

	// pass waitgroup as a pointer
	go f(&wg)

	// call wait
	wg.Wait()
	fmt.Println("Done working!")
}

```

output:
```go
Working...
Done working!
```

在上面的代码中，我们首先使用了 add 函数，它告诉等待组要阻塞多少个 goroutine。然后我们简单地将组作为指向 goroutine 的指针传递。
当 goroutine 完成工作时，我们调用 Done 方法告诉等待组停止阻塞。

### Waitgroups with the anonymous function

等待组也可以与匿名函数一起使用。这真的很简单。我们不使用另一个函数通过 goroutine 调用，而是使用匿名函数并做同样的事情。

Here is an example.

```go
package main
 
import (
    "fmt"
    "sync"
)
 
func main() {
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        fmt.Println("Running anonymous function")
        wg.Done()
    }()
    wg.Wait()
    fmt.Println("Done executing")
}

```

output:
```shell
Running anonymous function
Done executing
```

可以看出，匿名函数几乎相同。

### Importance of Go waitgroups
GoLang 等待组很重要，因为它们允许 goroutine 阻塞线程并执行它。没有它，我们需要手动休眠主线程以让 goroutine 执行。
等待组也可用于不同的用例，使其成为处理 goroutine 执行的最通用方式。


# Reference
[Goroutines in GoLang](https://golangdocs.com/goroutines-in-golang)

[Channels in GoLang](https://golangdocs.com/channels-in-golang)

[Multiple goroutines in GoLang](https://golangdocs.com/multiple-goroutines-in-golang)

[Select statement in GoLang](https://golangdocs.com/select-statement-in-golang)

[Waitgroups in GoLang](https://golangdocs.com/waitgroups-in-golang)










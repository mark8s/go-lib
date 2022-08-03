# Concurrency

## Atomic Operations in GoLang – atomic package（golang原子操作-原子包）
原子操作是在硬件级别实现的操作。 Go 有 atomic 包，它有助于在进行并发时实现同步。在这篇文章中，我们将看到 atomic 包必须提供什么。

### What is an atomic operation?
如果我们在单个共享变量上使用通道进行增量或减量等操作，goroutine 将不会同步并会产生错误的输出。
另一方面，原子操作是在硬件级别实现的，这意味着当我们创建一个共享原子变量并使用多个 goroutine 来更新它的值时，它将被正确更新。

### Importing GoLang atomic package
要使用原子函数，我们需要导入 sync/atomic 包。
```go
import "sync/atomic"
```

### Atomic functions in GoLang

该包包含 int32、int 64、uint32、uint64 等的加载、存储和加法操作。因为只有 int 是可以使用原子正确同步的原语。

让我们看一个例子，如果我们不使用 atomic 会发生什么。

```go
package main
 
import (
    "fmt"
    "sync"
)
 
func f(v *int, wg *sync.WaitGroup) {
    for i := 0; i < 3000; i++ {
        *v++
    }
    wg.Done()
}
 
func main() {
    var v int = 42
    var wg sync.WaitGroup
    wg.Add(2)
    go f(&v, &wg)
    go f(&v, &wg)
    wg.Wait()
 
    fmt.Println(v)
}
```

output:
```shell
3216
```
或
```shell
6042
```
或
```shell
3720
```

可以看出，每次程序运行时都会产生错误的输出。发生这种情况是因为访问和更改值的 goroutine 不同步。
因此，更改的值是任意的，将导致错误的操作次数。


现在，我们使用 atomic 将上面的代码转换为同步代码。

```go
package main
 
import (
    "fmt"
    "sync"
    "sync/atomic"
)
 
func f(v *uint32, wg *sync.WaitGroup) {
    for i := 0; i < 3000; i++ {
        atomic.AddUint32(v, 1)
    }
    wg.Done()
}
 
func main() {
    var v uint32 = 42
    var wg sync.WaitGroup
    wg.Add(2)
    go f(&v, &wg)
    go f(&v, &wg)
    wg.Wait()
 
    fmt.Println(v)
}
```

现在，每次我们运行它，它都会产生正确的输出。由于原子操作是同步的，因此无论如何它都会返回正确的输出。

```shell
6042
```

### Use of atomic operations in GoLang
当我们需要在不同的 goroutines 之间有一个共享变量时使用原子操作，这些变量将由它们更新。如果更新操作不同步，则会产生我们看到的问题。

原子操作通过同步对共享变量的访问并使输出正确来解决该问题。

## Mutex in GoLang（互斥锁）
互斥锁是并发中的一个重要主题。当他们谈论并发时，它们几乎出现在每一种编程语言中。在这篇文章中，我们将看看互斥锁解决了哪些问题以及如何解决。

### The race condition（竞争条件）
当多个 goroutine 尝试访问和更新共享数据时，就会出现竞争条件。相反，它们无法正确更新数据并产生不正确的输出。这种情况称为竞争条件，是由于重复的线程访问而发生的。

这是一个简单的程序来说明这个问题。
```go
package main
 
import (
    "fmt"
    "sync"
)
 
func f(v *int, wg *sync.WaitGroup) {
    *v++
    wg.Done()
}
 
func main() {
 
    var wg sync.WaitGroup
    var v int = 0
 
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go f(&v, &wg)
    }
 
    wg.Wait()
    fmt.Println("Finished", v)
}
```
output:
```shell
Finished 978
```
或
```shell
Finished 985
```
或
```shell
Finished 995
```

很明显，程序每次运行时都会输出错误的结果。这里正在发生竞争情况。所以，让我们尝试用互斥锁重写这个程序。

### Working of a mutex in GoLang
在我们编写那个程序之前，让我们看看到底什么是互斥锁？简而言之，互斥锁只是一种互斥。
这意味着当多个线程访问和修改共享数据之前，它们需要获取锁。然后当工作完成时，他们释放锁并让其他一些 goroutine 获取锁。

这允许 goroutine 同步访问数据并防止数据竞争。

```shell
// working of a mutex:
 
// acquire lock
// do operation
// release lock
```

### How to use the mutex in Go
现在，我们将修改上述程序，以免发生数据竞争。我们将使用互斥锁来阻止它。这是一个关于如何去做的方法。

```go
package main
 
import (
    "fmt"
    "sync"
)
 
func f(v *int, wg *sync.WaitGroup, m *sync.Mutex) {
    // acquire lock
    m.Lock()
    // do operation
    *v++
    // release lock
    m.Unlock()
    wg.Done()
}
 
func main() {
 
    var wg sync.WaitGroup
    // declare mutex
    var m sync.Mutex
    var v int = 0
 
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go f(&v, &wg, &m)
    }
 
    wg.Wait()
    fmt.Println("Finished", v)
}
```

output:
```shell
Finished 1000
```

在这个例子中添加互斥锁后，程序输出正确。

### Mutex vs atomic package
互斥锁与原子操作非常相似，但它们比这要复杂得多，这可能会让人感到困惑。原子利用 CPU 指令，而互斥锁利用锁定机制。因此，在更新整数等共享变量时，原子更快。
但是，当同时处理复杂的数据结构时，互斥锁的真正威力就来了。然后它是唯一的选择，因为原子不支持它。

## Concurrency in GoLang
并发是指将一个大任务拆分为较小的子任务并同时运行。这篇文章概述了 Go 如何处理并发。

### Goroutines
Goroutines是GO的组成部分。它们是GO程序中执行的单位。main function也是另一个goroutine。它们都有很小的堆栈尺寸，可以同时产卵数百万。他们非常轻巧。

### Channels
通道是 goroutine 通信的媒介。它是发送和接收信息的主要媒介。通道的发送和接收操作是阻塞的。它们可以是单向或双向的，并且可以通过它们发送任何数据。

### Select statement
select 语句的作用类似于 switch 语句，但仅适用于通道。它的执行完全取决于哪种情况可以快速产生输出。如果它变得不可判定，使得多个case产生输出，那么它将随机选择case。

### Waitgroups
等待组允许等待 goroutine。因此，当执行多个 goroutine 时，它将允许它等待每个 goroutine 完成。这是一个重要的想法，使 go 并发值得。同步包包含等待组结构。

### Atomics
原子是使用 CPU 指令提供同步访问的特殊数据类型。它作为 goroutine 之间的共享变量工作。它主要用于当我们需要通过多个 goroutine 同步对变量的读写访问时。 Go 中的 sync/atomic 包包含原子。

### Mutex
互斥锁是互斥的。通过使用互斥锁，goroutine 可以并发访问任何数据。这是线程中非常常见的概念，出现在许多不同的编程语言中。 Go 有一个包含可以使用的互斥锁结构的同步包。


## Reference

[Atomic Operations in GoLang – atomic package](https://golangdocs.com/atomic-operations-in-golang-atomic-package)

[Mutex in GoLang](https://golangdocs.com/mutex-in-golang)

[Concurrency in GoLang](https://golangdocs.com/concurrency-in-golang)
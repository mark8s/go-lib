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

该包包含 int32、int 64、uint32、uint64 等的加载、存储和加法操作。因为只有 int 是可以使用原子正确同步的原语












## Reference

[Atomic Operations in GoLang – atomic package](https://golangdocs.com/atomic-operations-in-golang-atomic-package)


# GRPC
在这里，我们将了解 gRPC 是什么以及如何使用 Go 编程语言实现它。一定要准备好笔记本或文本编辑器来做笔记以备将来参考。另外，如果您有兴趣，请浏览本网站上的其他一些教程。

## Introduction to gRPC with Golang

### Why gRPC?
今天，应用程序是用多种语言编写的——一种用于前端，另一种用于数据库，另一种用于后端，一种用于 Android 应用程序，另一种用于 iOS。所以，很明显，他们都需要相互交流。要做到这一点，他们都需要遵守一组单一的 API 合约：

- 沟通渠道
- 验证 
- 有效载荷格式 
- 数据模型 
- 错误处理

所有这些都必须高效（轻便且快速）且简单。我们使用 gRPC 完成所有这些工作。

### What is gRPC?
它是最初由 Google 开发的高性能开源功能丰富的框架，现在是云原生计算基金会（或 CNCF）的一部分，就像 Kubernetes 或 Prometheus 一样。

RPC 代表远程过程调用，而 g 不代表 Google。相反，它在良好、绿色和光荣、游戏、gon 等之间发生了变化（我知道……我知道……）。

它是一种协议，允许程序执行位于另一台计算机中的另一个程序的过程。最好的一点是，开发人员不必显式编码网络交互的细节，由底层框架自动处理。所以在客户端代码中，看起来我们只是直接调用了服务端代码的一个函数。即使客户端和服务器上的代码是用不同的编程语言编写的，它也可以工作。

它的工作原理是客户端有一个存根，它提供与服务器相同的方法或功能。 gRPC 会自动为您生成存根。存根将在底层调用 gRPC 框架，通过网络与服务器交换信息。

多亏了存根，客户端和服务器现在只需要关心实现其核心服务的逻辑。我们将在后续文章中了解如何借助协议缓冲区生成 gRPC 存根。

### Code generation by gRPC
代码生成是 gRPC 最重要的特性之一。为了为服务器和客户端生成存根，我们首先需要编写 API 合约，其中包括服务的描述及其在协议缓冲区文件中的有效负载消息（它具有 .proto 扩展名）。

```go
syntax = 'proto3';
 
message HelloRequest (
    string name = 1;
)
 
message HelloResponse (
    string greet = 1;
)
 
service Welcome (
rpc Hello(HelloRequest) returns (HelloResponse);
)
```

从这个 proto 文件中，服务器和客户端存根代码由协议缓冲区编译器（或 protoc）生成。生成的 Go 代码将是这样的：

```go
//Go code
...
 
type HelloRequest struct {
     Name string
}
 
type HelloResponse struct {
     Greet string
}
 
type WelcomeServiceClient interface {
     Hello(*HelloRequest) (*HelloResponse, error)
}
 
type WelcomeServiceServer interface {
     Hello(*HelloRequest) (*HelloResponse, error)
}
```

它以二进制格式表示数据，这种格式更小，传输速度更快。序列化也比一些基于文本的格式（如 JSON 或 XML）更有效。它在客户端和服务器之间提供了一个强类型的 API 契约，使用起来非常安全。它有一套很好的 API 演进规则，以确保向后和向前兼容。

在 HTTP/2（grpc 使用的传输协议）中也可以进行多路复用，这意味着客户端和服务器可以通过单个 TCP 连接并行发送多个请求和响应。这将有助于减少延迟并提高网络利用率。

### Types of gRPC

gRPC 有 4 种类型：

- 最简单的是一元的——客户端发送 1 条请求消息，服务器回复 1 条响应

- 客户端流式传输——客户端将发送多条消息流，并且它希望服务器只发回 1 个单一响应

- 服务器流式传输 - 客户端仅发送 1 条请求消息，服务器以多条消息流进行回复

- 双向或双向流——客户端和服务器将继续以任意顺序并行发送和接收多条消息。它非常灵活且没有阻塞，这意味着在发送下一条消息之前，没有任何一方需要等待响应。

这将它与 REST 之类的服务区分开来，后者只允许从客户端到服务器的单向请求。然而，虽然所有浏览器都完全支持 REST，但对 gRPC 的支持是有限的。

gRPC 的主要用途是微服务的实现。 （微服务一般也出现在一些电子游戏中，可以在游戏中为你的游戏角色购买特殊的衣服和配饰。）

由于 gRPC 支持低延迟和高速通信，以及强大的 API 合约，因此非常适合此类快速、安全的交易。

在下一部分中，我们将看到如何实现一元 gRPC 等等。


## Reference
[Introduction to gRPC with Golang](https://golangdocs.com/grpc-golang)



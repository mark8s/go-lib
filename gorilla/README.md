# mux
gorilla/mux是 gorilla Web 开发工具包中的路由管理库。gorilla Web 开发包是 Go 语言中辅助开发 Web 服务器的工具包。它包括 Web 服务器开发的各个方面，有表单数据处理包gorilla/schema，有 websocket 通信包gorilla/websocket，有各种中间件的包gorilla/handlers，有 session 管理包gorilla/sessions，有安全的 cookie 包gorilla/securecookie。本文先介绍gorilla/mux（下文简称mux），后续文章会依次介绍上面列举的 gorilla 包。

mux有以下优势：

- 实现了标准的http.Handler接口，所以可以与net/http标准库结合使用，非常轻量；

- 可以根据请求的主机名、路径、路径前缀、协议、HTTP 首部、查询字符串和 HTTP 方法匹配处理器，还可以自定义匹配逻辑；

- 可以在主机名、路径和请求参数中使用变量，还可以为之指定一个正则表达式；

- 可以传入参数给指定的处理器让其构造出完整的 URL；

- 支持路由分组，方便管理和维护。

## 快速使用
```shell
mkdir -p gorilla/mux && cd gorilla/mux
go mod init github.com/go-lib/gorilla/gorilla/mux
```

安装gorilla/mux库:
```shell
go get -u github.com/gorilla/mux
```



## 参考
[Go 每日一库之 gorilla/mux](https://darjun.github.io/2021/07/19/godailylib/gorilla/mux/)
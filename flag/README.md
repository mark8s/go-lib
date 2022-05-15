## flag包的用法

flag用于解析命令行选项。有过类 Unix 系统使用经验的童鞋对命令行选项应该不陌生。例如命令ls -al列出当前目录下所有文件和目录的详细信息，其中-al就是命令行选项。

命令行选项在实际开发中很常用，特别是在写工具的时候。

指定配置文件的路径，如redis-server ./redis.conf以当前目录下的配置文件redis.conf启动 Redis 服务器；
自定义某些参数，如python -m SimpleHTTPServer 8080启动一个 HTTP 服务器，监听 8080 端口。如果不指定，则默认监听 8000 端口。


## 参考
[Go 每日一库之 flag](https://darjun.github.io/2020/01/10/godailylib/flag/)


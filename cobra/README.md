## cobra包的用法
cobra是一个命令行程序库，可以用来编写命令行程序。同时，它也提供了一个脚手架， 用于生成基于 cobra 的应用程序框架。非常多知名的开源项目使用了 cobra 库构建命令行，如Kubernetes、Hugo、etcd等等等等。 本文介绍 cobra 库的基本使用和一些有趣的特性。

关于作者spf13，这里多说两句。spf13 开源不少项目，而且他的开源项目质量都比较高。 相信使用过 vim 的都知道spf13-vim，号称 vim 终极配置。 可以一键配置，对于我这样的懒人来说绝对是福音。他的viper是一个完整的配置解决方案。 完美支持 JSON/TOML/YAML/HCL/envfile/Java properties 配置文件等格式，还有一些比较实用的特性，如配置热更新、多查找目录、配置保存等。 还有非常火的静态网站生成器hugo也是他的作品。

## 快速使用
```shell
go get github.com/spf13/cobra/cobra
```

```shell
mkdir scaffold
cd scaffold/
go mod init github.com/mark8s/go-lib/cobra/scaffold
cobra init 
go mod tidy
```

```shell
go get github.com/mitchellh/go-homedir
go get github.com/spf13/viper
```

增加新的命令
```shell
cobra add date
```

## 参考
[Go 每日一库之 cobra](https://darjun.github.io/2020/01/17/godailylib/cobra/)


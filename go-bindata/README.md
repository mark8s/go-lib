# go-bindata
项目主页: https://github.com/go-bindata/go-bindata

简单说，将 任意 文件转成 go 源码。它还可以帮你把数据 压缩一下 。常用于将数据嵌入程序。

这些资源文件变成源码之后，数据储存在字节切片中 (raw byte slice)，只需要导入生成的源码，调用几个简单的函数就可访问，反正比 文件 IO 来得 简单和快 。因为是源码，也会加入编译，最后 包含在可执行文件中 ，发布时也就不再需要带着资源文件。


## 安装

```shell
go get -u github.com/go-bindata/go-bindata/...
```

## 使用

先看帮助信息

```shell
mark@LAPTOP-VH57ARI1 MINGW64 ~
$ go-bindata --help
Usage: C:\Users\mark\go\bin\go-bindata.exe [options] <input directories>

  -debug
        Do not embed the assets, but provide the embedding API. Contents will still be loaded from disk.
  -dev
        Similar to debug, but does not emit absolute paths. Expects a rootDir variable to already exist in the generated code's package.
  -fs
        Whether generate instance http.FileSystem interface code.
  -ignore value
        Regex pattern to ignore
  -mode uint
        Optional file mode override for all files.
  -modtime int
        Optional modification unix timestamp override for all files.
  -nocompress
        Assets will *not* be GZIP compressed when this flag is specified.
  -nomemcopy
        Use a .rodata hack to get rid of unnecessary memcopies. Refer to the documentation to see what implications this carries.
  -nometadata
        Assets will not preserve size, mode, and modtime info.
  -o string
        Optional name of the output file to be generated. (default "./bindata.go")
  -pkg string
        Package name to use in the generated code. (default "main")
  -prefix string
        Optional path prefix to strip off asset names.
  -tags string
        Optional set of build tags to include.
  -version
        Displays version information.

```

示例用法: `go-bindata --nocompress --nometadata -o ../pkg/vfs/assets.gen.go --pkg=vfs ./...`

-o: 指定打包后生成的go文件路径
-pkg: 指定打包后生成的go文件的包名
./...: 指定将要打包的文件，也可以是一个目录


## 自动生成
已经有固定的命令 + 参数搭配了，但是每次执行，不要说手敲麻烦又易错，就连复制粘贴都是体力活。

更不要说修改完资源容易忘掉重新执行转换。这时候就需要 go generate 和 make 出场了。

### go generate
`go generate`是go工具链自带的命令，自go1.4之后提供。

只要在某个 go 源文件开头写（注意 //go:generate 前面和中间没有任何空格，冒号是半角冒号）

```shell
// --pkg package的名称
//生成
//go:generate go-bindata --nocompress --nometadata -o ../pkg/vfs/assets.gen.go --pkg=vfs ./...
```

这之后只要执行
```shell
go generate
```
工具链就会自行扫描项目所有源码里的 //go:generate [args]... ，执行里面的 cmd args ，包括但不限于 go-bindata，任何在当前工作目录可以执行的命令，都行。

建议哪里的代码引用了资源文件，这行指令就放那个源码的开头。如果多处引用，则建议统一放程序入口。

## 示例参考

[使用go-bindata将文件编译进二进制](https://wiki.eryajf.net/pages/2bf6c3/#_1-%E5%AE%89%E8%A3%85)

















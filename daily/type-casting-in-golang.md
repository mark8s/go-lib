# Type-casting in GoLang

有时我们需要将一种数据类型转换为另一种。在这篇文章中，我们将详细介绍 Go 编程语言中的类型转换。

## Types in Go

There are many data-types in Go. The numbers, floating points, boolean and string.

- Number: int, int32, int64, uint32, uint64 etc
- Floating points: float32, float64, complex64, complex128
- Boolean: bool
- string: string

## What is type-casting?

类型转换意味着将一种类型转换为另一种类型。任何类型都可以转换为另一种类型，但这并不能保证该值将保持不变或实际上完全保留，正如我们将在本文中看到的那样。

## Type-casting syntax

一般类型转换的语法非常简单。只需使用其他类型名称作为函数来转换该值。

```shell
v := typeName(otherTypeValue)

e.g. i := int(32.987) // casting to integer
```

## Implicit type conversions

与其他语言不同，Go 不支持隐式类型转换。尽管在除数时，会根据场景进行隐式转换。所以我们需要非常小心在哪里使用什么类型。

## Basic Type-casting

任何需要更改变量类型的地方，都需要进行类型转换。有时类型转换可能不像之前讨论的那样直接。这是直接类型转换的示例。

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var a int = 42
    f := float64(a)
    fmt.Println(f)       // 42
}
```

## Conversions between string and int

有一个名为 strconv 的包，可用于在字符串和 int 值之间进行转换。下面是如何做到这一点的代码。

```go
package main
 
import (
    "fmt"
    "strconv"
)
 
func main() {
    var s string = "42"
    v, _ := strconv.Atoi(s)       // convert string to int
     
    fmt.Println(v)    // 42
     
    var i int = 42
    str := strconv.Itoa(i)        // convert int to string
     
    fmt.Println(str) // 42
}
```

##  Conversions between int and float

当 int 和 float 之间发生转换时，数字可能会失去精度。这是 int 和 float 之间的类型转换。

```go
package main
 
import (
    "fmt"
)
 
func main() {
    f := 12.34567
    i := int(f)  // loses precision
    fmt.Println(i)      // 12
     
    ii := 34
    ff := float64(ii)
     
    fmt.Println(ff)     // 34
}
```
正如您在上面的代码中所看到的，将 float 转换为 int 总是会丢失精度。

## Strings and bytes conversion

字符串是字节切片。因此，两者都可以毫不费力地相互转换。这是最简单的方法。

```go
package main
 
import (
    "fmt"
)
 
func main() {
    var s string = "Hello World"
    var b []byte = []byte(s)     // convert ty bytes
     
    fmt.Println(b)  // [72 101 108 108 111 32 87 111 114 108 100]
     
    ss := string(b)              // convert to string
     
    fmt.Println(ss)     // Hello World
}
```

## Type conversion during division

在除法期间，类型被隐式转换。这里有些例子。

```go
package main
 
import (
    "fmt"
)
 
func main() {
    a := 6/3        // both are int, a is int
    f := 6.3/3      // float and int, f is float
     
    fmt.Println(a, f)     // 2 2.1
}
```

## Drawbacks of type-casting in GoLang

类型转换是改变原始数据类型的好方法。但不幸的是，在某些情况下，它会失去精度。应该谨慎行事。类型转换的需求在于 Go 不支持隐式类型转换。所以在很多情况下，应该分开做。




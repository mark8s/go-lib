# Type assertions in GoLang

Go 中的类型断言有助于访问接口的底层类型并消除接口变量的歧义。在这篇文章中，我们将更深入地了解类型断言。

## What are type-assertions in GoLang?

类型断言揭示了接口变量内部的具体值。下面是一个显示断言的示例。

```go
package main
 
import "fmt"
 
func main() {
    var i interface{} = 42
    // 
    fmt.Println(i.(int))  // 42
}
```

## Type assertion syntax

类型断言的语法很简单。我们取接口变量，然后用括号访问类型，如下所示。

```go
// declare interface and assign
var i interface{} = "a string"
 
//type-assertion
valueOfI := i.(string)    // "a string"
```

interfaceVariable.(type) 是语法。

## Why use interfaces?

接口是一种可以取任何值的类型,这会产生歧义,关于这个变量什么都说不出来。类型断言和类型切换有助于理解它包含什么类型的值以及它是什么类型。

## Checking type-assertions

要检查类型断言，我们可以使用第二个返回值来检查正确性。这是如何使用它。

```go
package main
 
import "fmt"
 
func main() {
    var i interface{} = 42
     
    v, ok := i.(string)
     
    if ok == false  {
        fmt.Println("Wrong type assertion!") 
    } else {
        fmt.Println(v)
    }
     
    // Wrong type assertion!
     
    fmt.Println(i.(int))  // 42
}
```

## Type-assertions and panic

每当发生错误的断言时，就会发生panic。panic是以前看不到的错误。

```go
package main
 
import "fmt"
 
func main() {
    var i interface{} = 42
     
    v := i.(string)
     
    fmt.Println(v)
}
```

output:

```go
panic: interface conversion: interface {} is int, not string

goroutine 1 [running]:
main.main()
	D:/mark/project/go-lib/daily/type-assertions/example.go:8 +0x2e

Process finished with the exit code 2

```

## Why is it useful?

类型断言有助于消除接口变量的歧义。它有助于解开接口变量中的具体值。















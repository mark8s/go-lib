package main

import "fmt"

func main() {
	var i interface{} = 42

	v := i.(string)

	fmt.Println(v)
}

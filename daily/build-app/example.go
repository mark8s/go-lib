// main.go contents

//go:build !windows
// +build !windows

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(AnInt)

}

// include.go contents

var AnInt = 42

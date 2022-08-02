package main

import (
	"fmt"
)

func f() {
	var i int
	for i = 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
}

func main() {
	go f()
	f()
}

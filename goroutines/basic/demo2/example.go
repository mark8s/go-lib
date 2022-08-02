package main

import (
	"fmt"
	"time"
)

func f() {
	var i int
	for i = 0; i < 5; i++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Print(i, " ")
	}
}

func main() {
	go f()
	f()
}

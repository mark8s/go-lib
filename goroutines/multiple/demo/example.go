package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for _, c := range s {
		fmt.Print(string(c), " ")
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	// run two different goroutine
	go f("Hello")
	go f("World")

	// sleep main goroutine
	time.Sleep(1 * time.Second)
}

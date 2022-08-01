package main

import (
	"fmt"
	"time"
)

func PrintName(f string, l string) {
	fmt.Println(f, l)
}

func main() {
	var i int
	go func() {
		for i = 0; i < 7; i++ {
			fmt.Print(i, " ")
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(1 * time.Second)
	PrintName("John", "Doe")
}

package main

import (
	"fmt"
)

func g1(ch chan int) {
	ch <- 12
}

func g2(ch chan int) {
	ch <- 32
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go g1(ch1)
	go g2(ch2)

	select {
	case v1 := <-ch1:
		fmt.Println("Got: ", v1)
	case v2 := <-ch2:
		fmt.Println("Got: ", v2)
	}
}

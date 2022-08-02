package main

import (
	"fmt"
	"time"
)

func f(ch chan int) {
	for i := 0; i < 10; i++ {
		// send data to the channel
		ch <- i
	}

	// close the channel
	close(ch)
}

func g(ch chan int) {
	// loop over the data from the channel
	for v := range ch {
		fmt.Print(v, " ")
	}
}

func main() {
	ch := make(chan int)

	// send data to the channel
	go f(ch)

	// receive from the channel
	go g(ch)

	// sleep main goroutine
	time.Sleep(1 * time.Second)
}

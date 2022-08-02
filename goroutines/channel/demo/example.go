package main

import (
	"fmt"
)

func SendDataToChannel(ch chan int, value int) {
	ch <- value
}

func main() {

	var v int
	ch := make(chan int) // create a channel

	go SendDataToChannel(ch, 101) // send data via a goroutine

	v = <-ch // receive data from the channel

	fmt.Println(v) // 101
}

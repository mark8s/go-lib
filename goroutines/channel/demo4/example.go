package main

import "fmt"

func SendDataToChannel(ch chan string, s string) {
	ch <- s
	close(ch)
}

func main() {

	ch := make(chan string)

	go SendDataToChannel(ch, "Hello World!")

	// receive the second value as ok
	// that determines if the channel is closed or not
	v, ok := <-ch

	// check if closed
	if ok {
		fmt.Println("Channel closed")
	}

	fmt.Println(v) // Hello World!
}

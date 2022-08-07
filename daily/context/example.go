package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()
	c, cancelFunc := context.WithTimeout(c, time.Second)
	defer cancelFunc()
	myFunc(c, "Gopher")
}

func myFunc(c context.Context, arg string) {
	time.Sleep(2 * time.Second)
	fmt.Printf("Hello %s\n", arg)
}

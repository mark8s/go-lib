package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Second)
	defer cancel()
	myFunc(c, 2*time.Second, "Gopher")
}

func myFunc(c context.Context, d time.Duration, arg string) {
	select {
	case <-time.After(d):
		fmt.Printf("Hello %s", arg)
	case <-c.Done():
		log.Print(c.Err())
	}
}

package main

import (
	"fmt"
	// "time"
)

type Person struct {
	Name string
	Age  int
}

func SendPerson(ch chan Person, p Person) {
	ch <- p
}

func main() {

	p := Person{"John", 23}

	ch := make(chan Person)

	go SendPerson(ch, p)

	name := (<-ch).Name
	fmt.Println(name)
}

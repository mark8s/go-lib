package main

import (
	"fmt"
	"sync"
)

// pass waitgroup as a pointer
func f(wg *sync.WaitGroup) {
	// do work
	fmt.Println("Working...")

	// call done
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// add to the waitgroup counter
	wg.Add(1)

	// pass waitgroup as a pointer
	go f(&wg)

	// call wait
	wg.Wait()
	fmt.Println("Done working!")
}

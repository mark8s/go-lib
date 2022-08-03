package main

import (
	"fmt"
	"sync"
)

func f(v *int, wg *sync.WaitGroup) {
	*v++
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var v int = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go f(&v, &wg)
	}

	wg.Wait()
	fmt.Println("Finished", v)
}

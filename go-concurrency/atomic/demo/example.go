package main

import (
	"fmt"
	"sync"
)

func f(v *int, wg *sync.WaitGroup) {
	for i := 0; i < 3000; i++ {
		*v++
	}
	wg.Done()
}

func main() {
	var v int = 42
	var wg sync.WaitGroup
	wg.Add(2)
	go f(&v, &wg)
	go f(&v, &wg)
	wg.Wait()

	fmt.Println(v)
}

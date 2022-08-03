package main

import (
	"fmt"
	"sync"
)

func f(v *int, wg *sync.WaitGroup, m *sync.Mutex) {
	// acquire lock
	m.Lock()
	// do operation
	*v++
	// release lock
	m.Unlock()
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	// declare mutex
	var m sync.Mutex
	var v int = 0

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go f(&v, &wg, &m)
	}

	wg.Wait()
	fmt.Println("Finished", v)
}

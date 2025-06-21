package ex8

import (
	"fmt"
	"sync"
)

var counter int

func Solve() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go increment(&wg, &mu)
	}
	wg.Wait()

	fmt.Println("global int value: %d\n", counter)
}

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	counter += 1
	mu.Unlock()
}

// The output will be less than 1000 most of the time.

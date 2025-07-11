package ex19

import (
	"fmt"
	"sync"
)

func Solve() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 5)

	for ind := 1; ind <= 5; ind++ {
		wg.Add(1)
		go worker(&wg, ch, ind)
	}

	for ind := 1; ind <= 100; ind++ {
		ch <- ind
	}
	close(ch)

	wg.Wait()
}

func worker(wg *sync.WaitGroup, ch chan int, workerID int) {
	defer wg.Done()

	for val := range ch {
		fmt.Printf("value %d picked up by worked id: %d\n", val, workerID)
	}
}

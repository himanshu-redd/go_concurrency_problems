package ex6

import (
	"fmt"
	"sync"
	"time"
)

func Solve() {
	wg := sync.WaitGroup{}

	start := time.Now()
	wg.Add(3)
	go Sleep(&wg, 5)
	go Sleep(&wg, 1)
	go Sleep(&wg, 3)
	wg.Wait()

	end := time.Now()
	fmt.Printf("time consumed: %v\n", end.Sub(start))
}

func Sleep(wg *sync.WaitGroup, seconds int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(seconds))
}

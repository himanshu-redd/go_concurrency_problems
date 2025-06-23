package ex15

import (
	"fmt"
	"sync"
	"time"
)

func Solve() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go produce(&wg, ch)
	go consume(&wg, ch)
	wg.Wait()
}

func produce(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	count := 0

	for {
		ch <- count
		count++
	}
}

func consume(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for {
		select {
		case val := <-ch:
			fmt.Println("val: ", val)
		default: 
			time.Sleep(time.Duration(200) * time.Millisecond)
		}
	}
}

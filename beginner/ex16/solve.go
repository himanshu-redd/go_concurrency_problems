package ex16

import (
	"fmt"
	"sync"
)

func Solve() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	done := make(chan bool)

	wg.Add(2)
	go produce(&wg, ch, done)
	go consume(&wg, ch, done)
	wg.Wait()
}

func produce(wg *sync.WaitGroup, ch chan int, done chan bool) {
	defer wg.Done()
	count := 0

	for {
		ch <- count
		count++

		if count == 100 {
			done <- true
			close(ch)
			return
		}
	}
}

func consume(wg *sync.WaitGroup, ch chan int, done chan bool) {
	defer wg.Done()

	for {
		select {
		case val := <-ch:
			fmt.Println("val: ", val)

		case <-done:
			fmt.Println("closing chan")
			return
		}
	}
}

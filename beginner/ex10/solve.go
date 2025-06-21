package ex10

import (
	"fmt"
	"sync"
)

func Solve() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)

	go produce(&wg, ch, 1)
	go consume(&wg, ch)

	wg.Wait()
}

func produce(wg *sync.WaitGroup, ch chan int, val int) {
	defer wg.Done()
	ch <- 1
}

func consume(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	select{
	case val := <- ch:
		fmt.Println("chan value received: ", val)
	}
}
package ex3 

import (
	"fmt"
	"sync"
)

func Solve() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go produce(&wg, ch)
	go consume(&wg, ch)

	wg.Wait()
}

func produce(wg *sync.WaitGroup, ch chan int) {
	ch <- 1234
	wg.Done()
}

func consume (wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	val, ok  := <- ch
	if ok {
		fmt.Println("chan value: ", val)
	}
}
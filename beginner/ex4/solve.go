package main

import (
	"fmt"
	"sync"
)

func ex4(){
	ch := make(chan int, 3)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go produce4(wg, ch)
	go consume4(wg, ch)

	wg.Wait()
}

func produce4(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func consume4(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("val : ", val)
	}
}
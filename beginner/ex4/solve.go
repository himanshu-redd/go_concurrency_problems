package ex4

import (
	"fmt"
	"sync"
)

func Solve(){
	ch := make(chan int, 3)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go produce(wg, ch)
	go consume(wg, ch)

	wg.Wait()
}

func produce(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func consume(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("val : ", val)
	}
}
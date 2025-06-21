package ex11

import (
	"fmt"
	"sync"
	"time"
)

func Solve(){
	wg := sync.WaitGroup{}
	ch := make(chan int, 10)

	wg.Add(2)

	go producer(&wg, ch)
	go consumer(&wg, ch)

	wg.Wait()
}

func producer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for i := 1; i <= 100; i++ {
		fmt.Println("Sent: ", i)
		ch <- i
		if i % 10 == 0 {
			fmt.Println("current len of chan: ", len(ch))
			fmt.Println("current cap of chan: ", cap(ch))
			time.Sleep(time.Duration(100) * time.Millisecond)	
		}
	}

	close(ch)
}

func consumer(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	for val := range ch {
		fmt.Println("Receive: ", val)
	}
}
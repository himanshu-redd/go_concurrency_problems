package ex10

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
	for i := 1; i <= 10; i++ {
		fmt.Println("Sent: ", i) 
		ch <- i
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	close(ch)

}

func consume(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for val := range ch {
		fmt.Println("Received: ", val)
	}
}
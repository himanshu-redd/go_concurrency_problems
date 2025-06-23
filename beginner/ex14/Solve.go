package ex14

import (
	"fmt"
	"sync"
	"time"
)

func Solve() {
	wg := sync.WaitGroup{}
	ch1 := make(chan int)
	ch2 := make(chan int)

	wg.Add(2)
	go produce(&wg, ch1, ch2)
	go consume(&wg, ch1, ch2)
	wg.Wait()
}

func produce(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch1 <- i
		ch2 <- i * 2 
	}
	close(ch1)
	close(ch2)
}

func consume(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
	defer wg.Done()

	ch1Okay := true
	ch2Okay := true

	for ch1Okay || ch2Okay {
		select {

		case val1, ok := <-ch1:
			if ok {
			fmt.Println("value from chan 1: ", val1)
			} else {
				ch1Okay = false
			}

		case val2, ok := <-ch2:
			if ok {
				fmt.Println("value from chan 2: ", val2)
			} else {
				ch2Okay = false
			}

		default:
			fmt.Println("sleeping for 200 milliseconds  ")
			time.Sleep(time.Duration(200) * time.Millisecond)
		}
	}
}

package ex21

import (
	"fmt"
	"sync"
)

func Solve() {
	wg := sync.WaitGroup{}

	ch := make(chan string)

	wg.Add(2)

	go func(wg *sync.WaitGroup, ch chan string) {
		defer wg.Done()
		ping(ch)
	}(&wg, ch)

	go func(wg *sync.WaitGroup, ch chan string) {
		defer wg.Done()
		pong(ch)
	}(&wg, ch)

	wg.Wait()
}

func ping(ch chan string) {
	ch <- ""
	counter := 0

	for counter != 400 {
		counter++
		val, ok := <-ch

		if ok {
			fmt.Println(val)
			ch <- "PING"
		} else {
			break
		}
	}
	
	close(ch)
}

func pong(ch chan string) {
	counter := 0

	for {
		val, ok := <-ch
		counter++
		if ok {
			fmt.Println(val)
			if counter <= 400 {
				ch <- "PONG"
			}
		} else {
			break
		}
	}
}

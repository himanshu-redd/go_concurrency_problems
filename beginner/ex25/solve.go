package ex25

import (
	"fmt"
	"sync"
)

func Solve() {
	wg := sync.WaitGroup{}
	workerCnt := 100
	workerChans := make([]chan int, workerCnt)

	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		ch := initWorker(&wg, i+1)
		workerChans[i] = ch
	}

	wg.Add(2)

	producerCh := make(chan int)
	go producer(&wg, producerCh)
	go fanout(&wg, producerCh, workerChans)

	wg.Wait()
}

func producer(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- 127634 
	close(ch)
}

func fanout(wg *sync.WaitGroup, ch chan int, workerChans []chan int) {
	defer wg.Done()
	for {
		val, ok := <-ch
		if ok {
			for _, ch := range workerChans {
				ch <- val
			}
		} else {
			for _, ch := range workerChans {
				close(ch)
			}
			break
		}
	}
}

func initWorker(wg *sync.WaitGroup, workerID int) chan int {
	ch := make(chan int)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			val, ok := <-ch
			if ok {
				fmt.Println("worker id: ", workerID, ", value: ", val)
			} else {
				break
			}
		}
	}(wg)

	return ch
}

package ex25

import (
	"fmt"
	"sync"
)

func Solve() {
	wg := sync.WaitGroup{}
	workerCnt := 10
	var workerChans [10]chan int

	for i := 0; i < workerCnt; i++ {
		wg.Add(1)
		ch := initWorker(&wg, i+1)
		workerChans[i] = ch
	}

	wg.Add(2)

	producerCh := make(chan int)
	go producer(&wg, producerCh)
	go fanout(&wg, producerCh, workerChans, workerCnt)

	wg.Wait()
}

func producer(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func fanout(wg *sync.WaitGroup, ch chan int, workerChans [10]chan int, workerCnt int) {
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

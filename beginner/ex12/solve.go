package ex12

import (
	"fmt"
	"sync"
)

type Pair struct{
	Num int
	Sqr int
	ConsumerId int
}

func Solve(){
	wg := sync.WaitGroup{}
	consumerWg := sync.WaitGroup{}
	buffChan:= make(chan int, 10)
	unbuffChan := make(chan Pair)	

	wg.Add(3)
	consumerWg.Add(10)
	go produce(&wg, buffChan)
	for i := 1; i <= 10; i++ {
		go consume(&consumerWg, buffChan, unbuffChan, i)
	}

	go func(wg *sync.WaitGroup, consumerWg *sync.WaitGroup, ch chan Pair){
		wg.Done()
		consumerWg.Wait()
		close(unbuffChan)
	}(&wg, &consumerWg, unbuffChan)

	go readUnbuffChan(&wg, unbuffChan)

	wg.Wait()
}

func produce(wg *sync.WaitGroup, ch chan int){
	defer wg.Done()
	for i := 1; i <= 1000; i++ {
		ch <- i
	}
	close(ch)
}

func consume(cWg *sync.WaitGroup, bChan chan int, uChan chan Pair, consumerId int) {
	defer cWg.Done()
	for val := range bChan {
		uChan <- Pair{
			Num: val, 
			Sqr: (val * val),
			ConsumerId: consumerId,
		}
	}
}

func readUnbuffChan(wg *sync.WaitGroup, uChan chan Pair) {
	defer wg.Done()
	for val := range uChan{
		fmt.Printf("val: %d\t sqr: %d\t consumerId: %d\n", val.Num, val.Sqr, val.ConsumerId)
	}
}
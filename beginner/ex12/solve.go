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

	wg.Add(12)
	consumerWg.Add(10)
	go produce(&wg, buffChan)
	for i := 1; i <= 10; i++ {
		go consume(&wg, &consumerWg, buffChan, unbuffChan, i)
	}

	go func(consumerWg *sync.WaitGroup, ch chan Pair){
		consumerWg.Wait()
		close(unbuffChan)
	}(&consumerWg, unbuffChan)

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

func consume(wg *sync.WaitGroup, cWg *sync.WaitGroup, bChan chan int, uChan chan Pair, consumerId int) {
	defer wg.Done()
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
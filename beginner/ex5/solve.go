package ex5

import (
	"fmt"
	"sync"
)

func Solve(){
	unbufChan1 := make(chan int)
	unbufChan2 := make(chan int)
	unbufChan3 := make(chan int)
	buffChan := make(chan int, 3)

	wg := sync.WaitGroup{}
	wg.Add(4)
	go produce(&wg, unbufChan1, 1)
	go produce(&wg, unbufChan2, 2)
	go produce(&wg, unbufChan3, 3)
	go receivedUnbuffChans(&wg, buffChan, unbufChan1, unbufChan2, unbufChan3)

	wg.Add(1)
	go consume(&wg, buffChan)

	wg.Wait()

}

func produce(wg *sync.WaitGroup, ch chan int, val int) {
	defer wg.Done()
	fmt.Println("putting value in unbuffChan, val: ", val)
	ch <- val
	close(ch)
}

func receivedUnbuffChans(wg *sync.WaitGroup, buffChan chan int, unbuffChans ...chan int) {
	defer wg.Done()
	for _, ch := range unbuffChans {
		val := <- ch
		fmt.Println("received from unbuffered chan, val ", val)
		buffChan <- val
	}
	close(buffChan)
}

func consume (wg *sync.WaitGroup, buffchan chan int) {
	defer wg.Done()
	for val := range buffchan {
		fmt.Println("received from buffered chan, val:  ", val)
	}
}



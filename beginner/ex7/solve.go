package ex7

import (
	"fmt"
	"sync"

)

var counter int

func Solve(){
	wg := sync.WaitGroup{}
	wg.Add(1000)
	for i := 1; i <= 1000; i++ {
		go increment(&wg)
	}
	wg.Wait()

	fmt.Println("global int value: %d\n", counter)

}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter += 1
}

// The output will be less than 1000 most of the time. 
package ex2 

import (
	"fmt"
	"sync"
)

func Solve() {
	wg := sync.WaitGroup{}
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(wg *sync.WaitGroup){
			defer wg.Done()
			fmt.Println(i)
		}(&wg)
	}

	wg.Wait()
}

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
			printId(i+1, wg)
		}(&wg)
	}

	wg.Wait()
}

func printId(id int, wg *sync.WaitGroup) {
	fmt.Println(id)
	wg.Done()
}
package ex1

import (
	"fmt"
	"sync"
)

func Solve(){
	wg := sync.WaitGroup{}	
	wg.Add(1)
	go func(wg *sync.WaitGroup){
		fmt.Println("Hello from a goroutine!")
		wg.Done()
	}(&wg)
	wg.Wait()
}
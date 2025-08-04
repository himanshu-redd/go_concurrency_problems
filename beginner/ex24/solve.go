package ex24

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Solve(){
	var counter int32
	wg := sync.WaitGroup{}

	then := time.Now()
	for i := 1; i <= 10000000; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			counter++
		}()
	}
	wg.Wait()
	fmt.Println("Final value without atomic operation: ", counter,",total time consumed: ",  time.Since(then).Seconds())

	counter = 0
	then = time.Now()
	for i := 1; i <= 10000000; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}
	wg.Wait()
	fmt.Println("Final value with atomic operation: ", counter, ",total time consumed: ",  time.Since(then).Seconds())


	counter = 0
	then = time.Now()
	mu := sync.Mutex{}
	for i := 1; i <= 10000000; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Final value with mutex: ", counter, ",total time consumed: ",  time.Since(then).Seconds())
}
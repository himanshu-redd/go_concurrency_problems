package ex26

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Solve(){
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2 * time.Second))
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go longRunningOperation(ctx, &wg)

	wg.Wait()
}

func longRunningOperation(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	
	select {
	case <- time.After(time.Duration(3) * time.Second):
		fmt.Println("Long running operation done")

	case <- ctx.Done():
		fmt.Println("Context cancelled")
	}

	fmt.Println("Exiting long running operation")
}
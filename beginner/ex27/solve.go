package ex27

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Solve() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go expensiveOp(ctx, &wg)

	time.Sleep(time.Duration(1) * time.Second)
	cancel()

	wg.Wait()
}

func expensiveOp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(time.Duration(2) * time.Second):
		fmt.Println("expensive operation finished")

	case <-ctx.Done():
		fmt.Println("context cancelled")
	}

	fmt.Println("exiting expensive operation")
}

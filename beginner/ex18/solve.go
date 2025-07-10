package ex18

import (
	"fmt"
	"time"
)

func Solve() {
	ch := make(chan int)

	go func(ch chan int) {
		time.Sleep(time.Duration(3) * time.Second)
		ch <- 543
		close(ch)
	}(ch)

	ticker := time.Tick(time.Duration(4) * time.Second)

	for {
		select {
		case val := <-ch:
			fmt.Println("val : ", val)

		case <-ticker:
			fmt.Println("timeout received")
			return
		}
	}
}

package ex9

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Solve() {
	arr := make([]int, 10000)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(500)
	}

	chunk := 2 
	chunkSize := len(arr) / chunk

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	total := 0

	begin := time.Now()
	wg.Add(chunk)
	for i := 1; i <= chunk; i++ {
		left := (i - 1) * chunkSize
		right := left + chunkSize - 1

		go func(wg *sync.WaitGroup, mu *sync.Mutex, arr []int, left, right int, total *int) {
			defer wg.Done()
			localSum := 0
			for j := left; j <= right; j++ {
				mu.Lock()
				localSum += arr[j]
				mu.Unlock()
			}
			*total += localSum
		}(&wg, &mu, arr, left, right, &total)
	}
	wg.Wait()

	// well, It's taking 20X time asynchronously to calculate the sum than synchronously.
	fmt.Println("total sum : ", total, " time taken asynchronously: ", time.Now().Sub(begin))

	total = 0
	begin = time.Now()
	for i := 0; i < 10000; i++ {
		total += arr[i]
	}

	fmt.Println("total sum : ", total, " time taken synchronously: ", time.Now().Sub(begin))

	solveWithoutMutex(arr)
}


func solveWithoutMutex(arr []int) {
	wg := sync.WaitGroup{}
	total := 0

	chunk := 2
	ch := make(chan int, chunk)
	chunkSize := len(arr) / chunk

	begin := time.Now()

	wg.Add(chunk)

	for i := 1; i <= chunk; i++ {
		left := (i - 1) * chunkSize
		right := left + chunkSize - 1

		go func(wg *sync.WaitGroup, arr []int, left, right int, ch chan int) {
			defer wg.Done()
			partialSum := 0
			for j := left; j <= right; j++ {
				partialSum += arr[j]
			}
			ch <- partialSum
		}(&wg, arr, left, right, ch)
	}

	for i := 0; i < chunk; i++ {
		total += <-ch
	}
	wg.Wait()

	// this is the most optimized way
	fmt.Println("total sum : ", total, " time taken synchronously without mutex: ", time.Now().Sub(begin))
}

package ex22

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type Counter struct {
	Count int
	mu    sync.Mutex
}

const (
	BasePath = "./ex22/files/"
)

func Solve() {
	wg := sync.WaitGroup{}
	ch := make(chan string, 3)
	counter := &Counter{}

	files, err := ioutil.ReadDir(BasePath)
	if err != nil {
		panic(err)
	}

	fileCounter := 0
	filePaths := make([]string, 0)
	for _, file := range files {
		if !file.IsDir() {
			fileCounter++
			filePaths = append(filePaths, fmt.Sprintf("%s%s", BasePath, file.Name()))
		}
	}

	fmt.Println("total no of files: ", fileCounter)
	for _, filePath := range filePaths {
		fmt.Println("processing: ", filePath)
		wg.Add(1)
		go readFile(&wg, ch, counter, filePath, fileCounter)
	}

	for {
		content, ok := <-ch
		if ok {
			fmt.Print("\n" + content + "\n")
		} else {
			break
		}
	}

	wg.Wait()

}

func readFile(wg *sync.WaitGroup, ch chan string, counter *Counter, filePath string, noOfFiles int) {
	defer wg.Done()

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}

	ch <- string(bytes)

	counter.mu.Lock()
	counter.Count++
	counter.mu.Unlock()

	if counter.Count == noOfFiles {
		close(ch)
	}
}

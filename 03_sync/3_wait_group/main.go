package main

import (
	"fmt"
	"sync"
)

func main() {
	testWaitGroup()
}

func testWaitGroup() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(i)
		}(i, &wg)
	}
	wg.Wait()
	fmt.Println("done!")
}

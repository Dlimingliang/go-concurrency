package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	testOnce()
}

func testOnce() {
	var once sync.Once
	for i := 0; i < 5; i++ {
		go func(i int) {
			f := func() {
				fmt.Printf("i:=%d\n", i)
			}
			once.Do(f)
		}(i)
	}
	time.Sleep(1 * time.Second)
}

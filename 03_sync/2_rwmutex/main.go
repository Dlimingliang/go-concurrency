package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	m := make(map[int]int)
	mux := &sync.RWMutex{}
	go writeLoop(m, mux)
	go readLoop(1, m, mux)
	go readLoop(2, m, mux)
	go readLoop(3, m, mux)
	time.Sleep(30 * time.Millisecond)
}

func writeLoop(m map[int]int, mux *sync.RWMutex) {
	for {
		for i := 0; i < 10; i++ {
			mux.Lock()
			m[i] = i * 2
			mux.Unlock()
		}
	}
}

func readLoop(id int, m map[int]int, mux *sync.RWMutex) {
	for {
		mux.RLock()
		for k, v := range m {
			fmt.Println(id, ": ", k, "-", v)
		}
		mux.RUnlock()
	}
}

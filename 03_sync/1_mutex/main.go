package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	Counter map[string]int
	Mux     sync.Mutex
}

func main() {
	testSafeCounter()
}

func testSafeCounter() {
	key := "some key"
	counter := &SafeCounter{Counter: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go func() {
			counter.Increase(key)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(counter.Value(key))
}

func (s *SafeCounter) Increase(key string) {
	s.Mux.Lock()
	s.Counter[key]++
	s.Mux.Unlock()
}

func (s *SafeCounter) Value(key string) int {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	return s.Counter[key]
}

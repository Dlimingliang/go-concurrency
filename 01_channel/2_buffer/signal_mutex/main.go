package main

import (
	"fmt"
)

type Empty interface{}
type Semaphore chan Empty

func (s Semaphore) wait(size int) {
	e := new(Empty)
	for i := 0; i < size; i++ {
		s <- e
	}
}

func (s Semaphore) signal(size int) {
	for i := 0; i < size; i++ {
		<-s
	}
}

func (s Semaphore) Lock() {
	s.wait(1)
}

func (s Semaphore) Unlock() {
	s.signal(1)
}

func main() {
	done := make(chan bool)
	sem := make(Semaphore, 1)

	go func() {
		for i := 0; i < 10; i++ {
			sem.Lock()
			fmt.Printf("%d\n", i)
			sem.Unlock()
		}
		done <- true
	}()

	go func() {
		for i := 10; i < 20; i++ {
			sem.Lock()
			fmt.Printf("%d\n", i)
			sem.Unlock()
		}
		done <- true
	}()
	<-done
	<-done
	fmt.Println("完成")
}

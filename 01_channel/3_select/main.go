package main

import (
	"fmt"
	"time"
)

func main() {
	testSelect()
}

func testSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- i
		}
	}()

	go func() {
		for i := 10; i < 20; i++ {
			time.Sleep(1 * time.Second)
			ch2 <- i
		}
	}()

	for {
		select {
		case v := <-ch1:
			fmt.Println("ch1 receiving: ", v)
		case v := <-ch2:
			fmt.Println("ch2 receiving: ", v)
		}
	}
}

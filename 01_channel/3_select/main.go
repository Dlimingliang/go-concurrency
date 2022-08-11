package main

import (
	"fmt"
	"time"
)

func main() {
	testSelect()
}

func testSelect() {
	sum := 0
	done := make(chan bool, 2)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch1 <- i
		}
		done <- true
	}()

	go func() {
		for i := 10; i < 20; i++ {
			time.Sleep(1 * time.Second)
			ch2 <- i
		}
		done <- true
	}()

	for {
		if sum == 2 {
			break
		}
		select {
		case v := <-ch1:
			fmt.Println("ch1 receiving: ", v)
		case v := <-ch2:
			fmt.Println("ch2 receiving: ", v)
		case <-done:
			fmt.Println("ch end")
			sum++
		}
	}

}

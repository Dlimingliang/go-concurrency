package main

import (
	"fmt"
	"time"
)

func main() {
	//timeoutClose()
	ticker()
}

func ticker() {
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}

func timeoutClose() {
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

	timer := time.NewTimer(time.Second * 2)

	for {
		select {
		case v := <-ch1:
			timer.Reset(time.Second * 2)
			fmt.Println("ch1 receiving: ", v)
		case v := <-ch2:
			timer.Reset(time.Second * 2)
			fmt.Println("ch2 receiving: ", v)
		case <-timer.C:
			return
		}
	}

}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	loopReceiveChannel()
}

func loopReceiveChannel() {
	ch := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10)
			fmt.Println("putting: ", n)
			ch <- n
		}
		close(ch)
	}()
	for v := range ch {
		fmt.Println("receiving: ", v)
	}
}

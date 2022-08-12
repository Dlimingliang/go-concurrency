package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	ch := make(chan int, 10)
	go send(ch, ctx)
	consume(ch, ctx)
	fmt.Println("main done!")
}

func send(ch chan int, ctx context.Context) {
	t := time.Tick(1 * time.Second)
	for {
		select {
		case <-t:
			ch <- 1
		case <-ctx.Done():
			fmt.Println("send done!")
			return
		}
	}
}

func consume(ch chan int, ctx context.Context) {
	t := time.Tick(1 * time.Second)
	for {
		select {
		case <-t:
			fmt.Println(<-ch)
		case <-ctx.Done():
			fmt.Println("consume done!")
			return
		}
	}
}

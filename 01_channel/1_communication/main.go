package main

import "fmt"

func main() {
	//sendBeforeReceive()
	receiveBeforeSend()
}

func sendBeforeReceive() {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	value := <-ch
	fmt.Println(value)
}

func receiveBeforeSend() {
	ch := make(chan int)
	closeChannel := make(chan bool)
	go func(closeChannel chan bool) {
		value := <-ch
		fmt.Println(value)
		closeChannel <- true
	}(closeChannel)

	ch <- 1
	<-closeChannel
}

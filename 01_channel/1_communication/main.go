package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	//sendBeforeReceive()
	//receiveBeforeSend()
	//testSendAndGetData()
	//signalPattern()
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in, out chan int, prime int) {

	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

func signalPattern() {
	s := []int{1, 5, 2, 4, 3}
	done := make(chan bool)
	doSort := func(s []int) {
		sort.Ints(s)
		done <- true
	}
	i := 3
	go doSort(s[:i])
	go doSort(s[i:])
	<-done
	<-done
	fmt.Println(s)
}

func testSendAndGetData() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	//fmt.Printf("%s \n", <-ch)
	//getData(ch)
	time.Sleep(1 * time.Second)
}

func sendData(ch chan<- string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	//close(ch)
}

func getData(ch <-chan string) {
	//var input string
	//for {
	//	input = <-ch
	//	fmt.Printf("%s \n", input)
	//}
	for v := range ch {
		fmt.Printf("%s \n", v)
	}
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
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

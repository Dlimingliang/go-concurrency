package main

import "fmt"

func main() {
	generation := func(done chan interface{}, slice ...int) chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for _, v := range slice {
				select {
				case <-done:
					return
				case ch <- v:

				}
			}
		}()
		return ch
	}

	multiply := func(done chan interface{}, input chan int, multiplier int) chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for v := range input {
				select {
				case <-done:
					return
				case ch <- v * multiplier:
				}
			}
		}()
		return ch
	}

	add := func(done chan interface{}, input chan int, additive int) chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for v := range input {
				select {
				case <-done:
					return
				case ch <- v + additive:
				}
			}
		}()
		return ch
	}

	done := make(chan interface{})
	defer close(done)
	input := generation(done, 1, 2, 3, 4)
	out := multiply(done, add(done, multiply(done, input, 2), 1), 2)
	for v := range out {
		fmt.Println(v)
	}
}

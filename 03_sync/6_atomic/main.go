package main

import "sync/atomic"

func main() {

}

type Config struct {
	C atomic.Value
}

func (c *Config) Get() []int {
	return *c.C.Load().(*[]int)
}

func (c *Config) Set(n []int) {
	c.C.Store(&n)
}

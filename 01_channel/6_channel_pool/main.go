package main

import (
	"fmt"
	"time"
)

type Task struct {
	multiplyTwo func()
}

func NewTask(f func()) *Task {
	return &Task{multiplyTwo: f}
}

func (t *Task) execute() {
	t.multiplyTwo()
}

type Pool struct {
	EntryChannel chan *Task
	WorkChannel  chan *Task
	workNum      int
}

func NewPoll(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		WorkChannel:  make(chan *Task),
		workNum:      cap,
	}
}

func (p *Pool) run() {
	for i := 0; i < p.workNum; i++ {
		go p.work(i + 1)
	}
	for task := range p.EntryChannel {
		p.WorkChannel <- task
	}
}

func (p *Pool) work(workID int) {
	for task := range p.WorkChannel {
		task.execute()
		fmt.Println("work ID ", workID, " has executed")
	}
}

func main() {
	poll := NewPoll(4)

	go func() {
		for {
			task := NewTask(func() {
				fmt.Println(time.Now())
			})
			poll.EntryChannel <- task
		}

	}()
	poll.run()
}

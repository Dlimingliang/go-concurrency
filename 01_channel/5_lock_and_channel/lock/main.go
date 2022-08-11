package main

import (
	"fmt"
	"sync"
	"time"
)

var result = 4999950000
var sum int

type Task struct {
	Age int
}

type MutexPool struct {
	Mu    sync.Mutex
	Tasks []Task
}

func main() {
	//testMutex()
	testChanel()
}

func testChanel() {
	now := time.Now()
	pending, done := make(chan *Task, 100), make(chan *Task, 100)
	go func(pending chan *Task) {
		for i := 0; i < 100000; i++ {
			pending <- &Task{Age: i}
		}
	}(pending)

	for i := 0; i < 20; i++ {
		go Worker(pending, done)
	}

	for i := 0; i < 100000; i++ {
		res := <-done
		sum += res.Age
	}
	time.Sleep(2 * time.Second)
	fmt.Println("should: ", result)
	fmt.Println("actual", sum)
	fmt.Println(time.Since(now).Nanoseconds())
}

func Worker(pending, done chan *Task) {
	for {
		t := <-pending
		done <- t
	}
}

func testMutex() {
	now := time.Now()
	var tasks []Task
	for i := 0; i < 100000; i++ {
		tasks = append(tasks, Task{Age: i})
	}
	poll := MutexPool{Tasks: tasks}
	for i := 0; i < 20; i++ {
		//Worker(&poll)
		go mutexWorker(&poll)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("should: ", result)
	fmt.Println("actual", sum)
	fmt.Println(time.Since(now).Nanoseconds())
}

func mutexWorker(pool *MutexPool) {
	for {
		pool.Mu.Lock()
		// begin critical section:
		if len(pool.Tasks) == 0 {
			return
		}
		task := pool.Tasks[0]       // take the first task
		pool.Tasks = pool.Tasks[1:] // update the pool of tasks
		// end critical section
		process(task)
		pool.Mu.Unlock()
	}
}

func process(t Task) {
	sum += t.Age
}

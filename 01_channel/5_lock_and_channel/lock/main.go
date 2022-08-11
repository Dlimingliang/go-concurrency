package main

import (
	"fmt"
	"sync"
	"time"
)

var muteMap sync.Map

type Task struct {
	Num int
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
	for i := 0; i < 100000; i++ {
		go func(pending chan *Task, i int) {
			pending <- &Task{Num: i}
		}(pending, i)
	}

	for i := 0; i < 5; i++ {
		go Worker(pending, done)
	}

	for i := 0; i < 100000; i++ {
		<-done
	}
	fmt.Println(time.Since(now).Nanoseconds())
}

func Worker(pending, done chan *Task) {
	for {
		t := <-pending
		processPoint(t)
		done <- t
	}
}

func processPoint(t *Task) {
	if _, ok := muteMap.Load(t.Num); !ok {
		muteMap.Store(t.Num, true)
	} else {
		fmt.Println("已经存在", t.Num)
	}
}

func testMutex() {
	now := time.Now()
	var tasks []Task
	for i := 0; i < 100000; i++ {
		tasks = append(tasks, Task{Num: i})
	}
	poll := MutexPool{Tasks: tasks}
	for i := 0; i < 5; i++ {
		go mutexWorker(&poll)
	}
	fmt.Println(time.Since(now).Nanoseconds())
}

func mutexWorker(pool *MutexPool) {
	for {
		//pool.Mu.Lock()
		if len(pool.Tasks) == 0 {
			return
		}
		task := pool.Tasks[0]
		pool.Tasks = pool.Tasks[1:]
		//pool.Mu.Unlock()
		process(task)
	}
}

func process(t Task) {
	if _, ok := muteMap.Load(t.Num); !ok {
		muteMap.Store(t.Num, true)
	} else {
		fmt.Println("已经存在", t.Num)
	}
}

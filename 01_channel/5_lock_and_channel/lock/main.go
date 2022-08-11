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
	pending := make(chan *Task, 100)
	for i := 0; i < 10000; i++ {
		go func(pending chan *Task, i int) {
			pending <- &Task{Num: i}
		}(pending, i)
	}

	for i := 0; i < 5; i++ {
		go Worker(pending)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(muteMap)
}

func Worker(pending chan *Task) {
	for {
		t := <-pending
		processPoint(t)
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
	var tasks []Task
	for i := 0; i < 10000; i++ {
		tasks = append(tasks, Task{Num: i})
	}
	poll := MutexPool{Tasks: tasks}
	for i := 0; i < 5; i++ {
		go mutexWorker(&poll)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(muteMap)
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

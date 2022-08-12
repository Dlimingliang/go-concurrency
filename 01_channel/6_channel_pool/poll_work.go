package main

import (
	"fmt"
	"time"
)

var (
	MaxWorker = 5
	MaxQueue  = 10
	jobQueue  = make(JobQueue)
)

func main() {
	dispatcher := NewDispatcher(MaxWorker)
	dispatcher.Run()
	playLoadHandler()
	time.Sleep(1 * time.Second)
}

func playLoadHandler() {
	for i := 0; i < 10; i++ {
		task := Job{Payload: Payload{id: i}}
		jobQueue <- task
	}
}

type Payload struct {
	id int
}

func (p Payload) done() {
	fmt.Printf("%d done!\n", p.id)
}

type Job struct {
	Payload Payload
}

type JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChanel  chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChanel:  make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChanel
			select {
			case job := <-w.JobChanel:
				job.Payload.done()
			case <-w.quit:
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

//调度器
type Dispatcher struct {
	MaxWorkers int
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool, MaxWorkers: maxWorkers}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-jobQueue:
			go func(job Job) {
				jobChannel := <-d.WorkerPool
				jobChannel <- job
			}(job)
		}
	}
}

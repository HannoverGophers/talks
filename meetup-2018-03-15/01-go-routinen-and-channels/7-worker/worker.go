package main

import (
	"fmt"
	"sync"
)

type Task interface {
	Do()
}

type WorkerPool struct {
	pool chan Task
	wg   sync.WaitGroup
}

func NewWorkerPool(count int) *WorkerPool {
	wp := WorkerPool{
		pool: make(chan Task),
	}

	for i := 0; i < count; i++ {
		go wp.listenToPool()
	}
	return &wp
}

func (wp *WorkerPool) Add(task Task) {
	wp.wg.Add(1)
	wp.pool <- task
}

func (wp *WorkerPool) listenToPool() {
	for task := range wp.pool {
		task.Do()
		wp.wg.Done()
	}

	fmt.Println("Worker done!")
}

func (wp *WorkerPool) Close() {
	close(wp.pool)
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
}

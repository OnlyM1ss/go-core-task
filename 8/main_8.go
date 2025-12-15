package main

import (
	"fmt"
	"time"
)

// MyWaitGroup — кастомная wait-group на семафоре (канале), без sync.WaitGroup.
type MyWaitGroup struct {
	sem chan struct{}
	n   int
}

func NewMyWaitGroup() *MyWaitGroup {
	return &MyWaitGroup{
		sem: make(chan struct{}, 10000), // достаточно большой буфер
	}
}

func (wg *MyWaitGroup) Add(delta int) {
	wg.n += delta
}

func (wg *MyWaitGroup) Done() {
	wg.sem <- struct{}{}
}

func (wg *MyWaitGroup) Wait() {
	for i := 0; i < wg.n; i++ {
		<-wg.sem
	}
}

func main() {
	wg := NewMyWaitGroup()

	task := func(id int) {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Task done:", id)
	}

	const workers = 3
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go task(i)
	}

	wg.Wait()
	fmt.Println("All tasks done")
}

package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestMyWaitGroup(t *testing.T) {
	wg := NewMyWaitGroup()

	var completed int32
	const workers = 5

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			atomic.AddInt32(&completed, 1)
		}()
	}

	wg.Wait()

	if atomic.LoadInt32(&completed) != workers {
		t.Fatalf("expected %d completed tasks, got %d", workers, completed)
	}
}

package main

import (
	"testing"
	"time"
)

func TestRandomGeneratorProducesValues(t *testing.T) {

	out := make(chan int)
	quit := make(chan struct{})

	go randomGenerator(out, quit)

	const n = 3
	values := make([]int, 0, n)
	for i := 0; i < n; i++ {
		select {
		case v := <-out:
			values = append(values, v)
		case <-time.After(time.Second):
			t.Fatalf("timeout waiting for random value")
		}
	}
	close(quit)

	if len(values) != n {
		t.Fatalf("expected %d values, got %d", n, len(values))
	}
}

func TestRandomGeneratorClosesChannelOnQuit(t *testing.T) {
	out := make(chan int)
	quit := make(chan struct{})

	go randomGenerator(out, quit)

	close(quit)

	// ждём , чтобы генератор успел закрыть канал
	time.Sleep(100 * time.Millisecond)

	_, ok := <-out
	if ok {
		t.Fatalf("expected channel to be closed")
	}
}

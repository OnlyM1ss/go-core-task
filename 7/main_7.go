package main

import (
	"fmt"
	"sync"
)

// mergeChannels сливает N каналов в один выходной.
func mergeChannels(chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chans))
	for _, ch := range chans {
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 3; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 100; i < 103; i++ {
			ch2 <- i
		}
	}()

	go func() {
		defer close(ch3)
		for i := 200; i < 203; i++ {
			ch3 <- i
		}
	}()

	merged := mergeChannels(ch1, ch2, ch3)
	for v := range merged {
		fmt.Println("Merged value:", v)
	}
}

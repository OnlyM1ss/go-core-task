package main

import (
	"fmt"
	"math/rand"
	"time"
)

// кидает случайные числа в небуферизированный канал до закрытия
func randomGenerator(out chan<- int, quit <-chan struct{}) {
	for {
		select {
		case <-quit:
			close(out)
			return
		case out <- rand.Int():
		}
	}
}

func main() {
	out := make(chan int) // небуферизированный каал
	quit := make(chan struct{})

	go randomGenerator(out, quit)

	for i := 0; i < 5; i++ {
		fmt.Println("Random:", <-out)
	}

	close(quit)
	time.Sleep(100 * time.Millisecond)
}

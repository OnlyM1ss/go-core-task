package main

import (
	"fmt"
)

//  читает из in, преобразует в float64, возводит в куб и пишет в out.
func pipeline(in <-chan uint8, out chan<- float64) {
	for v := range in {
		f := float64(v)
		out <- f * f * f
	}
	close(out)
}

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go pipeline(in, out)

	go func() {
		for i := uint8(1); i <= 5; i++ {
			in <- i
		}
		close(in)
	}()

	for v := range out {
		fmt.Println("Result:", v)
	}
}

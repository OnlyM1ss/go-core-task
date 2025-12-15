package main

import (
	"reflect"
	"testing"
)

func TestPipelineCubes(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go pipeline(in, out)

	go func() {
		values := []uint8{1, 2, 3, 4}
		for _, v := range values {
			in <- v
		}
		close(in)
	}()

	var results []float64
	for v := range out {
		results = append(results, v)
	}

	want := []float64{1, 8, 27, 64}
	if !reflect.DeepEqual(results, want) {
		t.Fatalf("expected %v, got %v", want, results)
	}
}

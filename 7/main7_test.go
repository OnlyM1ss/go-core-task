package main

import (
	"sort"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		defer close(ch2)
		ch2 <- 10
		ch2 <- 20
	}()

	out := mergeChannels(ch1, ch2)
	var result []int
	for v := range out {
		result = append(result, v)
	}

	if len(result) != 4 {
		t.Fatalf("expected 4 values, got %d", len(result))
	}

	sorted := append([]int(nil), result...)
	sort.Ints(sorted)
	want := []int{1, 2, 10, 20}
	for i, v := range want {
		if sorted[i] != v {
			t.Fatalf("unexpected values: got %v, want %v", sorted, want)
		}
	}
}

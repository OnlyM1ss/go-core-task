package main

import (
	"reflect"
	"testing"
)

func TestIntersectIntsExample(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, inter := intersectInts(a, b)
	if !ok {
		t.Fatalf("expected true for intersection")
	}
	want := []int{64, 3}
	if !reflect.DeepEqual(inter, want) {
		t.Fatalf("expected %v, got %v", want, inter)
	}
}

func TestIntersectIntsNoIntersection(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	ok, inter := intersectInts(a, b)
	if ok {
		t.Fatalf("expected false for intersection")
	}
	if len(inter) != 0 {
		t.Fatalf("expected empty slice, got %v", inter)
	}
}

func TestIntersectIntsDuplicates(t *testing.T) {
	a := []int{1, 1, 2, 2}
	b := []int{2, 2, 1, 1}

	ok, inter := intersectInts(a, b)
	if !ok {
		t.Fatalf("expected true for intersection")
	}
	// по порядку появления во втором слайсе, без повторов
	want := []int{2, 1}
	if !reflect.DeepEqual(inter, want) {
		t.Fatalf("expected %v, got %v", want, inter)
	}
}

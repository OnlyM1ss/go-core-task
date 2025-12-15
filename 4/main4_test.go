package main

import (
	"reflect"
	"testing"
)

func TestDiffStringsBasic(t *testing.T) {
	s1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	s2 := []string{"banana", "date", "fig"}

	got := diffStrings(s1, s2)
	want := []string{"apple", "cherry", "43", "lead", "gno1"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("diffStrings() = %v, want %v", got, want)
	}
}

func TestDiffStringsEmptySecond(t *testing.T) {
	s1 := []string{"a", "b"}
	s2 := []string{}

	got := diffStrings(s1, s2)
	if !reflect.DeepEqual(got, s1) {
		t.Fatalf("expected %v, got %v", s1, got)
	}
}

func TestDiffStringsEmptyFirst(t *testing.T) {
	var s1 []string
	s2 := []string{"a", "b"}

	got := diffStrings(s1, s2)
	if len(got) != 0 {
		t.Fatalf("expected empty slice, got %v", got)
	}
}

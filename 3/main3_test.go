package main

import "testing"

func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	if v, ok := m.Get("a"); !ok || v != 1 {
		t.Fatalf("expected 1, true; got %d, %v", v, ok)
	}

	if v, ok := m.Get("b"); !ok || v != 2 {
		t.Fatalf("expected 2, true; got %d, %v", v, ok)
	}
}

func TestExistsAndRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("k", 10)

	if !m.Exists("k") {
		t.Fatalf("key k should exist")
	}

	m.Remove("k")
	if m.Exists("k") {
		t.Fatalf("key k should not exist after Remove")
	}
}

func TestCopyIndependence(t *testing.T) {
	m := NewStringIntMap()
	m.Add("x", 1)
	m.Add("y", 2)

	cp := m.Copy()

	if len(cp) != 2 || cp["x"] != 1 || cp["y"] != 2 {
		t.Fatalf("unexpected copy: %v", cp)
	}

	// изменяем оригинал
	m.Add("z", 3)
	if _, exists := cp["z"]; exists {
		t.Fatalf("copy should not contain new key z: %v", cp)
	}
}

func TestGetOnEmpty(t *testing.T) {
	m := &StringIntMap{} // без инициализации map
	if _, ok := m.Get("nope"); ok {
		t.Fatalf("expected false for non-existing key")
	}
	if m.Exists("nope") {
		t.Fatalf("expected false for non-existing key Exists")
	}
}

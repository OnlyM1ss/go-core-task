package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	if m.data == nil {
		m.data = make(map[string]int)
	}
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	if m.data == nil {
		return
	}
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	if m.data == nil {
		return map[string]int{}
	}
	cp := make(map[string]int, len(m.data))
	for k, v := range m.data {
		cp[k] = v
	}
	return cp
}

func (m *StringIntMap) Exists(key string) bool {
	if m.data == nil {
		return false
	}
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	if m.data == nil {
		return 0, false
	}
	v, ok := m.data[key]
	return v, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	fmt.Println("Exists(\"one\"):", m.Exists("one"))
	fmt.Println("Exists(\"three\"):", m.Exists("three"))

	if v, ok := m.Get("two"); ok {
		fmt.Println("Get(\"two\") =", v)
	}

	cp := m.Copy()
	fmt.Println("Copy:", cp)

	m.Remove("one")
	fmt.Println("After Remove(\"one\") Exists:", m.Exists("one"))
}

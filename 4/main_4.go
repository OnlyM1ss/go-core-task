package main

import "fmt"

// возвращает элементы, которые есть в первом слайсе, но отсутствуют во втором
func diffStrings(slice1, slice2 []string) []string {
	set := make(map[string]bool)
	for _, item := range slice2 {
		set[item] = true
	}

	var result []string
	for _, item := range slice1 {
		if !set[item] {
			result = append(result, item)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := diffStrings(slice1, slice2)
	fmt.Println(result) // [apple cherry 43 lead gno1]
}

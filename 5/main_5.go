package main

import "fmt"

func intersectInts(a, b []int) (bool, []int) {
	setA := make(map[int]struct{}, len(a))
	for _, v := range a {
		setA[v] = struct{}{}
	}
	result := make([]int, 0)
	seen := make(map[int]struct{})
	for _, v := range b {
		if _, inA := setA[v]; inA {
			if _, already := seen[v]; !already {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}
	}
	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	ok, inter := intersectInts(a, b)
	fmt.Println("Has intersection:", ok)
	fmt.Println("Intersection slice:", inter)
}

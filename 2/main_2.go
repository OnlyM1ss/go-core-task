package main

import (
	"fmt"
	"math/rand"
)

// возвращает новый слайс, содержащий только четные числа из исходного.
func sliceExample(original []int) []int {
	result := make([]int, 0, len(original))
	for _, v := range original {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

// добавляет число в конец слайса и возвращает новый слайс.
func addElements(slice []int, value int) []int {
	return append(slice, value)
}

// возвращает копию переданного слайса.
func copySlice(slice []int) []int {
	if slice == nil {
		return nil
	}
	dst := make([]int, len(slice))
	copy(dst, slice)
	return dst
}

/* removeElement удаляет элемент по индексу и возвращает новый слайс.
Если индекс некорректный, возвращает исходный слайс без паники.*/

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	// создаем новый слайс, чтобы не модифицировать исходный
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result
}

func main() {
	// 1. Создаем слайсов из 10 случайных значений
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100) // числа от 0 до 99
	}
	fmt.Println("originalSlice:", originalSlice)

	// 2. Тест
	evens := sliceExample(originalSlice)
	fmt.Println("Четные элементы:", evens)

	// 3. Тестируем
	withAdded := addElements(originalSlice, 999)
	fmt.Println("С добавленным элементом 999:", withAdded)

	// 4. Тестируем
	copied := copySlice(originalSlice)
	fmt.Println("Копия слайса:", copied)
	if len(copied) > 0 {
		originalSlice[0] = -1
	}
	fmt.Println("Измененный originalSlice:", originalSlice)
	fmt.Println("Копия после изменения originalSlice (должна не измениться):", copied)

	// 5. Тестируем
	if len(originalSlice) > 0 {
		withoutFirst := removeElement(originalSlice, 0)
		fmt.Println("Без первого элемента:", withoutFirst)
	}
	if len(originalSlice) > 2 {
		withoutMiddle := removeElement(originalSlice, 1)
		fmt.Println("Без элемента с индексом 1:", withoutMiddle)
	}
	if len(originalSlice) > 0 {
		withoutLast := removeElement(originalSlice, len(originalSlice)-1)
		fmt.Println("Без последнего элемента:", withoutLast)
	}
}

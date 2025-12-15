package main_4

import "fmt"

func difference(slice1, slice2 []string) []string {
 // Создаем map для быстрого поиска элементов второго слайса
 set := make(map[string]bool)
 for _, item := range slice2 {
  set[item] = true
 }

 // Проверяем элементы первого слайса
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

 result := difference(slice1, slice2)
 fmt.Println(result) // [apple cherry 43 lead gno1]
}
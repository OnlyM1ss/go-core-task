package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// возвращает типы переданных значений, используя fmt (reflect слишком тяжёлый)
func getAllTypes(vars ...interface{}) []string {
	types := make([]string, len(vars))
	for i, v := range vars {
		types[i] = fmt.Sprintf("%T", v)
	}
	return types
}

// преобразует всё в строку
func convertAllToString(vars ...interface{}) string {
	var result string
	for _, v := range vars {
		switch val := v.(type) {
		case int:
			result += strconv.Itoa(val)
		case float64:
			result += fmt.Sprintf("%.2f", val)
		case string:
			result += val
		case bool:
			result += strconv.FormatBool(val)
		case complex64:
			result += fmt.Sprintf("%v", val)
		default:
			result += fmt.Sprintf("%v", val)
		}
	}
	return result
}

// stringToRuneSlice преобразует строку в []RUNE
func stringToRuneSlice(s string) []rune {
	return []rune(s)
}

// hashWithSalt +++ СОЛЬ
func hashWithSalt(runes []rune, salt string) string {
	saltRunes := []rune(salt)
	mid := len(runes) / 2
	resultRunes := append(runes[:mid], append(saltRunes, runes[mid:]...)...)
	data := string(resultRunes)

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func main() {
	var numDecimal int = 42           // Десятичная
	var numOctal int = 052            // Восьмеричная
	var numHexadecimal int = 0x2A     // Шестнадцатеричная
	var pi float64 = 3.14             //  float64
	var name string = "Golang"        //  string
	var isActive bool = true          //  bool
	var complexNum complex64 = 1 + 2i // complex64

	vars := []interface{}{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}

	// 2. Определение типов через fmt %T
	types := getAllTypes(vars...)
	fmt.Println("Типы переменных:")
	for i, t := range types {
		fmt.Printf("  Переменная %d: %s\n", i+1, t)
	}

	// 3. Преобразование в строку и объединение
	combinedString := convertAllToString(vars...)
	fmt.Printf("\nОбъединённая строка: %s\n", combinedString)

	// 4. Преобразование в срез рун
	runeSlice := stringToRuneSlice(combinedString)
	fmt.Printf("Срез рун (первые 20): %q\n", runeSlice[:min(20, len(runeSlice))])

	// 5. Хэширование с солью
	salt := "go-2024"
	hashed := hashWithSalt(runeSlice, salt)
	fmt.Printf("\nХэш (SHA256): %s\n", hashed)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

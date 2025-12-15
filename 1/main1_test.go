package main

import (
	"crypto/sha256"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestGetAllTypes(t *testing.T) {
	var a int = 42
	var b float64 = 3.14
	var c string = "test"
	var d bool = true
	var e complex64 = 1 + 2i

	expected := []string{"int", "float64", "string", "bool", "complex64"}
	result := getAllTypes(a, b, c, d, e)

	//  reflect.DeepEqual для корректного сравнения срезов строк
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected types %v, got %v", expected, result)
	}
}

func TestConvertAllToString(t *testing.T) {
	var a int = 42
	var b float64 = 3.14
	var c string = "Golang"
	var d bool = true
	var e complex64 = 1 + 2i

	expected := "423.14Golangtrue(1+2i)"
	result := convertAllToString(a, b, c, d, e)

	if result != expected {
		t.Errorf("Expected string %q, got %q", expected, result)
	}
}

func TestStringToRuneSlice(t *testing.T) {
	input := "hello"
	expected := []rune{'h', 'e', 'l', 'l', 'o'}
	result := stringToRuneSlice(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected runes %v, got %v", expected, result)
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("423.14Golangtrue(1+2i)")
	salt := "go-2024"

	mid := len(runes) / 2
	saltRunes := []rune(salt)
	resultRunes := append(runes[:mid], append(saltRunes, runes[mid:]...)...)
	data := string(resultRunes)
	expectedHash := sha256.Sum256([]byte(data))
	expectedHex := hex.EncodeToString(expectedHash[:])

	result := hashWithSalt(runes, salt)

	if result != expectedHex {
		t.Errorf("Hash mismatch")
	}
}

func TestFullPipeline(t *testing.T) {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i

	vars := []interface{}{
		numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum,
	}

	types := getAllTypes(vars...)
	expectedTypes := []string{"int", "int", "int", "float64", "string", "bool", "complex64"}

	if !reflect.DeepEqual(types, expectedTypes) {
		t.Fatalf("Type list mismatch:\nExpected: %v\nGot:      %v", expectedTypes, types)
	}

	combined := convertAllToString(vars...)
	expectedCombined := "4242423.14Golangtrue(1+2i)"
	if combined != expectedCombined {
		t.Fatalf("Combined string mismatch:\nExpected: %q\nGot:      %q", expectedCombined, combined)
	}

	runeSlice := stringToRuneSlice(combined)
	if !reflect.DeepEqual(runeSlice, []rune(combined)) {
		t.Fatal("Rune slice conversion failed")
	}

	hash := hashWithSalt(runeSlice, "go-2024")
	if len(hash) != 64 || !isValidHex(hash) {
		t.Fatalf("Invalid hash: %q", hash)
	}
}

// функция для проверки корректности hex строки
func isValidHex(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

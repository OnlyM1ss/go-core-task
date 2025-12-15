package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	original := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}

	result := sliceExample(original)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("sliceExample() = %v, want %v", result, expected)
	}
}

func TestAddElements(t *testing.T) {
	original := []int{1, 2, 3}
	value := 4
	expected := []int{1, 2, 3, 4}

	result := addElements(original, value)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("addElements() = %v, want %v", result, expected)
	}

	// убедимся, что исходный слайс тоже логично изменился (append модифицирует базовый массив)
	if len(original) != 3 {
		t.Fatalf("ожидали исходную длину 3, получили %d", len(original))
	}
}

func TestCopySliceIndependence(t *testing.T) {
	original := []int{1, 2, 3, 4, 5}
	copied := copySlice(original)

	if !reflect.DeepEqual(original, copied) {
		t.Fatalf("copySlice() = %v, want %v", copied, original)
	}

	// изменяем original и проверяем, что copy не изменился
	original[0] = 999
	if copied[0] == 999 {
		t.Fatalf("копия изменилась при изменении оригинала: copied = %v", copied)
	}
}

func TestCopySliceNil(t *testing.T) {
	var original []int
	copied := copySlice(original)
	if copied != nil {
		t.Fatalf("ожидали nil копию для nil слайса, получили %v", copied)
	}
}

func TestRemoveElementMiddle(t *testing.T) {
	original := []int{10, 20, 30, 40, 50}
	result := removeElement(original, 2) // удаляем 30
	expected := []int{10, 20, 40, 50}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("removeElement() = %v, want %v", result, expected)
	}

	// original не должен измениться
	if !reflect.DeepEqual(original, []int{10, 20, 30, 40, 50}) {
		t.Fatalf("original слайс был изменен: %v", original)
	}
}

func TestRemoveElementFirstAndLast(t *testing.T) {
	original := []int{1, 2, 3, 4}

	resFirst := removeElement(original, 0)
	expFirst := []int{2, 3, 4}
	if !reflect.DeepEqual(resFirst, expFirst) {
		t.Fatalf("removeElement first = %v, want %v", resFirst, expFirst)
	}

	resLast := removeElement(original, len(original)-1)
	expLast := []int{1, 2, 3}
	if !reflect.DeepEqual(resLast, expLast) {
		t.Fatalf("removeElement last = %v, want %v", resLast, expLast)
	}
}

func TestRemoveElementInvalidIndex(t *testing.T) {
	original := []int{1, 2, 3}

	resNeg := removeElement(original, -1)
	if !reflect.DeepEqual(resNeg, original) {
		t.Fatalf("ожидали исходный слайс для отрицательного индекса, получили %v", resNeg)
	}

	resTooBig := removeElement(original, 3)
	if !reflect.DeepEqual(resTooBig, original) {
		t.Fatalf("ожидали исходный слайс для слишком большого индекса, получили %v", resTooBig)
	}
}

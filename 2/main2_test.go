package main

import "testing"

func TestSliceExample(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{2, 4, 6, 8, 10}
	result := sliceExample(input)
	if len(result) != len(expected) {
		t.Errorf("sliceExample() = %v, want %v", result, expected)
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("sliceExample() = %v, want %v", result, expected)
		}
	}
}

func TestAddElements(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	element := 42
	expected := []int{1, 2, 3, 4, 5, 42}
	result := addElements(input, element)
	if len(result) != len(expected) {
		t.Errorf("addElements() = %v, want %v", result, expected)
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("addElements() = %v, want %v", result, expected)
		}
	}
}

func TestCopySlice(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := copySlice(input)
	if len(result) != len(input) {
		t.Errorf("copySlice() = %v, want %v", result, input)
	}
	for i := range result {
		if result[i] != input[i] {
			t.Errorf("copySlice() = %v, want %v", result, input)
		}
	}
	result[0] = 42
	if input[0] == 42 {
		t.Errorf("copySlice() did not create a true copy; changes to result affected input")
	}
}

func TestRemoveElement(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	index := 2
	expected := []int{1, 2, 4, 5}
	result := removeElement(input, index)
	if len(result) != len(expected) {
		t.Errorf("removeElement() = %v, want %v", result, expected)
	}
	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("removeElement() = %v, want %v", result, expected)
		}
	}
}

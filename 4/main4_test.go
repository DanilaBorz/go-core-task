package main

import "testing"

func TestDifference(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := difference(slice1, slice2)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
		}
	}
}

func TestDifferenceEmpty(t *testing.T) {
	slice1 := []string{}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{}

	result := difference(slice1, slice2)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}

func TestDifferenceNoDifference(t *testing.T) {
	slice1 := []string{"banana", "date", "fig"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{}

	result := difference(slice1, slice2)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}
}

func TestDifferenceAllDifferent(t *testing.T) {
	slice1 := []string{"apple", "cherry", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := difference(slice1, slice2)
	if len(result) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
		}
	}
}

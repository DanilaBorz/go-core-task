package main

import "testing"

func TestStringIntMap(t *testing.T) {
	slice1 := []int{65, 3, 58, 678, 64}
	slice2 := []int{64, 2, 3, 43}
	expectedHasIntersection := true
	expectedIntersection := []int{64, 3}

	hasIntersection, intersection := difference(slice1, slice2)
	if hasIntersection != expectedHasIntersection {
		t.Errorf("Expected hasIntersection %v, got %v", expectedHasIntersection, hasIntersection)
	}

	if len(intersection) != len(expectedIntersection) {
		t.Errorf("Expected intersection length %d, got %d", len(expectedIntersection), len(intersection))
	}

	for i, v := range expectedIntersection {
		if intersection[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, intersection[i])
		}
	}
}

func TestStringIntMapNoIntersection(t *testing.T) {
	slice1 := []int{65, 58, 678}
	slice2 := []int{64, 2, 3, 43}
	expectedHasIntersection := false
	expectedIntersection := []int{}

	hasIntersection, intersection := difference(slice1, slice2)
	if hasIntersection != expectedHasIntersection {
		t.Errorf("Expected hasIntersection %v, got %v", expectedHasIntersection, hasIntersection)
	}

	if len(intersection) != len(expectedIntersection) {
		t.Errorf("Expected intersection length %d, got %d", len(expectedIntersection), len(intersection))
	}
}

func TestStringIntMapAllIntersection(t *testing.T) {
	slice1 := []int{64, 2, 3, 43}
	slice2 := []int{64, 2, 3, 43}
	expectedHasIntersection := true
	expectedIntersection := []int{64, 2, 3, 43}

	hasIntersection, intersection := difference(slice1, slice2)
	if hasIntersection != expectedHasIntersection {
		t.Errorf("Expected hasIntersection %v, got %v", expectedHasIntersection, hasIntersection)
	}

	if len(intersection) != len(expectedIntersection) {
		t.Errorf("Expected intersection length %d, got %d", len(expectedIntersection), len(intersection))
	}

	for i, v := range expectedIntersection {
		if intersection[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, intersection[i])
		}
	}
}

func TestStringIntMapEmpty(t *testing.T) {
	slice1 := []int{}
	slice2 := []int{64, 2, 3, 43}
	expectedHasIntersection := false
	expectedIntersection := []int{}

	hasIntersection, intersection := difference(slice1, slice2)
	if hasIntersection != expectedHasIntersection {
		t.Errorf("Expected hasIntersection %v, got %v", expectedHasIntersection, hasIntersection)
	}

	if len(intersection) != len(expectedIntersection) {
		t.Errorf("Expected intersection length %d, got %d", len(expectedIntersection), len(intersection))
	}
}

func TestStringIntMapEmptyBoth(t *testing.T) {
	slice1 := []int{}
	slice2 := []int{}
	expectedHasIntersection := false
	expectedIntersection := []int{}

	hasIntersection, intersection := difference(slice1, slice2)
	if hasIntersection != expectedHasIntersection {
		t.Errorf("Expected hasIntersection %v, got %v", expectedHasIntersection, hasIntersection)
	}

	if len(intersection) != len(expectedIntersection) {
		t.Errorf("Expected intersection length %d, got %d", len(expectedIntersection), len(intersection))
	}
}

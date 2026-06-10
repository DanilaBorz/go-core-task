package main

import "testing"

func TestGetType(t *testing.T) {
	tests := []struct {
		name     string
		value    interface{}
		expected string
	}{
		{"int", 42, "int"},
		{"float64", 3.14, "float64"},
		{"string", "Golang", "string"},
		{"bool", true, "bool"},
		{"complex64", complex64(1 + 2i), "complex64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.value); got != tt.expected {
				t.Errorf("getType() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCombineToString(t *testing.T) {
	result := combineToString(42, 3.14, "Golang", true, complex64(1+2i))
	expected := "423.14Golangtrue(1+2i)"
	if result != expected {
		t.Errorf("combineToString() = %v, want %v", result, expected)
	}
}

func TestHashWithSalt(t *testing.T) {
	tests := []struct {
		name  string
		input []rune
		salt  string
	}{
		{"normal", []rune("test"), "go-2024"},
		{"empty runes", []rune(""), "go-2024"},
		{"single", []rune("a"), "go-2024"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hashWithSalt(tt.input, tt.salt)
			if result == "" {
				t.Errorf("hashWithSalt() returned empty string")
			}
			if len(result) != 64 {
				t.Errorf("hashWithSalt() returned hash of incorrect length: got %d, want 64", len(result))
			}
		})

	}
}

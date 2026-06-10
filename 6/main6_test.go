package main

import "testing"

func TestGenerator(t *testing.T) {
	count := 10
	randomNumbers := generateRandomNumbers(count)

	for i := 0; i < count; i++ {
		num, ok := <-randomNumbers
		if !ok {
			t.Fatalf("Expected %d numbers, but channel closed early", count)
		}
		if num < 0 || num >= 100 {
			t.Errorf("Generated number %d is out of expected range (0-99)", num)
		}
	}

	// Проверяем, что канал закрыт после генерации всех чисел
	if _, ok := <-randomNumbers; ok {
		t.Error("Expected channel to be closed after generating all numbers")
	}
}

func TestGeneratorCount(t *testing.T) {
	count := 5
	randomNumbers := generateRandomNumbers(count)

	for i := 0; i < count; i++ {
		if _, ok := <-randomNumbers; !ok {
			t.Fatalf("Expected %d numbers, but channel closed early", count)
		}
	}

	// Проверяем, что канал закрыт после генерации всех чисел
	if _, ok := <-randomNumbers; ok {
		t.Error("Expected channel to be closed after generating all numbers")
	}
}

func TestGeneratorZeroCount(t *testing.T) {
	count := 0
	randomNumbers := generateRandomNumbers(count)

	// Проверяем, что канал закрыт сразу, так как count = 0
	if _, ok := <-randomNumbers; ok {
		t.Error("Expected channel to be closed immediately for count = 0")
	}
}

func TestGeneratorNegativeCount(t *testing.T) {
	count := -5
	randomNumbers := generateRandomNumbers(count)

	// Проверяем, что канал закрыт сразу, так как count < 0
	if _, ok := <-randomNumbers; ok {
		t.Error("Expected channel to be closed immediately for negative count")
	}
}

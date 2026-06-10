package main

import (
	"fmt"
	"math/rand"
)

// Задание 2
// Создайте слайс целых чисел originalSlice, содержащий 10 произвольных
//  значений, которые генерируются случайным образом (при каждом запуске должны получаться новые значения)

// Напишите функцию sliceExample, которая принимает слайс
//  и возвращает новый слайс, содержащий только четные числа из исходного слайса.

// Напишите функцию addElements, которая принимает слайс и число. Функция должна
//  добавлять это число в конец слайса и возвращать новый слайс.

// Напишите функцию copySlice, которая принимает слайс и возвращает его копию. Убедитесь,
// что изменения в оригинальном слайсе не влияют на его копию.

// Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно
// удалить. Функция должна возвращать новый слайс без элемента по указанному индексу.

// Напишите main функцию, в которой протестируете все вышеописанные функции.
// Выведите результаты на экран.

// Напишите unit тесты к созданным функциям
// Примечание. В качестве originalSlice можно использовать originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
func sliceExample(s []int) []int {
	var evenNumbers []int
	for _, num := range s {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	return evenNumbers
}

func addElements(s []int, element int) []int {
	return append(s, element)
}

func copySlice(s []int) []int {
	copied := make([]int, len(s))
	copy(copied, s)
	return copied
}

func removeElement(s []int, index int) []int {
	result := make([]int, 0, len(s)-1)
	result = append(result, s[:index]...)
	result = append(result, s[index+1:]...)
	return result
}

func main() {
	originalSlice := []int{}
	for i := 0; i < 10; i++ {
		originalSlice = append(originalSlice, rand.Intn(100))
	}
	fmt.Println("Original Slice:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Numbers:", evenSlice)

	newSlice := addElements(originalSlice, 42)
	fmt.Println("After Adding Element:", newSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	removedElementSlice := removeElement(originalSlice, 2)
	fmt.Println("After Removing Element at Index 2:", removedElementSlice)
}

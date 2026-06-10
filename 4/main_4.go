// Задание 4
// На вход подаются два неупорядоченных слайса строк. Например:

// slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
// slice2 := []string{"banana", "date", "fig"}
// Напишите функцию, которая возвращает слайс строк, содержащий элементы, которые есть в первом слайсе, но отсутствуют во втором.

// Напишите unit тесты к созданным функциям
package main

import "fmt"

func difference(slice1, slice2 []string) []string {
	excluded := make(map[string]struct{}, len(slice2))
	for _, str := range slice2 {
		excluded[str] = struct{}{}
	}

	result := make([]string, 0, len(slice1))
	for _, str1 := range slice1 {
		if _, found := excluded[str1]; !found {
			result = append(result, str1)
		}
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	result := difference(slice1, slice2)
	fmt.Println("Difference:", result)
}

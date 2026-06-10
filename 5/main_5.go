// Задание 5
// На вход подаются два неупорядоченных слайса int любой длины. Например:

// a := []int{65, 3, 58, 678, 64}
// b := []int{64, 2, 3, 43}
// Напишите функцию, которая проверяет, есть ли
// пересечения значений между двумя слайсами и возвращает:

// bool значение есть ли хотя бы одно пересечение
// в значениях входных срезов
// срез []int с пересеченными значениями (если таких значений нет,
// то возвращать пустой срез). Т. е. если взать слайсы a и b из премере,
// то вернет true, []int{64, 3}.
// Напишите unit тесты к созданным функциям

package main

import "fmt"

func difference(slice1, slice2 []int) (bool, []int) {
	intersection := make(map[int]struct{})
	for _, v := range slice1 {
		intersection[v] = struct{}{}
	}

	hasIntersection := false
	result := make([]int, 0)

	for _, v := range slice2 {
		if _, found := intersection[v]; found {
			hasIntersection = true
			result = append(result, v)
		}
	}

	return hasIntersection, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	hasIntersection, intersection := difference(a, b)
	fmt.Println("Has intersection:", hasIntersection)
	fmt.Println("Intersection values:", intersection)
}

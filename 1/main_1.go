// Задание 1
// Напишите программу на Go, которая:

// Создает несколько переменных различных типов данных:
// int (три числа в десятичной, восьмеричной и шеснадцатиричной системах)
// float64
// string
// bool
// complex64
// Определяет тип каждой переменной и выводит его на экран.
// Преобразует все переменные в строковый тип и объединяет их в одну строку.
// Преобразовать эту строку в срез рун.
// Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
// Напишите unit тесты к созданным функциям
// Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

// Входные числа из пункта 1 могут быть:

// var numDecimal int = 42           // Десятичная система
// var numOctal int = 052            // Восьмеричная система
// var numHexadecimal int = 0x2A     // Шестнадцатиричная система
// var pi float64 = 3.14             // Тип float64
// var name string = "Golang"         // Тип string
// var isActive bool = true           // Тип bool
// var complexNum complex64 = 1 + 2i  // Тип complex64
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	testValues()
}

func getType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func combineToString(values ...interface{}) string {
	var result string
	for _, v := range values {
		result += fmt.Sprintf("%v", v)
	}
	return result
}

func hashWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	hasher := sha256.New()
	leftPart := string(runes[:mid])
	rightPart := string(runes[mid:])
	hasher.Write([]byte(leftPart + salt + rightPart))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func testValues() {
	var numDecimal = 42               // Десятичная система
	var numOctal = 052                // Восьмеричная система (42)
	var numHexadecimal = 0x2A         // Шестнадцатиричная система (42)
	var pi = 3.14                     // Тип float64
	var name = "Golang"               // Тип string
	var isActive = true               // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	fmt.Println("Типы переменных")
	fmt.Printf("numDecimal (%d): %s\n", numDecimal, getType(numDecimal))
	fmt.Printf("numOctal (%d): %s\n", numOctal, getType(numOctal))
	fmt.Printf("numHexadecimal (%d): %s\n", numHexadecimal, getType(numHexadecimal))
	fmt.Printf("pi (%.2f): %s\n", pi, getType(pi))
	fmt.Printf("name (%s): %s\n", name, getType(name))
	fmt.Printf("isActive (%t): %s\n", isActive, getType(isActive))
	fmt.Printf("complexNum (%.1f%+.1fi): %s\n", real(complexNum), imag(complexNum), getType(complexNum))

	combined := combineToString(numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum)
	fmt.Printf("\nОбъединенная строка: %s\n", combined)

	runes := []rune(combined)
	fmt.Printf("Срез рун: %v\n", runes)
	fmt.Printf("Длина среза рун: %d\n", len(runes))

	hashed := hashWithSalt(runes, "go-2024")
	fmt.Printf("\nSHA256 хэш с солью 'go-2024': %s\n", hashed)
}

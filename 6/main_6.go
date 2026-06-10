// Задание 6
// Напишите генератор случайных чисел используя небуфферизированный канал.

// Напишите unit тесты к созданным функциям

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumbers := generateRandomNumbers(10)
	for num := range randomNumbers {
		fmt.Println(num)
	}
}

func generateRandomNumbers(count int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < count; i++ {
			ch <- rand.Intn(100) // ^_^ число от 0 до 99
		}
	}()
	return ch
}

package main

import (
	"fmt"
	"sync"
	"time"
)

// Задание 7
// Напишите программу на Go, которая сливает N каналов в один.

// Напишите unit тесты к созданным функциям

func mergeChannels(channels ...<-chan int) <-chan int {
	merged := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch <-chan int) {
			defer wg.Done()
			for num := range ch {
				merged <- num
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
	}()
	go func() {
		time.Sleep(2 * time.Second) // Добавляем задержку, чтобы показать, что каналы могут работать асинхронно
		defer close(ch2)
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
	}()
	go func() {
		defer close(ch3)
		for i := 10; i < 15; i++ {
			ch3 <- i
		}
	}()
	merged := mergeChannels(ch1, ch2, ch3)
	for num := range merged {
		fmt.Println(num)
	}
}

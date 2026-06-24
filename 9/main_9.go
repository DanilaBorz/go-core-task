package main

// Задание 9
// Сделать конвейер чисел Даны два канала.
// В первый пишутся числа типа uint8. Нужно, чтобы
// числа читались из первого канала по мере поступления,
//  затем эти числа должны преобразовываться в float64 и
// возводиться в куб и результат записывался во второй канал.

// Напишите main функцию, в которой протестируете весь
//  вышеописанный функционал. Выведите результаты на экран.

// Напишите unit тесты к созданным функциям

func uint8Generator(out chan<- uint8, nums ...uint8) {
	for _, num := range nums {
		out <- num
	}
	close(out)
}

func uint8ToFloat64(in <-chan uint8, out chan<- float64) {
	for num := range in {
		out <- float64(num)
	}
	close(out)
}
func cubePipeline(in <-chan uint8, out chan<- float64) {
	for num := range in {
		value := float64(num)
		out <- value * value * value
	}
	close(out)
}

func main() {
	uint8Chan := make(chan uint8)
	float64Chan := make(chan float64)

	go uint8Generator(uint8Chan, 1, 2, 3, 4, 5)
	go uint8ToFloat64(uint8Chan, float64Chan)
	go cubePipeline(uint8Chan, float64Chan)
}

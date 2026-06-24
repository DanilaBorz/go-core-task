// Задание 8
// Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

// Напишите unit тесты к созданным функциям

package main

import "sync"

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		done: make(chan struct{}),
	}
}

type WaitGroup struct {
	mu    sync.Mutex
	count int
	done  chan struct{}
}

func (wg *WaitGroup) Add(delta int) {
	if delta < 0 {
		return
	}

	wg.mu.Lock()
	defer wg.mu.Unlock()

	if wg.count == 0 {
		wg.done = make(chan struct{})
	}
	wg.count += delta
}

func (wg *WaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	if wg.count == 0 {
		return
	}

	wg.count--
	if wg.count == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Wait() {
	wg.mu.Lock()
	done := wg.done
	count := wg.count
	wg.mu.Unlock()

	if count == 0 {
		return
	}

	<-done
}

func main() {
	customWaitGroup := NewWaitGroup()
	customWaitGroup.Add(3)

	go func() {
		defer customWaitGroup.Done()
		// Выполнение задачи 1
	}()

	go func() {
		defer customWaitGroup.Done()
		// Выполнение задачи 2
	}()

	go func() {
		defer customWaitGroup.Done()
		// Выполнение задачи 3
	}()

	customWaitGroup.Wait()
	// Все задачи завершены, можно продолжать выполнение
}

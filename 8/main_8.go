// Задание 8
// Сделать кастомную waitGroup на семафоре, не используя sync.WaitGroup.

// Напишите unit тесты к созданным функциям

package main

func NewWaitGroup() *WaitGroup {
	wg := &WaitGroup{
		sem:  make(chan struct{}, 1),
		done: make(chan struct{}),
	}
	wg.sem <- struct{}{}
	return wg
}

type WaitGroup struct {
	sem   chan struct{}
	count int
	done  chan struct{}
}

func (wg *WaitGroup) lock() {
	<-wg.sem
}

func (wg *WaitGroup) unlock() {
	wg.sem <- struct{}{}
}

func (wg *WaitGroup) Add(delta int) {
	if delta < 0 {
		return
	}

	wg.lock()
	defer wg.unlock()

	if wg.count == 0 {
		wg.done = make(chan struct{})
	}
	wg.count += delta
}

func (wg *WaitGroup) Done() {
	wg.lock()
	defer wg.unlock()

	if wg.count == 0 {
		return
	}

	wg.count--
	if wg.count == 0 {
		close(wg.done)
	}
}

func (wg *WaitGroup) Wait() {
	wg.lock()
	done := wg.done
	count := wg.count
	wg.unlock()

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

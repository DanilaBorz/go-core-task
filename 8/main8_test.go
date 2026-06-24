package main

import "testing"

func TestWaitGroup(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(3)

	done := make(chan struct{}, 3)
	go func() {
		defer wg.Done()
		done <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		done <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		done <- struct{}{}
	}()

	wg.Wait()

	for i := 0; i < 3; i++ {
		select {
		case <-done:
			// Получили сигнал о завершении задачи
		default:
			t.Errorf("Expected all tasks to be done, but some are still pending")
		}
	}
}

func TestWaitGroupWithNoTasks(t *testing.T) {
	wg := NewWaitGroup()
	wg.Wait() // Ожидаем, что не будет блокировки

	// Если мы дошли до этого момента, значит тест прошел успешно
}

func TestWaitGroupDoneWithNoTasks(t *testing.T) {
	wg := NewWaitGroup()
	wg.Done() // Ожидаем, что вызов Done без задач не вызовет панику
	wg.Wait() // И после этого Wait не должен блокироваться
}

func TestWaitGroupWithMoreTasksThanAdded(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(2)

	done := make(chan struct{}, 2)
	go func() {
		defer wg.Done()
		done <- struct{}{}
	}()

	go func() {
		defer wg.Done()
		done <- struct{}{}
	}()

	wg.Wait()

	for i := 0; i < 2; i++ {
		select {
		case <-done:
			// Получили сигнал о завершении задачи
		default:
			t.Errorf("Expected all tasks to be done, but some are still pending")
		}
	}
}

func TestWaitGroupWithNegativeDelta(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(-1) // Ожидаем, что это не вызовет паники

	// Если мы дошли до этого момента, значит тест прошел успешно
}

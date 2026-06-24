package main

import (
	"testing"
	"time"
)

func TestWaitGroupWaitsForAllTasks(t *testing.T) {
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

func TestWaitGroupBlocksUntilDone(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(1)

	waited := make(chan struct{})
	go func() {
		wg.Wait()
		close(waited)
	}()

	select {
	case <-waited:
		t.Fatal("Wait returned before Done was called")
	case <-time.After(20 * time.Millisecond):
	}

	wg.Done()

	select {
	case <-waited:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return after Done was called")
	}
}

func TestWaitGroupWithNoTasks(t *testing.T) {
	wg := NewWaitGroup()
	wg.Wait() // Ожидаем, что не будет блокировки

	// Если мы дошли до этого момента, значит тест прошел успешно
}

func TestWaitGroupDoneWithNoTasksDoesNotBlock(t *testing.T) {
	wg := NewWaitGroup()
	wg.Done() // Ожидаем, что вызов Done без задач не вызовет панику
	wg.Wait() // И после этого Wait не должен блокироваться
}

func TestWaitGroupIgnoresExtraDoneCalls(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(2)

	wg.Done()
	wg.Done()
	wg.Done()
	wg.Wait()
}

func TestWaitGroupWithNegativeDelta(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(-1) // Ожидаем, что это не вызовет паники

	// Если мы дошли до этого момента, значит тест прошел успешно
}

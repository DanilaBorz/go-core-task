package main

import (
	"sort"
	"testing"
	"time"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go sendValues(ch1, 1, 2, 3)
	go sendValues(ch2, 4, 5)
	go sendValues(ch3, 6)

	var result []int
	for value := range mergeChannels(ch1, ch2, ch3) {
		result = append(result, value)
	}

	sort.Ints(result)

	expected := []int{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Fatalf("expected %d values, got %d", len(expected), len(result))
	}

	for i, value := range expected {
		if result[i] != value {
			t.Fatalf("expected %v, got %v", expected, result)
		}
	}
}

func TestMergeChannelsDoesNotWaitForEarlierSlowChannel(t *testing.T) {
	slow := make(chan int)
	fast := make(chan int)

	go func() {
		defer close(slow)
		time.Sleep(50 * time.Millisecond)
		slow <- 1
	}()

	go sendValues(fast, 2)

	merged := mergeChannels(slow, fast)

	select {
	case value := <-merged:
		if value != 2 {
			t.Fatalf("expected fast channel value first, got %d", value)
		}
	case <-time.After(20 * time.Millisecond):
		t.Fatal("mergeChannels waited for the earlier slow channel")
	}
}

func sendValues(ch chan<- int, values ...int) {
	defer close(ch)
	for _, value := range values {
		ch <- value
	}
}

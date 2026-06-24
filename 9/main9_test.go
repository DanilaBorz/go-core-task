package main

import (
	"reflect"
	"testing"
)

func TestUint8Generator(t *testing.T) {
	out := make(chan uint8)

	go uint8Generator(out, 1, 2, 3)

	var got []uint8
	for value := range out {
		got = append(got, value)
	}

	want := []uint8{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestUint8ToFloat64(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		in <- 1
		in <- 2
		in <- 255
	}()
	go uint8ToFloat64(in, out)

	var got []float64
	for value := range out {
		got = append(got, value)
	}

	want := []float64{1, 2, 255}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestCubePipeline(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		in <- 2
		in <- 3
		in <- 10
	}()
	go cubePipeline(in, out)

	var got []float64
	for value := range out {
		got = append(got, value)
	}

	want := []float64{8, 27, 1000}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestCubePipelineWithNoValues(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	close(in)

	go cubePipeline(in, out)

	got, ok := <-out
	if ok {
		t.Fatalf("expected output channel to be closed, got value %v", got)
	}
}

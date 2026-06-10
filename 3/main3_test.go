package main

import "testing"

func TestNewStringIntMap(t *testing.T) {
	sim := NewStringIntMap()
	if sim == nil {
		t.Error("Expected NewStringIntMap to return a non-nil value")
	}
}

func TestAddAndGet(t *testing.T) {
	sim := NewStringIntMap()
	sim.Add("key1", 10)

	value, exists := sim.Get("key1")
	if !exists {
		t.Error("Expected key1 to exist in the map")
	}
	if value != 10 {
		t.Errorf("Expected value for key1 to be 10, got %d", value)
	}
}

func TestRemove(t *testing.T) {
	sim := NewStringIntMap()
	sim.Add("key1", 10)
	sim.Remove("key1")

	_, exists := sim.Get("key1")
	if exists {
		t.Error("Expected key1 to be removed from the map")
	}
}

func TestCopy(t *testing.T) {
	sim := NewStringIntMap()
	sim.Add("key1", 10)
	sim.Add("key2", 20)

	copyMap := sim.Copy()
	if len(copyMap) != 2 {
		t.Errorf("Expected copy map to have 2 elements, got %d", len(copyMap))
	}
	if copyMap["key1"] != 10 || copyMap["key2"] != 20 {
		t.Error("Expected copy map to contain the correct key-value pairs")
	}
}

func TestExists(t *testing.T) {
	sim := NewStringIntMap()
	sim.Add("key1", 10)

	if !sim.Exists("key1") {
		t.Error("Expected key1 to exist in the map")
	}
	if sim.Exists("key2") {
		t.Error("Expected key2 to not exist in the map")
	}
}

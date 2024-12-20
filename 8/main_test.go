package main

import "testing"

func TestWaitGroup(t *testing.T) {
	cwg := NewCustomWaitGroup()

	cwg.Add(3)

	completed := 0

	for i := 0; i < 3; i++ {
		go func() {
			defer cwg.Done()
			completed++
		}()
	}

	cwg.Wait()

	if completed != 3 {
		t.Errorf("expected 3 completed tasks, but got %d", completed)
	}
}

func TestCustomWaitGroup_PanicOnNegativeCounter(t *testing.T) {
	cwg := NewCustomWaitGroup()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic but did not get one")
		}
	}()

	cwg.Add(-1)
}

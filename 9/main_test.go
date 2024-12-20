package main

import "testing"

func TestCounter(t *testing.T) {
	ch := make(chan uint8)
	go counter(ch, 10)

	expected := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var result []uint8

	for v := range ch {
		result = append(result, v)
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], v)
		}
	}
}

func TestCuber(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go cuber(out, in)

	go func() {
		for i := uint8(0); i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	expected := []float64{0, 1, 8, 27, 64, 125, 216, 343, 512, 729}
	var result []float64

	for v := range out {
		result = append(result, v)
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %f, got %f", expected[i], v)
		}
	}
}

func TestPrinter(t *testing.T) {

	out := make(chan float64)

	go func() {
		for i := 0; i < 10; i++ {
			out <- float64(i * i * i)
		}
		close(out)
	}()

	printer(out)
}

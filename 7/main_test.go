package main

import "testing"

type orTest struct {
	name     string
	channels []chan int
}

var orTests = []orTest{
	{
		name: "test with 2 int channels",
		channels: []chan int{
			make(chan int),
			make(chan int),
		},
	},
	{
		name: "test with 3 int channels",
		channels: []chan int{
			make(chan int),
			make(chan int),
			make(chan int),
		},
	},
	{
		name: "test with 4 int channels",
		channels: []chan int{
			make(chan int),
			make(chan int),
			make(chan int),
		},
	},
	{
		name: "test with 10 int channels",
		channels: []chan int{
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
			make(chan int),
		},
	},
}

func TestOr(t *testing.T) {
	for _, tt := range orTests {
		t.Run(tt.name, func(t *testing.T) {
			testOr(t, tt)
		})
	}
}

func testOr(t *testing.T, tt orTest) {
	for i, ch := range tt.channels {
		go func(c chan int, index int) {
			defer close(ch)
			for j := 0; j < 5; j++ {
				c <- j + index*10
			}
		}(ch, i)
	}

	channels := make([]<-chan int, len(tt.channels))
	for i, ch := range tt.channels {
		channels[i] = ch
	}

	result := or(channels...)
	received := make(map[int]bool)
	for val := range result {
		if _, ok := received[val]; ok {
			t.Errorf("duplicate value received: %d", val)
		}
		received[val] = true
	}

	expectedCount := 5 * len(tt.channels)
	if len(received) != expectedCount {
		t.Errorf("expected %d unique values, got %d", expectedCount, len(received))
	}
}

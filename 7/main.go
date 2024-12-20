package main

import "sync"

func or[T any](channels ...<-chan T) <-chan T {
	sink := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(с <-chan T) {
			defer wg.Done()
			for v := range channel {
				sink <- v
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(sink)
	}()

	return sink
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i < 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 1; i < 10; i++ {
			ch2 <- i
		}
	}()

	result := or(ch1, ch2)

	for val := range result {
		println(val)
	}
}

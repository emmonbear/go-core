package main

import "sync"

func or(channels ...<-chan int) <-chan int {
	sink := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		channel := channel
		go func() {
			defer wg.Done()
			for  v := range channel {
				sink <- v
			}
		}()
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
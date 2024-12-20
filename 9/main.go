package main

import "fmt"

// counter generates a sequence of natural numbers (0 to count-1) and sends them on the out channel.
// It then closes the out channel once all values are sent.
func counter(out chan<- uint8, count int) {
	for i := 0; i < count; i++ {
		out <- (uint8(i))
	}
	close(out)
}

// cuber reads numbers from the in channel, cubes them, and sends the result on the out channel.
// It closes the out channel once all values have been processed.
func cuber(out chan<- float64, in <-chan uint8) {
	for v := range in {
		result := float64(v) * float64(v) * float64(v)
		out <- result
	}
	close(out)
}

// printer reads numbers from the in channel and prints them to the console.
func printer(in <-chan float64) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan uint8)
	cubes := make(chan float64)

	go counter(naturals, 10)
	go cuber(cubes, naturals)
	printer(cubes)
}

package main

import "fmt"

func counter(out chan<- uint8, count int) {
	for i := 0; i < count; i++ {
		out <- (uint8(i))
	}
	close(out)
}

func cuber(out chan<- float64, in <-chan uint8) {
	for v := range in {
		result := float64(v) * float64(v) * float64(v)
		out <- result
	}
	close(out)
}

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
package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	// minValue is the minimum value of the range for generated random numbers.
	minValue = -500

	// maxValue is the maximum value of the range for generated random numbers.
	maxValue = 500
)

// generateRandomNumber generates a specified number of random integers and sends them to the channel.
// The random integers are in the range of minValue to maxValue.
//
// Parameters:
//   - ch: The channel to send the random numbers to.
//   - count: The number of random numbers to generate.
func generateRandomNumber(ch chan<- int, count int) {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		ch <- gen.Intn((maxValue-minValue)+1) + minValue
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go generateRandomNumber(ch, 10)
	for num := range ch {
		fmt.Println(num)
	}
}

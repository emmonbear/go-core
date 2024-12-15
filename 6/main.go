package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generateRandomNumber generates a specified number of random integers and sends them to the channel.
// The random integers are in the range of -500 to 499.
//
// Parameters:
//   - ch: The channel to send the random numbers to.
//   - count: The number of random numbers to generate.
func generateRandomNumber(ch chan<- int, count int) {
	gen := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		ch <- gen.Intn(1000) - 500
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

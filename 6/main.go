package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

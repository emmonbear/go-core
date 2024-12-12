package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateNewSlice(length int) []int {
	gen := rand.New(rand.NewSource(time.Now().Unix()))

	slice := make([]int, length)

	for i := range slice {
		slice[i] = gen.Intn(100)
	}

	return slice
}

func main() {
	originalSlice := GenerateNewSlice(10)
	fmt.Println("Original slice: ", originalSlice)
}

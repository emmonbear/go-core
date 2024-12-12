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

func SliceExample(slice []int) []int {
	var evenSlice []int
	for _, num := range slice {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}
	return evenSlice
}

func main() {
	originalSlice := GenerateNewSlice(10)
	slice := SliceExample(originalSlice)
	fmt.Println("Original slice: ", originalSlice)
	fmt.Println("Slice slice: ", slice)
}

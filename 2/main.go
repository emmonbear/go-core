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

func AppendInt(slice []int, elems ...int) []int {
	sliceLen := len(slice)
	newLen := sliceLen + len(elems)

	if newLen > cap(slice) {
		newCap := newLen
		if newCap < cap(slice)*2 {
			newCap = 2 * cap(slice)
		}
		newSlice := make([]int, newLen, newCap)
		copy(newSlice, slice)
		slice = newSlice
	}

	copy(slice[sliceLen:], elems)

	return slice[:newLen]
}

func main() {
	originalSlice := GenerateNewSlice(10)
	slice := SliceExample(originalSlice)
	fmt.Println("Original slice: ", originalSlice)
	fmt.Println("Slice slice: ", slice)
	add := AppendInt(slice, 1, 2, 3, 4)
	fmt.Println("Add: ", add)
}

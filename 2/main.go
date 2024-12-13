package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateNewSlice creates a slice of random integers with the specified length.
// Each integer is randomly generated in the range [0, 100).
// The randomness is seeded with the current Unix timestamp to ensure different results on each execution.
//
// Parameters:
// - length: The length of the slice to be generated.
//
// Returns:
// - A slice of integers with random values.
func GenerateNewSlice(length int) []int {
	gen := rand.New(rand.NewSource(time.Now().Unix()))

	slice := make([]int, length)

	for i := range slice {
		slice[i] = gen.Intn(100)
	}

	return slice
}

// SliceExample filters the input slice to return a new slice containing only even integers.
//
// Parameters:
// - slice: The input slice to be filtered.
//
// Returns:
// - A new slice containing only even integers from the input slice.
func SliceExample(slice []int) []int {
	var evenSlice []int
	for _, num := range slice {
		if num%2 == 0 {
			evenSlice = append(evenSlice, num)
		}
	}
	return evenSlice
}

// AddElements appends multiple integers to the end of the slice, ensuring the slice has enough capacity.
// If necessary, it creates a new slice with increased capacity to accommodate the new elements.
//
// Parameters:
// - slice: The original slice to which the elements will be added.
// - elems: The integers to be added to the slice.
//
// Returns:
// - A new slice with the added integers.
func AddElements(slice []int, elems ...int) []int {
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

func CopySlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)

	return result
}

func RemoveElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func main() {
	originalSlice := GenerateNewSlice(10)
	even := SliceExample(originalSlice)
	fmt.Println("Original slice: ", originalSlice)
	fmt.Println("Even slice: ", even)
	add := AddElements(originalSlice, 1, 2, 3, 4)
	fmt.Println("Add: ", add)

	copy := CopySlice(originalSlice)
	originalSlice = AddElements(originalSlice, 1, 2, 321)
	fmt.Println("Original slice: ", originalSlice)
	fmt.Println("Copy slice: ", copy)

	remove := RemoveElement(originalSlice, -1)
	fmt.Println(remove)

}

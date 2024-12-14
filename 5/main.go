package main

import "fmt"

// Compare function compares two slices of integers and finds the common elements.
// It returns a boolean indicating whether there is at least one common element,
// and a slice containing the unique common elements.
//
// Parameters:
// - slice1 []int: The first slice of integers to compare.
// - slice2 []int: The second slice of integers to compare.
//
// Returns:
// - bool: true if there are any common elements, false otherwise.
// - []int: A slice containing the unique common elements between slice1 and slice2.
func Compare(slice1, slice2 []int) (bool, []int) {
	set2 := make(map[int]struct{})
	for _, v := range slice2 {
		set2[v] = struct{}{}
	}

	var result []int
	seen := make(map[int]struct{})
	for _, v := range slice1 {
		if _, found := set2[v]; found {
			if _, alreadySeen := seen[v]; !alreadySeen {
				result = append(result, v)
				seen[v] = struct{}{}
			}
		}
	}

	if len(result) > 0 {
		return true, result
	}

	return false, []int{}
}

func main() {
	a := []int{1, 1, 2, 1}
	b := []int{1}

	expect, res := Compare(a, b)

	fmt.Println(expect, res)
}

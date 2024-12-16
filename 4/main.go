package main

import "fmt"

// Difference takes two slices of strings and returns a slice of strings
// containing the elements that are in slice1 but not in slice2.
//
// The function uses a map (set2) to store the elements of slice2 for
// efficient lookup while iterating over slice1.
//
// Parameters:
//
//	slice1: A slice of strings, from which the difference will be calculated.
//	slice2: A slice of strings, the elements of which will be excluded from slice1.
//
// Returns:
//
//	A slice of strings that contains the elements from slice1 that are not in slice2.
func Diffence(slice1, slice2 []string) []string {
	set2 := make(map[string]struct{})
	for _, v := range slice2 {
		set2[v] = struct{}{}
	}

	var result []string

	for _, v := range slice1 {
		if _, found := set2[v]; !found {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	set := Diffence(slice1, slice2)

	fmt.Println(set)
}

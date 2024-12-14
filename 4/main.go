package main

import "fmt"

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
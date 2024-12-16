package main

import "testing"

type compareTest struct {
	name     string
	slice1   []int
	slice2   []int
	expected []int
	exists   bool
}

var compareTests = []compareTest{
	{
		name:     "intersections exist",
		slice1:   []int{65, 3, 58, 678, 64},
		slice2:   []int{64, 2, 3, 43},
		expected: []int{3, 64},
		exists:   true,
	},
	{
		name:     "no intersections",
		slice1:   []int{65, 3, 58, 678},
		slice2:   []int{2, 43, 100},
		expected: []int{},
		exists:   false,
	},
	{
		name:     "first slice is empty",
		slice1:   []int{},
		slice2:   []int{2, 3, 43},
		expected: []int{},
		exists:   false,
	},
	{
		name:     "second slice is empty",
		slice1:   []int{2, 3, 43},
		slice2:   []int{},
		expected: []int{},
		exists:   false,
	},
	{
		name:     "same slices",
		slice1:   []int{1, 2, 3, 4, 5},
		slice2:   []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
		exists:   true,
	},
	{
		name:     "one slice is a subset of the other",
		slice1:   []int{1, 2, 3, 4, 5},
		slice2:   []int{2, 3},
		expected: []int{2, 3},
		exists:   true,
	},
	{
		name:     "repeated elements in first slice",
		slice1:   []int{2, 3, 2, 3, 5},
		slice2:   []int{3, 2},
		expected: []int{2, 3},
		exists:   true,
	},
	{
		name:     "no common elements, large slices",
		slice1:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		slice2:   []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		expected: []int{},
		exists:   false,
	},
	{
		name:     "all elements in slice1 are present in slice2",
		slice1:   []int{1, 2, 3, 4, 5},
		slice2:   []int{5, 4, 3, 2, 1},
		expected: []int{1, 2, 3, 4, 5},
		exists:   true,
	},
}

func TestCompare(t *testing.T) {
	for _, tt := range compareTests {
		t.Run(tt.name, func(t *testing.T) {
			testCompare(t, tt)
		})
	}
}

func testCompare(t *testing.T, tt compareTest) {
	exists, result := Compare(tt.slice1, tt.slice2)

	if exists != tt.exists {
		t.Errorf("expected exists %t, got %t", tt.exists, exists)
	}

	if len(result) != len(tt.expected) {
		t.Errorf("len(result) expected %d, got %d", len(tt.expected), len(result))
	}

	for i, v := range result {
		if tt.expected[i] != v {
			t.Errorf("expected %v, got %v", tt.expected, result)
		}
	}
}

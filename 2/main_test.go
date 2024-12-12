package main

import (
	"reflect"
	"testing"
)

type addElementsTest struct {
	name     string
	slice    []int
	values   []int
	expected []int
}

type sliceExampleTest struct {
	name     string
	slice    []int
	expected []int
}

type generateNewSliceTest struct {
	name   string
	length int
}

type copySliceTest struct {
	name     string
	slice    []int
	expected []int
}

var addElementsTests = []addElementsTest{
	{
		name:     "test 1",
		slice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		values:   []int{11, 12, 13},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
	},
	{
		name:     "test 2",
		slice:    []int{5, 10, 15},
		values:   []int{20, 25},
		expected: []int{5, 10, 15, 20, 25},
	},
	{
		name:     "test 3",
		slice:    []int{100, 200, 300},
		values:   []int{400},
		expected: []int{100, 200, 300, 400},
	},
	{
		name:     "test 4",
		slice:    []int{1, 2},
		values:   []int{3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "test 5",
		slice:    []int{},
		values:   []int{1, 2, 3},
		expected: []int{1, 2, 3},
	},
	{
		name:     "test 6",
		slice:    []int{10},
		values:   []int{20, 30},
		expected: []int{10, 20, 30},
	},
}

var sliceExampleTests = []sliceExampleTest{
	{
		name:     "test 1",
		slice:    []int{1, 2, 3, 4, 5},
		expected: []int{2, 4},
	},
	{
		name:     "test 2",
		slice:    []int{6, 7, 8, 9, 10},
		expected: []int{6, 8, 10},
	},
	{
		name:     "test 3",
		slice:    []int{1, 3, 5, 7, 9},
		expected: []int{}, // No even numbers
	},
	{
		name:     "test 4",
		slice:    []int{2, 4, 6, 8, 10},
		expected: []int{2, 4, 6, 8, 10}, // All numbers are even
	},
	{
		name:     "test 5",
		slice:    []int{10, 11, 12, 13, 14},
		expected: []int{10, 12, 14}, // Only even numbers
	},
	{
		name:     "test 6",
		slice:    []int{},
		expected: []int{}, // Empty slice
	},
}

var generateNewSliceTests = []generateNewSliceTest{
	{
		name:   "test 1",
		length: 10,
	},
	{
		name:   "test 2",
		length: 101,
	},
	{
		name:   "test 3",
		length: 21,
	},
}

var copySliceTests = []copySliceTest{
	{
		name:     "test 1",
		slice:    []int{1, 2, 3, 4, 5},
		expected: []int{1, 2, 3, 4, 5},
	},
	{
		name:     "test 2",
		slice:    []int{2},
		expected: []int{2},
	},
	{
		name:     "test 3",
		slice:    []int{10, 20, 30, 40},
		expected: []int{10, 20, 30, 40},
	},
	{
		name:     "test 4",
		slice:    []int{5, 15, 25, 35},
		expected: []int{5, 15, 25, 35},
	},
	{
		name:     "test 5",
		slice:    []int{-1, -2, -3, -4},
		expected: []int{-1, -2, -3, -4},
	},
	{
		name:     "test 6",
		slice:    []int{0, 0, 0, 0},
		expected: []int{0, 0, 0, 0},
	},
	{
		name:     "test 7",
		slice:    []int{1},
		expected: []int{1},
	},
	{
		name:     "test 8",
		slice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name:     "test 9",
		slice:    []int{100},
		expected: []int{100},
	},
}

func TestAddElements(t *testing.T) {
	for _, tt := range addElementsTests {
		t.Run(tt.name, func(t *testing.T) {
			testAddElements(t, tt)
		})
	}
}

func TestSliceExample(t *testing.T) {
	for _, tt := range sliceExampleTests {
		t.Run(tt.name, func(t *testing.T) {
			testSliceExample(t, tt)
		})
	}
}

func TestGenerateNewSlice(t *testing.T) {
	for _, tt := range generateNewSliceTests {
		t.Run(tt.name, func(t *testing.T) {
			testGenerateNewSlice(t, tt)
		})
	}
}

func TestCopySlice(t *testing.T) {
	for _, tt := range copySliceTests {
		t.Run(tt.name, func(t *testing.T) {
			testCopySlice(t, tt)
		})
	}
}

func testAddElements(t *testing.T, tt addElementsTest) {
	results := AddElements(tt.slice, tt.values...)

	for i, result := range results {
		if result != tt.expected[i] {
			t.Fatalf("expected %v, got %v", tt.expected, results)
		}
	}
}

func testSliceExample(t *testing.T, tt sliceExampleTest) {
	results := SliceExample(tt.slice)

	for i, result := range results {
		if result != tt.expected[i] {
			t.Fatalf("expected %v, got %v", tt.expected, results)
		}
	}
}

func testGenerateNewSlice(t *testing.T, tt generateNewSliceTest) {
	results := GenerateNewSlice(tt.length)

	for _, result := range results {
		if !(result < 100 && result >= 0) {
			t.Fatalf("expected [0;100), got %v", results)
		}
	}
}

func testCopySlice(t *testing.T, tt copySliceTest) {
	results := CopySlice(tt.slice)

	for i, result := range results {
		if result != tt.expected[i] {
			t.Fatalf("expected %v, got %v", tt.expected[i], result)
		}
	}
	tt.slice[0] = 999 // Modify the original slice
	if reflect.DeepEqual(tt.slice, results) {
		t.Fatalf("expected copy to remain unchanged, but it was modified")
	}
}

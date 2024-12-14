package main

import (
	"testing"
)

type differenceTest struct {
	name     string
	slice1   []string
	slice2   []string
	expected []string
}

var differenceTests = []differenceTest{
	{
		name: "test 1",
		slice1: []string{
			"apple",
			"banana",
			"cherry",
			"date",
			"43",
			"lead",
			"gno1",
		},
		slice2: []string{
			"banana",
			"date",
			"fig",
		},
		expected: []string{
			"apple",
			"cherry",
			"43",
			"lead",
			"gno1",
		},
	},
	{
		name:   "slice1 is empty",
		slice1: []string{},
		slice2: []string{
			"banana", "date", "fig",
		},
		expected: []string{},
	},
	{
		name: "slice2 is empty",
		slice1: []string{
			"apple", "banana", "cherry", "date",
		},
		slice2: []string{},
		expected: []string{
			"apple", "banana", "cherry", "date",
		},
	},
	{
		name: "no common elements",
		slice1: []string{
			"apple", "banana", "cherry", "date",
		},
		slice2: []string{
			"fig", "grape", "kiwi",
		},
		expected: []string{
			"apple", "banana", "cherry", "date",
		},
	},
	{
		name: "all elements from slice1 are in slice2",
		slice1: []string{
			"banana", "date",
		},
		slice2: []string{
			"banana", "date", "fig",
		},
		expected: []string{},
	},
	{
		name: "all elements in slice1 and slice2 are the same",
		slice1: []string{
			"apple", "banana", "cherry",
		},
		slice2: []string{
			"apple", "banana", "cherry",
		},
		expected: []string{},
	},
}

func TestDifference(t *testing.T) {
	for _, tt := range differenceTests {
		t.Run(tt.name, func(t *testing.T) {
			testDifference(t, tt)
		})
	}
}

func testDifference(t *testing.T, tt differenceTest) {
	result := Diffence(tt.slice1, tt.slice2)

	if len(result) != len(tt.expected) {
		t.Errorf("expected %v, got %v", tt.expected, result)
		return
	}

	for i := range result {
		if result[i] != tt.expected[i] {
			t.Errorf("expected %v, got %v", tt.expected, result)
			return
		}
	}
}

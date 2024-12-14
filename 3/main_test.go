package main

import "testing"

type addTest struct {
	name        string
	key         string
	value       int
	mapInstance *StringIntMap
	expected    *StringIntMap
}

type removeTest struct {
	name        string
	key         string
	mapInstance *StringIntMap
	expected    *StringIntMap
}

var addTests = []addTest{
	{
		name:        "test 1",
		key:         "string1",
		value:       2,
		mapInstance: NewStringIntMap(),
		expected: &StringIntMap{
			data: map[string]int{"string1": 2},
		},
	},
	{
		name:        "test 2",
		key:         "string2",
		value:       3,
		mapInstance: NewStringIntMap(),
		expected:    &StringIntMap{data: map[string]int{"string2": 3}},
	},
	{
		name:        "test 3",
		key:         "string1",
		value:       4,
		mapInstance: &StringIntMap{data: map[string]int{"string1": 2}},
		expected:    &StringIntMap{data: map[string]int{"string1": 4}},
	},
	{
		name:        "test 4",
		key:         "string2",
		value:       7,
		mapInstance: &StringIntMap{data: map[string]int{"string1": 2}},
		expected:    &StringIntMap{data: map[string]int{"string1": 2, "string2": 7}},
	},
	{
		name:        "test 5",
		key:         "",
		value:       0,
		mapInstance: &StringIntMap{data: map[string]int{"string1": 2}},
		expected:    &StringIntMap{data: map[string]int{"string1": 2, "": 0}},
	},
	{
		name:        "test 7",
		key:         "string1",
		value:       0,
		mapInstance: &StringIntMap{data: map[string]int{"string1": 2}},
		expected:    &StringIntMap{data: map[string]int{"string1": 0}},
	},
}

var removeTests = []removeTest{
	{
		name: "test 1",
		key:  "string1",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"string1": 1,
				"string2": 2,
				"string3": 3,
			},
		},
		expected: &StringIntMap{
			data: map[string]int{
				"string2": 2,
				"string3": 3,
			},
		},
	},
	{
		name: "Remove non-existing key",
		key:  "string4",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"string1": 1,
				"string2": 2,
				"string3": 3,
			},
		},
		expected: &StringIntMap{
			data: map[string]int{
				"string1": 1,
				"string2": 2,
				"string3": 3,
			},
		},
	},
	{
		name: "Remove the only key",
		key:  "string3",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"string3": 3,
			},
		},
		expected: &StringIntMap{
			data: map[string]int{},
		},
	},
	{
		name: "Remove from empty map",
		key:  "string1",
		mapInstance: &StringIntMap{
			data: map[string]int{},
		},
		expected: &StringIntMap{
			data: map[string]int{},
		},
	},
	{
		name: "Remove key from large map",
		key:  "string10",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"string1":  1,
				"string2":  2,
				"string3":  3,
				"string4":  4,
				"string5":  5,
				"string10": 10,
				"string11": 11,
			},
		},
		expected: &StringIntMap{
			data: map[string]int{
				"string1":  1,
				"string2":  2,
				"string3":  3,
				"string4":  4,
				"string5":  5,
				"string11": 11,
			},
		},
	},
}

func TestAdd(t *testing.T) {
	for _, tt := range addTests {
		t.Run(tt.name, func(t *testing.T) {
			testAdd(t, tt)
		})
	}
}

func TestRemove(t *testing.T) {
	for _, tt := range removeTests {
		t.Run(tt.name, func(t *testing.T) {
			testRemove(t, tt)
		})
	}
}

func testAdd(t *testing.T, tt addTest) {
	tt.mapInstance.Add(tt.key, tt.value)

	for key, value := range tt.expected.data {
		if got := tt.mapInstance.data[key]; got != value {
			t.Errorf("Add() = %v, want %v", got, value)
		}
	}
}

func testRemove(t *testing.T, tt removeTest) {
	tt.mapInstance.Remove(tt.key)

	for key, value := range tt.expected.data {
		if got := tt.mapInstance.data[key]; got != value {
			t.Errorf("Add() = %v, want %v", got, value)
		}
	}
}

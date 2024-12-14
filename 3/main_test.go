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

type copyTest struct {
	name        string
	mapInstance *StringIntMap
	expected    *StringIntMap
}

type existsTest struct {
	name        string
	key         string
	mapInstance *StringIntMap
	expected    bool
}

type getTest struct {
	name        string
	key         string
	mapInstance *StringIntMap
	expected    int
	exists      bool
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

var copyTests = []copyTest{
	{
		name: "copy non-empty",
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
		name: "Copy empty map",
		mapInstance: &StringIntMap{
			data: map[string]int{},
		},
		expected: &StringIntMap{
			data: map[string]int{},
		},
	},
	{
		name: "Copy map with one element",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 42,
			},
		},
		expected: &StringIntMap{
			data: map[string]int{
				"key1": 42,
			},
		},
	},
	{
		name: "Copy map with duplicate values",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 10,
				"key2": 10,
				"key3": 10,
			},
		},
		expected: &StringIntMap{
			data: map[string]int{
				"key1": 10,
				"key2": 10,
				"key3": 10,
			},
		},
	},
}

var existsTests = []existsTest{
	{
		name: "key exists",
		key:  "key1",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
				"key3": 6,
				"key4": 7,
			},
		},
		expected: true,
	},
	{
		name: "Key does not exist in map",
		key:  "key5",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
				"key3": 6,
				"key4": 7,
			},
		},
		expected: false,
	},
	{
		name: "Key does not exist in map",
		key:  "key5",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
				"key3": 6,
				"key4": 7,
			},
		},
		expected: false,
	},
	{
		name: "Key exists with value zero",
		key:  "key1",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 0,
				"key2": 4,
			},
		},
		expected: true,
	},
	{
		name: "Key does not exist but map has other keys",
		key:  "key999",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
				"key3": 6,
			},
		},
		expected: false,
	},
}

var getTests = []getTest{
	{
		name: "key exists",
		key:  "key1",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
				"key3": 6,
			},
		},
		expected: 2,
		exists:   true,
	},
	{
		name: "Key does not exist",
		key:  "key5",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
				"key3": 6,
			},
		},
		expected: 0,
		exists:   false,
	},
	{
		name: "Key exists with zero value",
		key:  "key2",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 0,
				"key3": 6,
			},
		},
		expected: 0,
		exists:   true,
	},
	{
		name: "Empty map",
		key:  "key1",
		mapInstance: &StringIntMap{
			data: map[string]int{},
		},
		expected: 0,
		exists:   false,
	},
	{
		name: "Check another key in non-empty map",
		key:  "key3",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
				"key3": 6,
			},
		},
		expected: 6,
		exists:   true,
	},
	{
		name: "Key does not exist but map has other keys",
		key:  "key999",
		mapInstance: &StringIntMap{
			data: map[string]int{
				"key1": 2,
				"key2": 4,
			},
		},
		expected: 0,
		exists:   false,
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

func TestCopy(t *testing.T) {
	for _, tt := range copyTests {
		t.Run(tt.name, func(t *testing.T) {
			testCopy(t, tt)
		})
	}
}

func TestExists(t *testing.T) {
	for _, tt := range existsTests {
		t.Run(tt.name, func(t *testing.T) {
			testExists(t, tt)
		})
	}
}

func TestGet(t *testing.T) {
	for _, tt := range getTests {
		t.Run(tt.name, func(t *testing.T) {
			testGet(t, tt)
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
			t.Errorf("Remove() = %v, want %v", got, value)
		}
	}
}

func testCopy(t *testing.T, tt copyTest) {
	result := tt.mapInstance.Copy()
	tt.mapInstance.data["string1"] = 10
	for key, value := range tt.expected.data {
		if got := result[key]; got != value {
			t.Errorf("Copy() = %v, want %v", got, value)
		}
	}
}

func testExists(t *testing.T, tt existsTest) {
	result := tt.mapInstance.Exists(tt.key)
	if result != tt.expected {
		t.Errorf("expected %t, got %t", tt.expected, result)
	}
}

func testGet(t *testing.T, tt getTest) {
	value, exists := tt.mapInstance.Get(tt.key)

	if exists != tt.exists {
		t.Errorf("expected %t, got %t", tt.exists, exists)
	}

	if value != tt.expected {
		t.Errorf("expected %d, got %d", tt.expected, value)
	}
}

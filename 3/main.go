package main

import "fmt"

// StringIntMap represents a map that associates string keys with integer values.
type StringIntMap struct {
	data map[string]int
}

// NewStringIntMap creates and returns a new StringIntMap instance.
func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

// Add inserts or updates a key-value pair in the map.
func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

// Remove deletes the key-value pair associated with the given key.
func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

// Copy creates a new map that is a copy of the current map.
func (m *StringIntMap) Copy() map[string]int {
	if m.data == nil {
		return nil
	}
	result := make(map[string]int, len(m.data))

	for key, value := range m.data {
		result[key] = value
	}

	return result
}

// Exists checks if the given key exists in the map.
// It returns true if the key is found, false otherwise.
func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]

	return exists
}

// Get retrieves the value associated with the given key.
// It returns the value and a boolean indicating if the key exists.
func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]

	return value, exists
}

func main() {
	m := NewStringIntMap()
	m.Add("string 1", 1)
	m.Add("string 2", 2)
	m.Add("string 3", 3)
	m.Remove("string 1")

	fmt.Println(m)

	copy := m.Copy()
	m.Remove("string 3")
	fmt.Println(m)
	fmt.Println(copy)

	exist := m.Exists("string 4")

	fmt.Println(exist)

	get, ok := m.Get("string 2")
	fmt.Println(get, ok)
}

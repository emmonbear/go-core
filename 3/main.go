package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{
		data: make(map[string]int),
	}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	if m.data == nil {
		return nil
	}
	result := make(map[string]int, len(m.data))

	for key, value :=  range m.data {
		result[key] = value
	}

	return result
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.data[key]

	return exists
}

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

	exist :=m.Exists("string 4")

	fmt.Println(exist)

	get, ok := m.Get("string 3")
	fmt.Println(get, ok)


}

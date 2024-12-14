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

func main() {
	m := NewStringIntMap()
	m.Add("string 1", 1)
	m.Add("string 2", 2)
	m.Remove("string 1")

	fmt.Println(m)
}

package main

import "testing"

type generateRandomNumberTest struct {
	name  string
	count int
}

var generateRandomNumberTests = []generateRandomNumberTest{
	{
		name:  "test 1",
		count: 10,
	},
	{
		name:  "test 2",
		count: 5,
	},
	{
		name:  "test 3",
		count: 0,
	},
}

func TestGenerateRandomNumberTest(t *testing.T) {
	for _, tt := range generateRandomNumberTests {
		t.Run(tt.name, func(t *testing.T) {
			testGenerateRandomNumberTest(t, tt)
		})
	}
}

func testGenerateRandomNumberTest(t *testing.T, tt generateRandomNumberTest) {
	ch := make(chan int)
	go generateRandomNumber(ch, tt.count)

	count := 0
	for num := range ch {
		count++
		if num < minValue || num > maxValue {
			t.Errorf("generated number %d is out of range", num)
		}
	}

	if count != tt.count {
		t.Errorf("expected %d numbers, got %d", tt.count, count)
	}

}

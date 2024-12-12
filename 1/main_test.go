package main

import "testing"

type valuesToStringTest struct {
	name         string
	values       *Values
	expected     string
	expectedRune []rune
}

var valuesToStringTests = []valuesToStringTest{
	{
		name: "default",
		values: &Values{
			numDecimal:     42,
			numOctal:       052,
			numHexadecimal: 0x2A,
			pi:             3.14,
			name:           "Golang",
			isActive:       true,
			complexNum:     1 + 2i,
		},
		expected: "4242423.14Golangtrue(1+2i)",
		expectedRune: []rune{52, 50, 52, 50, 52, 50, 51, 46, 49, 52, 71, 111, 108, 97,
			110, 103, 116, 114, 117, 101, 40, 49, 43, 50, 105, 41},
	},
	{
		name: "default",
		values: &Values{
			numDecimal:     1242,
			numOctal:       0212352,
			numHexadecimal: 0x2AFFAB,
			pi:             .32,
			name:           "Test123",
			isActive:       false,
			complexNum:     1.231 + 2.31i,
		},
		expected: "12427089028179630.32Test123false(1.231+2.31i)",
		expectedRune: []rune{49, 50, 52, 50, 55, 48, 56, 57, 48, 50, 56, 49, 55, 57, 54, 51, 48,
			46, 51, 50, 84, 101, 115, 116, 49, 50, 51, 102, 97, 108, 115, 101, 40, 49, 46, 50,
			51, 49, 43, 50, 46, 51, 49, 105, 41},
	},
}

func TestValuesToString(t *testing.T) {
	for _, tt := range valuesToStringTests {
		t.Run(tt.name, func(t *testing.T) {
			testValuesToString(t, tt)
		})
	}

}

func testValuesToString(t *testing.T, tt valuesToStringTest) {
	res := tt.values.ValuesToString()
	r := tt.values.ValuesToRune()
	if res != tt.expected {
		t.Fatalf("expected %s, got %v", tt.expected, res)
	}

	for i, rr := range r {
		if rr != tt.expectedRune[i] {
			t.Fatalf("expected %d, got %d", tt.expectedRune[i], r)
		}
	}
}

package main

import (
	"fmt"
)

type Values struct {
	numDecimal     int
	numOctal       int
	numHexadecimal int
	pi             float64
	name           string
	isActive       bool
	complexNum     complex64
}

func (v *Values) String() string {
	return fmt.Sprintf(
		"var numDecimal %T = %d\t// Десятичная система\n"+
			"var numOctal %T = %d\t// Восьмеричная система\n"+
			"var numHexadecimal %T = %d\t// Шестнадцатиричная система\n"+
			"var pi %T = %f\t// Тип float64\n"+
			"var name %T = %s\t// Тип string\n"+
			"var isActive %T = %t\t// Тип bool\n"+
			"var complexNum %T = %g\t// Тип complex64\n",
		v.numDecimal, v.numDecimal,
		v.numOctal, v.numOctal,
		v.numHexadecimal, v.numHexadecimal,
		v.pi, v.pi,
		v.name, v.name,
		v.isActive, v.isActive,
		v.complexNum, v.complexNum,
	)
}

func main() {
	val := Values{
		numDecimal:     42,
		numOctal:       052,
		numHexadecimal: 0x2A,
		pi:             3.14,
		name:           "Golang",
		isActive:       true,
		complexNum:     1 + 2i,
	}

	fmt.Println(&val)

}

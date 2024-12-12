package main

import (
	"crypto/sha256"
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

const salt = "go-2024"

func (v *Values) String() string {
	return fmt.Sprintf(
		"var numDecimal %T = %d\t// Десятичная система\n"+
			"var numOctal %T = %d\t// Восьмеричная система\n"+
			"var numHexadecimal %T = %d\t// Шестнадцатиричная система\n"+
			"var pi %T = %.2f\t// Тип float64\n"+
			"var name %T = %s\t// Тип string\n"+
			"var isActive %T = %t\t// Тип bool\n"+
			"var complexNum %T = %g\t// Тип complex64",
		v.numDecimal, v.numDecimal,
		v.numOctal, v.numOctal,
		v.numHexadecimal, v.numHexadecimal,
		v.pi, v.pi,
		v.name, v.name,
		v.isActive, v.isActive,
		v.complexNum, v.complexNum,
	)
}

func (v *Values) ValuesToString() string {
	return fmt.Sprintf("%d%d%d%.2f%s%t%g",
		v.numDecimal,
		v.numOctal,
		v.numHexadecimal,
		v.pi,
		v.name,
		v.isActive,
		v.complexNum,
	)
}

func (v *Values) ValuesToRune() []rune {
	return []rune(v.ValuesToString())
}

func (v *Values) HashWithSalt() string {
	runes := v.ValuesToRune()
	saltRune := []rune(salt)
	halfLength := len(runes) / 2

	runes = append(runes[:halfLength], append(saltRune, runes[halfLength:]...)...)
	hash := sha256.Sum256([]byte(string(runes)))

	return fmt.Sprintf("%X", hash)
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

	res := val.ValuesToString()

	fmt.Println(res)

	r := val.ValuesToRune()

	for _, rr := range r {
		fmt.Printf("%d,", rr)
	}
	println()
	hash := val.HashWithSalt()

	fmt.Println(hash)
}

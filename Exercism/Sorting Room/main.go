package main

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

type numberBoxContaining struct {
	number int
}

func (box numberBoxContaining) Number() int {
	return box.number
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

type AnotherFancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

func (i AnotherFancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	if _, ok := fnb.(FancyNumber); ok {
		num, err := strconv.Atoi(fnb.Value())
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}

func main() {
	fmt.Println(DescribeNumber(-12.345))
	fmt.Println(DescribeNumberBox(numberBoxContaining{12}))
	fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))
	// Output: 10
	fmt.Println(ExtractFancyNumber(AnotherFancyNumber{"4"}))
	// Output: 0
}

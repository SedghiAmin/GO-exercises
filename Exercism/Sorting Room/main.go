package main

import "fmt"

type NumberBox interface {
	Number() int
}

type numberBoxContaining struct {
	number int
}

func (box numberBoxContaining) Number() int {
	return box.number
}

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

func main() {
	fmt.Println(DescribeNumber(-12.345))
	fmt.Println(DescribeNumberBox(numberBoxContaining{12}))
}

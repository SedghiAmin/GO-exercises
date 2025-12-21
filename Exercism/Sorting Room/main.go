package main

import "fmt"

func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

func main() {
	fmt.Printf(DescribeNumber(-12.345))
}

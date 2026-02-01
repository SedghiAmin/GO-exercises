package main

import "fmt"

func Reverse(input string) string {
	r := make([]rune, len(input))
	step := len(input) - 1
	for _, char := range input {
		r[step] = char
		step--
	}
	return string(r)
}

func main() {
	input := "AMIN"
	fmt.Println(Reverse(input))
}

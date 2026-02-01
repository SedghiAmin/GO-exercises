package main

import "fmt"

func Reverse(input string) string {
	runes := []rune(input)
	r := make([]rune, len(runes))
	step := len(runes) - 1
	for _, char := range runes {
		r[step] = char
		step--
	}
	return string(r)
}

func main() {
	input := "AMIN"
	fmt.Println(Reverse(input))
}

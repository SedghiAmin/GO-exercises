package main

import (
	"fmt"
	"strings"
)

func IsPangram(input string) bool {
	alphabet := make(map[rune]bool)

	for _, char := range strings.ToUpper(input) {
		if char >= 'A' && char <= 'Z' {
			alphabet[char] = true
		}
	}

	return len(alphabet) == 26
}

func main() {
	input := "a quick movement of the enemy will jeopardize five gunboats"
	fmt.Println(IsPangram(input)) //flase

	input = "The quick brown fox jumps over the lazy dog."
	fmt.Println(IsPangram(input)) //true
}

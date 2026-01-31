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

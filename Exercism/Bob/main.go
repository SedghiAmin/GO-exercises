package main

import (
	"unicode"
)

func haveletter(remark string) bool {
	input := []rune(remark)
	for _, char := range input {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false
}

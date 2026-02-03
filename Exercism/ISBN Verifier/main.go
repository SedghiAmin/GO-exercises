package main

import (
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	sum := 0
	if len(isbn) != 10 {
		return false
	} else {
		var digit int
		step := 10
		for i := 0; i < 10; i++ {
			if i == 9 && (isbn[i] == 'X' || isbn[i] == 'x') {
				digit = 10
			} else {
				if !unicode.IsDigit(rune(isbn[i])) {
					return false
				} else {
					digit = int(isbn[i] - '0')
				}

			}
			sum += digit * step
			step--
		}
	}
	return sum%11 == 0
}

package main

import (
	"fmt"
	"slices"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, fmt.Errorf("span must be smaller than string length")
	}
	if span < 0 {
		return 0, fmt.Errorf("span must not be negative")
	}
	var sum int64
	result := make([]int64, len(digits))
	sum = 1
	digitsRunes := []rune(digits)
	for i := 0; i <= len(digitsRunes)-span; i++ {
		for j := i; j < i+span; j++ {
			if unicode.IsDigit(digitsRunes[j]) {
				sum *= int64(digitsRunes[j] - '0')
			} else {
				return 0, fmt.Errorf("digits input must only contain digits")
			}
		}
		result[i] = sum
		sum = 1
	}
	return slices.Max(result), nil
}

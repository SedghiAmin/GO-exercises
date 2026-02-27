package romannumerals

import (
	"errors"
	"strings"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("the number must be between 1 to 3999")
	}
	in := input
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	output := make([]string, 0)

	for i, num := range values {
		if num <= in {
			for num <= in {
				output = append(output, symbols[i])
				in -= num
			}
		}
	}
	return strings.Join(output, ""), nil
}

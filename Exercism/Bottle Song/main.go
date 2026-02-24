package main

import (
	"fmt"
	"strings"
	"unicode"
)

func numToWord(num int) string {
	switch num {
	case 1:
		return "one green bottle"
	case 2:
		return "two green bottles"
	case 3:
		return "three green bottles"
	case 4:
		return "four green bottles"
	case 5:
		return "five green bottles"
	case 6:
		return "six green bottles"
	case 7:
		return "seven green bottles"
	case 8:
		return "eight green bottles"
	case 9:
		return "nine green bottles"
	case 10:
		return "ten green bottles"
	}
	return ""
}

func CapitalizeFirst(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	var builder strings.Builder
	builder.Grow(len(s))

	builder.WriteString(string(unicode.ToUpper(r[0])))

	if len(r) > 1 {
		builder.WriteString(string(r[1:]))
	}

	return builder.String()
}

func Recite(startBottles, takeDown int) []string {
	if startBottles < 1 || takeDown < 1 {
		return []string{}
	}
	if startBottles < takeDown {
		return []string{}
	}
	result := make([]string, 0, startBottles-takeDown+1)
	for i := startBottles; i > startBottles-takeDown; i-- {
		result = append(result, fmt.Sprintf("%v hanging on the wall,", CapitalizeFirst(numToWord(i))))
		result = append(result, fmt.Sprintf("%v hanging on the wall,", CapitalizeFirst(numToWord(i))))
		result = append(result, fmt.Sprintf("And if one green bottle should accidentally fall,"))
		if i == 1 {
			result = append(result, fmt.Sprintf("There'll be no green bottles hanging on the wall."))
		} else {
			result = append(result, fmt.Sprintf("There'll be %v hanging on the wall.", numToWord(i-1)))
		}
		if i-1 > startBottles-takeDown {
			result = append(result, fmt.Sprintf(""))
		}

	}
	return result
}

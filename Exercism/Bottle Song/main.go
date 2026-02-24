package main

import (
	"fmt"
	"strings"
)

func NumToWord(num int) string {
	switch num {
	case 1:
		return "One green bottle"
	case 2:
		return "Two green bottles"
	case 3:
		return "Three green bottles"
	case 4:
		return "Four green bottles"
	case 5:
		return "Five green bottles"
	case 6:
		return "Six green bottles"
	case 7:
		return "Seven green bottles"
	case 8:
		return "Eight green bottles"
	case 9:
		return "Nine green bottles"
	case 10:
		return "Ten green bottles"
	}
	return ""
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
		result = append(result, fmt.Sprintf("%v hanging on the wall,", NumToWord(i)))
		result = append(result, fmt.Sprintf("%v hanging on the wall,", strings.ToLower(NumToWord(i))))
		result = append(result, fmt.Sprintf("And if one green bottle should accidentally fall,"))
		if i == 1 {
			result = append(result, fmt.Sprintf("There'll be no green bottles hanging on the wall."))
		} else {
			result = append(result, fmt.Sprintf("There'll be %v hanging on the wall.", NumToWord(i-1)))
		}
		if i-1 > startBottles-takeDown {
			result = append(result, fmt.Sprintf(""))
		}

	}
	return result
}

func main() {
	fmt.Println(Recite(5, 3))
}

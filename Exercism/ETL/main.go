package main

import (
	"fmt"
	"strings"
)

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int)
	for score, str := range in {
		for _, char := range str {
			lChar := strings.ToLower(char)
			out[lChar] = score
		}
	}
	return out
}

func main() {

	oldEnglishScores := map[int][]string{
		1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	}

	newEnglishScores := Transform(oldEnglishScores)

	fmt.Println("Score for 'c':", newEnglishScores["c"]) // 3
	fmt.Println("Score for 'z':", newEnglishScores["z"]) // 10
}

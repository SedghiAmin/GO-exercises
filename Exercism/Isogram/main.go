package main

import (
	"fmt"
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToUpper(word)
	m := make(map[rune]bool, len(word))
	for _, c := range word {
		if c == ' ' || c == '-' {
			continue
		}
		if _, exist := m[c]; exist {
			return false
		}
		m[c] = true
	}
	return true
}

func main() {
	fmt.Println(IsIsogram("lumberjacks"))
	fmt.Println(IsIsogram("background"))
	fmt.Println(IsIsogram("downstream"))
	fmt.Println(IsIsogram("six-year-old"))
	fmt.Println(IsIsogram("isograms"))
}

package main

import (
	"fmt"
	"strings"
)

func Score(word string) int {
	sum := 0
	word = strings.ToUpper(word)
	if len(strings.ToLower(word)) < 1 {
		return sum
	}
	for _, c := range strings.ToLower(word) {
		switch c {
		case 'd', 'g':
			sum += 2
		case 'b', 'c', 'm', 'p':
			sum += 3
		case 'f', 'h', 'v', 'w', 'y':
			sum += 4
		case 'k':
			sum += 5
		case 'j', 'x':
			sum += 8
		case 'q', 'z':
			sum += 10
		default:
			sum += 1
		}
	}
	return sum
}

func main() {
	fmt.Println(Score("hello"))   // 8
	fmt.Println(Score("world"))   // 9
	fmt.Println(Score(""))        // 0
	fmt.Println(Score("quirky"))  // 22
	fmt.Println(Score("cabbage")) // 14
}

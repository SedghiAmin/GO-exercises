package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	result := Frequency{}
	reg := regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z0-9]+)?`)
	words := reg.FindAllString(strings.ToLower(phrase), -1)
	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	testCases := []struct {
		name string
		text string
	}{
		{"simple", "Hello World"},
		{"punctuations", "Hello, World! How are you?"},
		{"contractions", "I don't know what you're saying"},
		{"numbers", "I have 2 apples and 3 bananas"},
		{"all", "That's 100% correct! Don't you agree?"},
	}

	for _, tc := range testCases {
		words := WordCount(tc.text)

		fmt.Printf("\n%s:\n", tc.name)
		fmt.Printf("Text: %q\n", tc.text)
		fmt.Printf("Words: %v\n", words)
		fmt.Printf("Counts: %d\n", len(words))
	}
}

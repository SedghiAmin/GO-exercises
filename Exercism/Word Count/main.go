package main

import (
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

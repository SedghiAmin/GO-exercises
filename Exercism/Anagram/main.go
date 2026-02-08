package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func NormalizeWords(str string) string {
	lowered := strings.Map(unicode.ToLower, str)
	runes := []rune(lowered)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func Detect(subject string, candidates []string) []string {
	NormalizeSubject := NormalizeWords(subject)
	anagrams := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		if strings.EqualFold(subject, candidate) {
			continue
		}
		if NormalizeWords(candidate) == NormalizeSubject {
			anagrams = append(anagrams, candidate)
		}
	}
	return anagrams
}

func main() {
	target := "stone"
	candidates := []string{"stone", "tones", "banana", "tons", "notes", "Seton"}

	result := Detect(target, candidates)
	fmt.Println(result) // output: [tones notes Seton]
}

package main

import (
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

package anagram

import (
	"sort"
	"strings"
	"unicode"
)

func NormalizeWord(str string) string{
    lowStr:= strings.Map(unicode.ToLower, str)
    runes:= []rune(lowStr)
    sort.Slice(runes, func(i,j int)bool{
        return runes[i] < runes[j]
    })
    return string(runes)
}

func Detect(subject string, candidates []string) []string {
	normalizeSubject:= NormalizeWord(subject)
    anagram:= make([]string, 0 , len(candidates))
    for _, candidate:= range candidates{
        if strings.EqualFold(candidate , subject){
            continue
        }
        if NormalizeWord(candidate) == normalizeSubject{
            anagram= append(anagram, candidate)
        }
    }
    return anagram
}

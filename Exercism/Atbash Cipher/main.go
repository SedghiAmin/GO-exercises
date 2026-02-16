package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var output strings.Builder
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)
	groupByFive := 0
	for _, c := range s {
		if unicode.IsLetter(c) {
			offset := c - 'a'
			output.WriteRune('z' - offset)
			groupByFive++
		} else if unicode.IsDigit(c) {
			output.WriteRune(c)
			groupByFive++
		}
	}
	if output.String() == "" {
		return ""
	}
	var result strings.Builder
	for i, c := range output.String() {
		if i > 0 && i%5 == 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(c)
	}
	return result.String()
}
func main() {
	fmt.Println(Atbash("test"))                                         //gvhg
	fmt.Println(Atbash("x123 yes"))                                     //c123b vh
	fmt.Println(Atbash("gvhg"))                                         //test
	fmt.Println(Atbash("hello world"))                                  //svool dliow
	fmt.Println(Atbash("Testing,1 2 3, testing."))                      //gvhgr mt123 gvhgr mt
	fmt.Println(Atbash("The quick brown fox jumps over the lazy dog.")) //gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt
}

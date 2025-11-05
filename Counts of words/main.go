package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)

	count := make(map[string]int)

	for _ , value := range(words){
		count[value] ++
	}

	return count
}

func main(){
	wc.Test(WordCount)
}
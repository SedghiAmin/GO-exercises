package main

import(
	"strings"
	"fmt"
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
	var s = "I ate a donut. Then I ate another donut."
	fmt.Println(WordCount(s))
}
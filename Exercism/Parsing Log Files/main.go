package main

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[*-=~]*>`)
	return re.Split(text, -1)
}

func main() {
	fmt.Println(IsValidLine("[ERR] A good error here"))
	fmt.Printf("%#v", SplitLogLine("section 1<*>section 2<~~~>section 3"))
}

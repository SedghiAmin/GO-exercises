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

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*?password[^"]*?"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(IsValidLine("[ERR] A good error here"))
	fmt.Printf("%#v\n", SplitLogLine("section 1<*>section 2<~~~>section 3"))

}

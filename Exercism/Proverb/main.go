package main

import "fmt"

func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return rhyme
	}
	r := make([]string, len(rhyme))
	for i := 0; i < len(rhyme)-1; i++ {
		r[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
	}
	r[len(rhyme)-1] = "And all for the want of a " + rhyme[0] + "."
	return r
}

func main() {
	result := Proverb([]string{"nail", "shoe", "horse"})

	for i, line := range result {
		fmt.Printf("Line %d: %s\n", i+1, line)
	}
}

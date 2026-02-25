package main

import (
	"fmt"
	"strings"
)

func Verse(i int) string {
	if i < 1 {
		return ""
	}
	days := []string{
		"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth", "thirteenth", "fourteenth",
	}

	gifts := []string{
		"a Partridge in a Pear Tree", "two Turtle Doves", "three French Hens",
		"four Calling Birds", "five Gold Rings", "six Geese-a-Laying",
		"seven Swans-a-Swimming", "eight Maids-a-Milking", "nine Ladies Dancing",
		"ten Lords-a-Leaping", "eleven Pipers Piping", "twelve Drummers Drumming",
	}

	result := strings.Builder{}
	tmp := fmt.Sprintf("On the %v day of Christmas my true love gave to me: ", days[i-1])
	result.WriteString(tmp)
	for j := i - 1; j >= 0; j-- {
		tmp = fmt.Sprintf(gifts[j])
		result.WriteString(tmp)
		if j-1 == 0 {
			result.WriteString(", and ")
		}
		if j-1 > 0 {
			result.WriteString(", ")
		}
	}
	result.WriteString(".")
	return result.String()
}

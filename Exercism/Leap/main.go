package main

import "fmt"

func IsLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsLeapYear(2024))
}

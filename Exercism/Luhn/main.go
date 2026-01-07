package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}
	j := 1
	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		charStr := string(id[i])
		v, err := strconv.Atoi(charStr)
		if err != nil {
			return false
		}
		if j%2 == 0 {
			if v *= 2; v > 9 {
				v -= 9
			}
		}
		sum += v
		j++
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func main() {

	fmt.Println(Valid("066 123 478"))                                 //false
	fmt.Println(Valid("4539 3195 0343 6467"))                         //true
	fmt.Println(Valid("9999999999 9999999999 9999999999 9999999999")) //true
}

package main

import (
	"errors"
	"fmt"
)

func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("incorrect number")
	}
	if n == 1 {
		return 2, nil
	}
	i := 2
	prime := i
	num := 3
	p := 0
	for i <= n {
		for j := 2; j < num; j++ {
			if num%j == 0 {
				p++
				break
			}
		}
		if p == 0 {
			prime = num
			i++
		}
		num++
		p = 0
	}
	return prime, nil
}

func main() {
	fmt.Println(Nth(1)) // 2
	fmt.Println(Nth(2)) // 3
	fmt.Println(Nth(3)) // 5
	fmt.Println(Nth(4)) // 7
	fmt.Println(Nth(5)) // 11
	fmt.Println(Nth(6)) // 13
	fmt.Println(Nth(0)) // 0, errors:incorrect number
}

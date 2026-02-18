package main

import "fmt"

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	marked := make(map[int]bool, limit)
	prime := make([]int, 0)
	for i := 2; i <= limit; i++ {
		if !marked[i] {
			prime = append(prime, i)
			for j := i * i; j <= limit; j += i {
				marked[j] = true
			}
		}
	}
	return prime
}

func main() {
	fmt.Println(Sieve(1000))
}

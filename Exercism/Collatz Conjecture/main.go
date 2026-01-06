package main

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n > 0 {
		step := 0
		for n != 1 {
			if n%2 == 0 {
				n /= 2
			} else {
				n = n*3 + 1
			}
			step++
		}
		return step, nil
	} else {
		return 0, errors.New("the number is invalid")
	}
}

func main() {
	tests := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 19, 27}

	for _, n := range tests {
		steps, err := CollatzConjecture(n)
		if err != nil {
			fmt.Printf("Error for %d: %v\n", n, err)
		} else {
			fmt.Printf("Collatz(%d) = %d steps\n", n, steps)
		}
	}
}

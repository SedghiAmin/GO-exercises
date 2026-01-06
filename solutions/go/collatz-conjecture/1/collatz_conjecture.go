package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n > 0 {
		step := 0
		for n != 1{
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
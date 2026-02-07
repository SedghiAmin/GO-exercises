package main

import "fmt"

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	//var k Kind
	if (a+b < c || b+c < a || a+c < b) || (a <= 0 || b <= 0 || c <= 0) {
		return NaT
	} else {
		z := 1
		if a == b {
			z++
		}
		if b == c {
			z++
		}
		if a == c {
			z++
		}
		if z >= 3 {
			return Equ
		}
		if z >= 2 {
			return Iso
		}
	}
	return Sca
}
func main() {

	fmt.Println("Test 1 (2, 2, 2):", KindFromSides(2, 2, 2))
	fmt.Println("Test 2 (3, 4, 4):", KindFromSides(3, 4, 4))
	fmt.Println("Test 3 (3, 4, 5):", KindFromSides(3, 4, 5))
	fmt.Println("Test 4 (0, 4, 5):", KindFromSides(0, 4, 5))
	fmt.Println("Test 5 (1, 1, 3):", KindFromSides(1, 1, 3))
}

package main

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

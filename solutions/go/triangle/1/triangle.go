// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
// type Kind

type Kind int

const (
	NaT Kind = iota
	Equ
	Iso
	Sca
)

func KindFromSides(a, b, c float64) Kind {
	//var k Kind
	if (a+b < c || b+c < a || a+c < b) || (a <= 0 || b<=0 || c<=0) {
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

package main

import (
	"errors"
	"fmt"
)

type Histogram map[byte]int

type DNA string

func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for i := 0; i < len(d); i++ {
		switch d[i] {
		case 'A', 'C', 'G', 'T':
			h[d[i]]++
		default:
			return nil, errors.New(" invalid nucleotides")
		}
	}
	return h, nil
}

func main() {

	strand := DNA("ACGTACGT")
	histogram, err := strand.Counts()

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Results:")
		fmt.Printf("A: %d\n", histogram['A'])
		fmt.Printf("C: %d\n", histogram['C'])
		fmt.Printf("G: %d\n", histogram['G'])
		fmt.Printf("T: %d\n", histogram['T'])
	}

	invalidStrand := DNA("ACGTXACGT")
	_, err2 := invalidStrand.Counts()
	if err2 != nil {
		fmt.Println("Invalid string error:", err2)
	}
}

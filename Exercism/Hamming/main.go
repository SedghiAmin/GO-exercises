package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("not equal length")
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}

func main() {
	// Example from the problem statement
	strand1 := "GAGCCTACTAACGGGAT"
	strand2 := "CATCGTAATGACGGCCT"

	distance, err := Distance(strand1, strand2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Hamming distance between\n%q\nand\n%q\nis: %d\n",
			strand1, strand2, distance)
	}
}

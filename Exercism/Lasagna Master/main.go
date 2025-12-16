package main

import "fmt"

func PreparationTime(layers []string, prepartionTime int) int {
	if prepartionTime > 0 {
		return len(layers) * prepartionTime
	}
	return len(layers) * 2
}

func Quantities(layers []string) (int, float64) {
	var s, n int
	for _, l := range layers {
		if l == "sauce" {
			s += 1
		} else if l == "noodles" {
			n += 1
		}
	}
	return n * 50, float64(s) * 0.2
}

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	// => 18
	fmt.Println(PreparationTime(layers, 0))
	// => 12
	fmt.Println(Quantities(layers))
	// => 100, 0.4
}

package main

import "fmt"

func PreparationTime(layers []string, prepartionTime int) int {
	if prepartionTime > 0 {
		return len(layers) * prepartionTime
	}
	return len(layers) * 2
}

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	// => 18
	fmt.Println(PreparationTime(layers, 0))
	// => 12

}

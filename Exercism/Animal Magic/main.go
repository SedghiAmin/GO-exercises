package main

import (
	"fmt"
	"math/rand"
	"time"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * 12
}

func main() {
	fmt.Println(RollADie())           // d will be assigned a random int, 1 <= d <= 20
	fmt.Println(GenerateWandEnergy()) // f will be assigned a random float64, 0.0 <= f < 12.0
	fmt.Println(ShuffleAnimals())
}

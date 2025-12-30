package main

import (
	"fmt"
	"math/rand"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(20) + 1
}

func main() {
	fmt.Println(RollADie()) // d will be assigned a random int, 1 <= d <= 20
}

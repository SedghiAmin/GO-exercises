package main

import (
	"fmt"
)

// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
	if !knightIsAwake {
		return true
	}
	return false
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	}
	return false
}

func main() {
	var knightIsAwake = false
	var archerIsAwake = true
	var prisonerIsAwake = false
	fmt.Printf("Can Spy: %v\n", CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Printf("Can Fast Attack: %v\n", CanFastAttack(knightIsAwake))
}

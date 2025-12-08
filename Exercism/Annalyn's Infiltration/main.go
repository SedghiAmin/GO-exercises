package main

import (
	"fmt"
)

func CanFastAttack(knightIsAwake bool) bool {
	if !knightIsAwake {
		return true
	}
	return false
}

func main() {
	var knightIsAwake = true
	fmt.Println(CanFastAttack(knightIsAwake))
}

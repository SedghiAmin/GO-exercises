package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index >= len(slice) {
		return -1 //If the index is out of bounds (ie. if it is negative or after the end of the stack), return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

func main() {
	fmt.Printf("%#v\n", FavoriteCards())
	fmt.Printf("%#v\n", GetItem([]int{1, 2, 4, 1}, 10))
	fmt.Printf("%#v\n", SetItem([]int{1, 2, 4, 1}, -1, 6))
}

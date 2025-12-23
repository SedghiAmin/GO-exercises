package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	f, exist := cb[file]
	if !exist || len(f) != 8 {
		return 0
	}
	count := 0
	for _, val := range f {
		if val {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	count := 0
	index := rank - 1
	for _, file := range cb {
		if file[index] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, file := range cb {
		count += len(file)
	}
	return count
}

func main() {
	cb := Chessboard{
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{false, false, false, false, true, false, false, false},
		"C": File{false, false, false, false, false, false, false, false},
	}
	fmt.Println(CountInFile(cb, "A")) // Output: 3
	fmt.Println(CountInFile(cb, "B")) // Output: 1
	fmt.Println(CountInFile(cb, "Z")) // Output: 0 (file doesn't exist)

	fmt.Println(CountInRank(cb, 1)) // Output: 1 (only file A, rank 1)
	fmt.Println(CountInRank(cb, 3)) // Output: 1 (files A and B, rank 3)
	fmt.Println(CountInRank(cb, 9)) // Output: 0 (invalid rank)
	fmt.Println(CountInRank(cb, 0)) // Output: 0 (invalid rank)

	fmt.Println(CountAll(cb)) // Output: 24
}

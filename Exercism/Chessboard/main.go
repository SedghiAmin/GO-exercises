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

func main() {
	cb := Chessboard{
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{false, false, false, false, true, false, false, false},
	}
	fmt.Println(CountInFile(cb, "A")) // Output: 3
	fmt.Println(CountInFile(cb, "B")) // Output: 1
	fmt.Println(CountInFile(cb, "Z")) // Output: 0 (file doesn't exist)
}

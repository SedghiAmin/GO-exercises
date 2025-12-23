# Chessboard Occupancy Counter (Go)

This project is a simple Go program that models a chessboard and provides utility functions to count squares and occupied positions by file (column), rank (row), or across the entire board.

It is designed as a practice exercise for working with:
- Custom types
- Maps and slices
- Iteration and validation logic in Go

# Chessboard Representation

The chessboard is represented using two custom types:

type File []bool
type Chessboard map[string]File


- File: A slice of 8 booleans where each value represents whether a square in that file is occupied.
- Chessboard: A map with keys "A" to "H" (files), each mapped to a File. 
 
Each bool value means:

- true → square is occupied
- false → square is empty

# Functions

CountInFile
func CountInFile(cb Chessboard, file string) int


Counts how many squares are occupied in a given file (column).

Rules:
- If the file does not exist, returns 0
- If the file does not contain exactly 8 squares, returns 0

CountInRank
func CountInRank(cb Chessboard, rank int) int


Counts how many squares are occupied in a given rank (row).

Rules:
- Valid ranks are from 1 to 8
- If the rank is invalid, returns 0

CountAll
func CountAll(cb Chessboard) int


Counts the total number of squares present in the chessboard, regardless of whether they are occupied.

Note: This function counts squares based on the actual data present in the map, not a fixed 8×8 board.

CountOccupied
func CountOccupied(cb Chessboard) int


Counts how many squares are occupied (true) across the entire chessboard.

Example Usage
```go
cb := Chessboard{
"A": File{true, false, true, false, false, false, false, true},
"B": File{false, false, false, false, true, false, false, false},
"C": File{false, false, false, false, false, false, false, false},
}

fmt.Println(CountInFile(cb, "A")) // 3
fmt.Println(CountInRank(cb, 1))   // 1
fmt.Println(CountAll(cb))         // 24
fmt.Println(CountOccupied(cb))    // 4
```

# Key Concepts Practiced
- Custom type definitions
- Maps with slice values
- Defensive programming (input validation)
- Iteration over maps and slices

# Notes
- The implementation does not assume that all files (A–H) are present.
- This makes the logic more flexible and suitable for partial or dynamic boards.
- Rank indexing is zero-based internally (rank - 1), matching Go slice indexing.
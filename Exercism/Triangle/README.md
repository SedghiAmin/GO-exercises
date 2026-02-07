# Triangle Type Classifier

A simple Go program that determines the type of a triangle based on its side lengths.

## Overview

This program classifies triangles into four categories:
- **Equilateral**: All three sides are equal
- **Isosceles**: At least two sides are equal
- **Scalene**: All sides are different
- **Not a Triangle**: Invalid triangle based on triangle inequality theorem

## Triangle Validity Rules

For a shape to be a valid triangle:
1. All sides must have positive length (> 0)
2. The sum of any two sides must be greater than the third side:
    - a + b > c
    - b + c > a
    - a + c > b

## Code Structure

### Types and Constants
```go
type Kind int

const (
    NaT Kind = iota  // Not a Triangle (0)
    Equ              // Equilateral (1)
    Iso              // Isosceles (2)
    Sca              // Scalene (3)
)
```

### Main Function
`KindFromSides(a, b, c float64) Kind`
- Takes three side lengths as float64 parameters
- Returns a `Kind` value indicating the triangle type

## Usage

```go
package main

func main() {
    // Test cases
    fmt.Println("Test 1 (2, 2, 2):", KindFromSides(2, 2, 2))  // Equilateral
    fmt.Println("Test 2 (3, 4, 4):", KindFromSides(3, 4, 4))  // Isosceles
    fmt.Println("Test 3 (3, 4, 5):", KindFromSides(3, 4, 5))  // Scalene
    fmt.Println("Test 4 (0, 4, 5):", KindFromSides(0, 4, 5))  // Not a triangle
    fmt.Println("Test 5 (1, 1, 3):", KindFromSides(1, 1, 3))  // Not a triangle
}
```

### Output Interpretation
The program returns integer values representing triangle types:
- `0` = Not a triangle (NaT)
- `1` = Equilateral triangle (Equ)
- `2` = Isosceles triangle (Iso)
- `3` = Scalene triangle (Sca)

## Algorithm

The classification algorithm works as follows:

1. **Validity Check**:
    - Check if any side is ≤ 0
    - Check triangle inequality (a + b > c, etc.)

2. **Type Determination**:
    - Count equal side pairs (z counter starts at 1)
    - If z ≥ 3 → Equilateral (all sides equal)
    - If z ≥ 2 → Isosceles (at least two equal sides)
    - Otherwise → Scalene (all sides different)

## Examples

| Sides (a, b, c) | Result | Reason |
|-----------------|--------|--------|
| (2, 2, 2) | Equilateral | All sides equal |
| (3, 4, 4) | Isosceles | Two sides equal |
| (3, 4, 5) | Scalene | All sides different |
| (0, 4, 5) | Not a triangle | Side length ≤ 0 |
| (1, 1, 3) | Not a triangle | 1 + 1 < 3 |

## Compilation and Execution

```bash
# Save the code to triangle.go
go build triangle.go  # Compile
./triangle           # Run the executable

# Or run directly
go run triangle.go
```

## Dependencies

- Go 1.x or higher
- Standard library only (no external dependencies)

## License

This is example code for educational purposes.
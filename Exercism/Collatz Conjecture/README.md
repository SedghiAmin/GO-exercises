# Collatz Conjecture in Go

A Go implementation of the famous Collatz Conjecture (also known as the 3n+1 problem), one of mathematics' most intriguing unsolved problems.

## What is the Collatz Conjecture?

The Collatz Conjecture is a mathematical conjecture that states: for any positive integer `n`, the following sequence will always eventually reach 1:

1. If `n` is even: divide it by 2
2. If `n` is odd: multiply it by 3 and add 1
3. Repeat the process with the resulting number

The conjecture is that no matter what positive integer you start with, you will always eventually reach 1.

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/collatz-go.git
cd collatz-go
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "errors"
)

func CollatzConjecture(n int) (int, error) {
    if n > 0 {
        step := 0
        for n != 1 {
            if n%2 == 0 {
                n /= 2
            } else {
                n = n*3 + 1
            }
            step++
        }
        return step, nil
    } else {
        return 0, errors.New("the number is invalid")
    }
}
```

### Running the Example

```bash
go run collatz.go
```

Output:
```
Collatz(1) = 0 steps
Collatz(2) = 1 steps
Collatz(3) = 7 steps
Collatz(4) = 2 steps
Collatz(5) = 5 steps
Collatz(6) = 8 steps
Collatz(7) = 16 steps
Collatz(8) = 3 steps
Collatz(9) = 19 steps
Collatz(10) = 6 steps
Collatz(12) = 9 steps
Collatz(19) = 20 steps
Collatz(27) = 111 steps
```

### Testing Specific Numbers

```go
func main() {
    // Test a single number
    steps, err := CollatzConjecture(27)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Collatz(27) takes %d steps to reach 1\n", steps)
    }
    
    // Test with invalid input
    steps, err = CollatzConjecture(0)
    if err != nil {
        fmt.Printf("Error: %v\n", err)  // Output: Error: the number is invalid
    }
}
```

## API Reference

### `CollatzConjecture(n int) (int, error)`

Calculates the number of steps required to reach 1 using the Collatz sequence.

**Parameters:**
- `n int`: Starting positive integer

**Returns:**
- `int`: Number of steps to reach 1
- `error`: Error if input is invalid (n ≤ 0)

**Examples:**
```go
steps, err := CollatzConjecture(6)   // Returns 8, nil
steps, err := CollatzConjecture(0)   // Returns 0, error
```

## Examples

### Example 1: Basic Calculation
```go
steps, err := CollatzConjecture(10)
// steps = 6, err = nil
```

### Example 2: Batch Processing
```go
numbers := []int{1, 2, 3, 4, 5}
for _, n := range numbers {
    steps, _ := CollatzConjecture(n)
    fmt.Printf("Number: %d, Steps: %d\n", n, steps)
}
```

### Example 3: Error Handling
```go
steps, err := CollatzConjecture(-5)
if err != nil {
    fmt.Println("Invalid input:", err)
}
```

## Algorithm Details

The algorithm follows these steps:
1. **Input Validation**: Check if the input is a positive integer
2. **Iterative Process**:
    - While n ≠ 1:
        - If n is even: n = n ÷ 2
        - If n is odd: n = 3n + 1
        - Increment step counter
3. **Return Result**: Return the total number of steps

**Time Complexity**: O(k) where k is the number of steps (varies per input)
**Space Complexity**: O(1)

## Interesting Facts

1. **27 is special**: It takes 111 steps to reach 1, which is unusually high
2. **Largest tested**: The conjecture has been verified for all numbers up to 2⁶⁸ ≈ 2.95×10²⁰
3. **Still unsolved**: Despite its simplicity, the conjecture remains unproven
4. **Known as**: 3n+1 problem, Syracuse problem, Ulam's conjecture, Hasse's algorithm

## Test Cases

The implementation includes tests for:
- ✅ Positive integers
- ✅ Number 1 (0 steps)
- ✅ Even numbers
- ✅ Odd numbers
- ✅ Large sequences (like 27)
- ✅ Invalid inputs (zero and negative numbers)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - see LICENSE file for details

## Acknowledgments

- Lothar Collatz (1910–1990) who first proposed the conjecture
- The mathematical community for decades of research
- All contributors to this implementation

**Note**: This is a simple educational implementation. For research purposes, consider implementing caching mechanisms for better performance with large numbers.
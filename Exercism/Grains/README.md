# Chessboard Grains Calculator

A Go implementation of the classic "wheat and chessboard" problem that demonstrates exponential growth calculations and bitwise operations.

## 📖 The Legend

This program solves the famous mathematical problem: If you place one grain of wheat on the first square of a chessboard, two on the second, four on the third, and continue doubling each square, how many grains would you have?

## 🚀 Features

- **Exponential Calculation**: Computes grains on any chessboard square using bitwise shifting
- **Total Sum Calculation**: Calculates the total grains on all 64 squares using mathematical formula
- **Input Validation**: Validates square numbers (1-64) and returns appropriate errors
- **Efficient Implementation**: Uses bitwise operations for optimal performance

## 📦 Installation

```bash
# Clone the repository
git clone <repository-url>
cd chessboard-grains

# Run the program
go run main.go
```

## 🛠️ API Reference

### `Square(n int) (uint64, error)`

Calculates the number of grains on a specific chessboard square.

**Parameters:**
- `n` (int): The square number (1-64)

**Returns:**
- `uint64`: Number of grains on square `n`
- `error`: Error if `n` is outside valid range (1-64)

**Formula:** `grains = 2^(n-1)`

**Example:**
```go
grains, err := Square(1)  // Returns: 1, nil
grains, err := Square(3)  // Returns: 4, nil  
grains, err := Square(0)  // Returns: 0, error
```

### `Total() uint64`

Calculates the total number of grains on the entire chessboard.

**Returns:**
- `uint64`: Total grains on all 64 squares

**Formula:** `total = 2^64 - 1`

**Example:**
```go
total := Total()  // Returns: 18,446,744,073,709,551,615
```

## 📊 Mathematical Background

### Individual Squares
The number of grains doubles on each successive square:
- Square 1: 1 = 2⁰
- Square 2: 2 = 2¹
- Square 3: 4 = 2²
- ...
- Square n: 2ⁿ⁻¹

### Total Grains
The sum of a geometric series:
```
Total = 2⁰ + 2¹ + 2² + ... + 2⁶³ = 2⁶⁴ - 1
```

### Key Numbers
- Square 64: 9,223,372,036,854,775,808 grains
- Total: 18,446,744,073,709,551,615 grains
- For comparison: World wheat production is ~770 million tons annually

## 💡 Implementation Details

### Bitwise Optimization
The program uses bitwise left shift (`<<`) for efficient power-of-two calculations:
```go
1 << (n-1)  // Equivalent to 2^(n-1)
(1 << 64) - 1  // Equivalent to 2^64 - 1
```

### Type Considerations
- Uses `uint64` for maximum range (up to 2⁶⁴ - 1)
- Square 64 calculation (2⁶³) fits within `uint64` limits
- Total calculation (2⁶⁴ - 1) is the maximum `uint64` value

## 🧪 Usage Examples

```go
package main

import (
    "fmt"
)

func main() {
    // Calculate total grains
    fmt.Printf("Total grains on chessboard: %d\n", Total())
    
    // Calculate grains on specific squares
    squares := []int{1, 2, 3, 4, 5, 10, 20, 32, 64}
    for _, square := range squares {
        grains, err := Square(square)
        if err != nil {
            fmt.Printf("Error for square %d: %v\n", square, err)
        } else {
            fmt.Printf("Square %d: %d grains\n", square, grains)
        }
    }
    
    // Handle error case
    grains, err := Square(65)
    if err != nil {
        fmt.Println("Expected error:", err)
    }
}
```

**Sample Output:**
```
Total grains on chessboard: 18446744073709551615
Square 1: 1 grains
Square 2: 2 grains
Square 3: 4 grains
Square 4: 8 grains
Square 5: 16 grains
Square 10: 512 grains
Square 20: 524288 grains
Square 32: 2147483648 grains
Square 64: 9223372036854775808 grains
Expected error: square number must be between 1 and 64
```

## 🔧 Running Tests

To test the implementation:

```bash
# Create a test file
touch main_test.go

# Add test cases (see example below)
# Run tests
go test -v
```

### Example Test File:
```go
package main

import (
    "testing"
)

func TestSquare(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        want     uint64
        wantErr  bool
    }{
        {"square 1", 1, 1, false},
        {"square 2", 2, 2, false},
        {"square 3", 3, 4, false},
        {"square 64", 64, 9223372036854775808, false},
        {"square 0", 0, 0, true},
        {"square 65", 65, 0, true},
        {"square negative", -1, 0, true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Square(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("Square() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Square() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestTotal(t *testing.T) {
    want := uint64(18446744073709551615)
    if got := Total(); got != want {
        t.Errorf("Total() = %v, want %v", got, want)
    }
}
```

## 🚀 Performance

- **Time Complexity**: O(1) for both functions
- **Space Complexity**: O(1)
- Uses bitwise operations instead of slower `math.Pow()` calls
- No loops or recursion needed

## 📚 Learning Points

This implementation demonstrates:
- Bitwise operations in Go
- Working with large integers (`uint64`)
- Geometric series calculations
- Input validation and error handling
- Clean function design and separation of concerns

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is open source and available under the [MIT License](LICENSE).

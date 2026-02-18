# Armstrong Number Checker in Go

This program provides a function to determine whether a given integer is an **Armstrong number** (also known as a **narcissistic number** or **pluperfect number**).

## What is an Armstrong Number?

An Armstrong number is a number that equals the sum of its own digits each raised to the power of the number of digits.

### Examples:

- **153** is an Armstrong number because: 1³ + 5³ + 3³ = 1 + 125 + 27 = 153
- **370** is an Armstrong number because: 3³ + 7³ + 0³ = 27 + 343 + 0 = 370
- **371** is an Armstrong number because: 3³ + 7³ + 1³ = 27 + 343 + 1 = 371
- **407** is an Armstrong number because: 4³ + 7³ + 0³ = 64 + 343 + 0 = 407
- **9474** is an Armstrong number because: 9⁴ + 4⁴ + 7⁴ + 4⁴ = 6561 + 256 + 2401 + 256 = 9474

## Code Implementation

```go
package main

import (
    "fmt"
    "math"
)

// IsNumber checks if a given integer is an Armstrong number
func IsNumber(n int) bool {
    // 0 is an Armstrong number (0¹ = 0)
    if n == 0 {
        return true
    }
    
    // Handle negative numbers (convert to positive for digit extraction)
    // Note: Negative numbers are not typically considered Armstrong numbers,
    // but the function will process them by their absolute value
    if n < 0 {
        n = -n
    }
    
    // Store the original number for comparison
    number := n
    
    // Extract digits
    digits := make([]int, 0)
    for n > 0 {
        digits = append(digits, n%10) // Get the last digit
        n /= 10                        // Remove the last digit
    }
    
    // Calculate the sum of digits raised to the power of digit count
    sum := 0.0
    power := len(digits)
    
    for i := range digits {
        sum += math.Pow(float64(digits[i]), float64(power))
    }
    
    // Check if the sum equals the original number
    return sum == float64(number)
}

func main() {
    // Test cases
    fmt.Println(IsNumber(10))   // false
    fmt.Println(IsNumber(9))    // true (9¹ = 9)
    fmt.Println(IsNumber(153))  // true (1³ + 5³ + 3³ = 153)
}
```

## How It Works

### Step-by-Step Process

1. **Special Case Handling**: The number 0 is immediately identified as an Armstrong number.

2. **Negative Number Handling**: If the input is negative, the function converts it to positive for digit extraction. (Note: Negative numbers are not typically considered Armstrong numbers in mathematical literature.)

3. **Digit Extraction**: The function repeatedly divides the number by 10 to extract each digit:
    - `n % 10` gives the last digit
    - `n /= 10` removes the last digit
    - This continues until all digits are extracted

4. **Power Calculation**: For each digit, the function calculates:
    - `digit ^ (number of digits)`
    - Uses `math.Pow()` from the standard library

5. **Comparison**: The sum of these powers is compared to the original number.

## Algorithm Analysis

### Time Complexity
- **O(d)**: where d is the number of digits in the number
- For an n-digit number, the algorithm performs d operations for digit extraction and d operations for power calculation

### Space Complexity
- **O(d)**: stores all digits in a slice for processing

## Usage Examples

```go
func main() {
    testNumbers := []int{
        0, 1, 2, 3, 4, 5, 6, 7, 8, 9,      // All single-digit numbers are Armstrong numbers
        10, 11, 12,                         // Two-digit numbers (few are Armstrong)
        153, 370, 371, 407,                  // Three-digit Armstrong numbers
        9474,                                // Four-digit Armstrong number
        54748,                               // Five-digit Armstrong number
        548834,                              // Six-digit Armstrong number
        -153,                                // Negative number (processed as positive)
    }
    
    for _, num := range testNumbers {
        if IsNumber(num) {
            fmt.Printf("%d is an Armstrong number\n", num)
        } else {
            fmt.Printf("%d is NOT an Armstrong number\n", num)
        }
    }
}
```

## Complete Test Suite

```go
package main

import "testing"

func TestIsNumber(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
    }{
        {"Zero", 0, true},
        {"Single digit 1", 1, true},
        {"Single digit 5", 5, true},
        {"Single digit 9", 9, true},
        {"Two-digit 10", 10, false},
        {"Two-digit 11", 11, false},
        {"Two-digit 12", 12, false},
        {"Three-digit 153", 153, true},
        {"Three-digit 154", 154, false},
        {"Three-digit 370", 370, true},
        {"Three-digit 371", 371, true},
        {"Three-digit 407", 407, true},
        {"Four-digit 9474", 9474, true},
        {"Four-digit 9475", 9475, false},
        {"Five-digit 54748", 54748, true},
        {"Six-digit 548834", 548834, true},
        {"Negative number", -153, true}, // Processed as positive 153
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := IsNumber(tt.input)
            if result != tt.expected {
                t.Errorf("IsNumber(%d) = %v; expected %v", 
                         tt.input, result, tt.expected)
            }
        })
    }
}
```

## Known Armstrong Numbers

Here are some well-known Armstrong numbers:

| Digits | Armstrong Numbers |
|--------|-------------------|
| 1 | 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 |
| 3 | 153, 370, 371, 407 |
| 4 | 1634, 8208, 9474 |
| 5 | 54748, 92727, 93084 |
| 6 | 548834 |
| 7 | 1741725, 4210818, 9800817, 9926315 |

## Performance Considerations

### Optimization without `math.Pow`

For better performance, especially with many calculations, you can avoid `math.Pow`:

```go
func IsNumberOptimized(n int) bool {
    if n < 0 {
        n = -n
    }
    
    original := n
    digits := make([]int, 0)
    
    for n > 0 {
        digits = append(digits, n%10)
        n /= 10
    }
    
    power := len(digits)
    sum := 0
    
    for _, digit := range digits {
        // Calculate digit^power without math.Pow
        product := 1
        for i := 0; i < power; i++ {
            product *= digit
        }
        sum += product
    }
    
    return sum == original
}
```

## Limitations

1. **Integer Overflow**: For very large numbers, the sum might exceed the maximum integer value
2. **Negative Numbers**: The mathematical definition typically excludes negative numbers
3. **Performance**: Using `math.Pow` for integer powers is slightly inefficient

## Mathematical Background

Armstrong numbers are named after Michael F. Armstrong, who used them as a programming exercise in the 1960s. They're also called:
- **Narcissistic numbers**
- **Pluperfect numbers**
- **Plus perfect numbers**

The property is that the number is "in love with itself" - hence the name "narcissistic".

## License

This code is provided as an educational example and is free to use, modify, and distribute.
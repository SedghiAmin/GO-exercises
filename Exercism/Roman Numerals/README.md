# Roman Numeral Converter in Go

A Go program that converts Arabic numerals (1-3999) to traditional Roman numerals.

## 📖 Overview

This Go implementation converts integers to Roman numerals following the traditional rules and conventions of Roman numeral notation. It handles all standard cases including subtraction notation for numbers like 4 (IV), 9 (IX), 40 (XL), etc.

## 🚀 Features

- **Complete Range Support**: Converts numbers from 1 to 3999 (traditional Roman numeral limits)
- **Proper Subtraction Notation**: Correctly implements IV, IX, XL, XC, CD, CM
- **Error Handling**: Returns appropriate errors for invalid input ranges
- **Efficient Algorithm**: Uses subtraction method with pre-defined value-symbol mapping
- **Clean Interface**: Simple function API with clear input/output

## 📦 Installation

```bash
# Clone or download the package
git clone <repository-url>
cd roman-numeral-converter
```

```go
package main

import (
    "fmt"
)

func main() {
    // Convert a number to Roman numeral
    result, err := ToRomanNumeral(1996)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("1996 in Roman numerals:", result) // Output: MCMXCVI
    }
}
```

### Function Signature

```go
func ToRomanNumeral(input int) (string, error)
```

**Parameters:**
- `input int`: Arabic numeral to convert (must be 1-3999)

**Returns:**
- `string`: Roman numeral representation
- `error`: Error if input is outside valid range (1-3999)

## 📊 Conversion Examples

| Arabic | Roman Numeral | Notes |
|--------|---------------|-------|
| 1 | I | Basic unit |
| 4 | IV | Subtraction notation |
| 9 | IX | Subtraction notation |
| 10 | X | Basic symbol |
| 14 | XIV | Combination |
| 37 | XXXVII | Your example case |
| 40 | XL | Subtraction notation |
| 49 | XLIX | Complex subtraction |
| 90 | XC | Subtraction notation |
| 99 | XCIX | Double subtraction |
| 100 | C | Basic symbol |
| 400 | CD | Subtraction notation |
| 444 | CDXLIV | Multiple subtractions |
| 900 | CM | Subtraction notation |
| 999 | CMXCIX | Complex case |
| 1000 | M | Basic symbol |
| 1996 | MCMXCVI | Complex modern date |
| 2023 | MMXXIII | Recent year |
| 3999 | MMMCMXCIX | Maximum traditional value |

## 🔍 Algorithm Details

The converter uses a **greedy algorithm** that works as follows:

1. **Pre-defined Mapping**: Uses two parallel arrays:
    - `values`: [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
    - `symbols`: ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]

2. **Iterative Subtraction**:
    - Start with the largest value (1000 = "M")
    - While the input is greater than or equal to the current value:
        - Append the corresponding symbol to output
        - Subtract the value from input
    - Move to the next smaller value
    - Repeat until input reaches 0

3. **Example: Converting 37**:
   ```
   37 >= 10? Yes → Add "X" → Remainder: 27
   27 >= 10? Yes → Add "X" → Remainder: 17
   17 >= 10? Yes → Add "X" → Remainder: 7
   7 >= 5? Yes → Add "V" → Remainder: 2
   2 >= 1? Yes → Add "I" → Remainder: 1
   1 >= 1? Yes → Add "I" → Remainder: 0
   Result: XXXVII
   ```

## 🧪 Testing

### Built-in Example

```go
func main() {
    fmt.Println(ToRomanNumeral(37)) // Output: XXXVII <nil>
}
```

### Extended Test Suite

```go
package main

import (
    "fmt"
    "strings"
)

func testConverter() {
    testCases := []struct {
        input    int
        expected string
    }{
        {1, "I"},
        {4, "IV"},
        {9, "IX"},
        {10, "X"},
        {14, "XIV"},
        {37, "XXXVII"},
        {49, "XLIX"},
        {99, "XCIX"},
        {444, "CDXLIV"},
        {999, "CMXCIX"},
        {1996, "MCMXCVI"},
        {2023, "MMXXIII"},
        {3999, "MMMCMXCIX"},
    }

    fmt.Println("Roman Numeral Converter Test Results:")
    fmt.Println(strings.Repeat("=", 50))

    for _, tc := range testCases {
        result, err := ToRomanNumeral(tc.input)
        status := "✅"
        if err != nil {
            status = "❌"
            result = err.Error()
        } else if result != tc.expected {
            status = "❌"
        }
        fmt.Printf("%s %4d → %-15s (expected: %s)\n", 
            status, tc.input, result, tc.expected)
    }
}

func main() {
    testConverter()
}
```

## ⚠️ Error Cases

The function returns an error for:
- Numbers less than 1
- Numbers greater than 3999 (traditional Roman numeral limit)

```go
// Error examples
result, err := ToRomanNumeral(0)    // Error: "the number must be between 1 to 3999"
result, err := ToRomanNumeral(4000) // Error: "the number must be between 1 to 3999"
result, err := ToRomanNumeral(-5)   // Error: "the number must be between 1 to 3999"
```

## 🏗️ Code Structure

```go
// Core conversion function
func ToRomanNumeral(input int) (string, error) {
    // 1. Input validation (1-3999)
    // 2. Initialize parallel arrays for values and symbols
    // 3. Iterate through values from largest to smallest
    // 4. Use greedy algorithm to build Roman numeral
    // 5. Return concatenated result
}
```

## 📝 Implementation Notes

1. **Traditional Limits**: Roman numerals traditionally only go up to 3,999 (MMMCMXCIX)
2. **No Zero**: The Roman numeral system has no representation for zero
3. **Subtraction Rules**: Only specific subtractions are allowed (IV, IX, XL, XC, CD, CM)
4. **Order Matters**: Symbols must be in descending order of value
5. **No More Than Three**: No symbol appears more than three times consecutively

## 🔧 Performance

- **Time Complexity**: O(1) - Constant time as it iterates through fixed 13 values
- **Space Complexity**: O(1) - Fixed-size arrays and string builder
- **Efficiency**: Uses `strings.Join` for efficient string concatenation

## 🤝 Contributing

Contributions are welcome! Please ensure:

1. All conversions follow traditional Roman numeral rules
2. Error handling is maintained for edge cases
3. Test cases are updated for new functionality
4. Code follows Go conventions and is properly documented

## 📄 License

This project is open source and available for educational and practical use.

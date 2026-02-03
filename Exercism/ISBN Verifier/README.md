# ISBN-10 Validator

A simple Go implementation for validating ISBN-10 (International Standard Book Number) codes.

## Overview

This Go package provides a function to validate ISBN-10 codes according to the official ISBN-10 verification formula. The validator handles both hyphenated and non-hyphenated formats, and properly handles the special check character 'X' (representing the value 10).

## Features

- ✅ Validates ISBN-10 codes with and without hyphens
- ✅ Handles the special check character 'X' (case-insensitive)
- ✅ Implements the official ISBN-10 verification algorithm
- ✅ Returns boolean result for easy integration
- ✅ Includes comprehensive error checking

## ISBN-10 Format

The ISBN-10 format consists of:
- 9 digits (0-9)
- 1 check character (either a digit 0-9 or 'X' representing 10)
- Optional hyphens for formatting

### Verification Formula

The ISBN-10 is validated using this formula:

```
(d₁ × 10 + d₂ × 9 + d₃ × 8 + d₄ × 7 + d₅ × 6 + d₆ × 5 + d₇ × 4 + d₈ × 3 + d₉ × 2 + d₁₀ × 1) mod 11 == 0
```

Where `d₁` to `d₉` are digits and `d₁₀` is the check character (digit or 'X').

## Installation

```bash
# Clone the repository
git clone <repository-url>
cd isbn-validator
```

## Usage

### Basic Usage

```go
package main

import "fmt"

func main() {
    // Test various ISBN-10 codes
    testISBNs := []string{
        "3-598-21508-8",  // Valid
        "3-598-21507-X",  // Valid with 'X'
        "359821507X",     // Valid without hyphens
        "3-598-21508-9",  // Invalid
    }
    
    for _, isbn := range testISBNs {
        valid := IsValidISBN(isbn)
        fmt.Printf("ISBN: %-20s Valid: %v\n", isbn, valid)
    }
}
```

### Function Signature

```go
func IsValidISBN(isbn string) bool
```

**Parameters:**
- `isbn`: The ISBN-10 string to validate (can contain hyphens)

**Returns:**
- `true` if the ISBN is valid
- `false` if the ISBN is invalid

## Examples

### Valid ISBNs

```go
IsValidISBN("3-598-21508-8")   // returns true
IsValidISBN("3-598-21507-X")   // returns true
IsValidISBN("3-598-21507-x")   // returns true (case-insensitive)
IsValidISBN("3598215088")      // returns true (no hyphens)
IsValidISBN("0-7475-3269-9")   // returns true
```

### Invalid ISBNs

```go
IsValidISBN("3-598-21508-9")   // returns false (wrong check digit)
IsValidISBN("3-598-21507-A")   // returns false (invalid character)
IsValidISBN("3-598-21508")     // returns false (too short)
IsValidISBN("3-598-21508-81")  // returns false (too long)
IsValidISBN("")                // returns false (empty string)
```

## Algorithm Details

The validator performs these steps:

1. **Preprocessing**: Removes all hyphens from the input string
2. **Length Validation**: Checks if the cleaned string has exactly 10 characters
3. **Character Validation**:
    - For positions 1-9: Verifies each character is a digit (0-9)
    - For position 10: Allows digits 0-9 or 'X'/'x'
4. **Weight Calculation**:
    - Converts characters to their numeric values
    - Multiplies each digit by its decreasing weight (10, 9, 8, ..., 1)
5. **Modulo Check**: Calculates the sum modulo 11 and checks if it equals 0

## Testing

The package includes built-in test cases in the `main()` function. To run the tests:

```bash
go run main.go
```

Expected output:
```
Check ISBN-10:
==============
'3-598-21508-8': true
'3-598-21507-X': true
'359821507X': true
'3-598-21508-9': false
```

## Edge Cases Handled

- ✅ Mixed case 'X' or 'x'
- ✅ Multiple or no hyphens
- ✅ Invalid characters in any position
- ✅ Incorrect length strings
- ✅ Empty strings

## Dependencies

- Standard Go library only
- No external dependencies required

## Performance

The algorithm runs in O(n) time complexity where n is the length of the input string, making it efficient for batch processing of ISBN codes.

## Limitations

- Only validates ISBN-10 format (not ISBN-13)
- Does not verify the publisher or country codes
- No ISBN lookup functionality (only format validation)

## Contributing

Feel free to submit issues and enhancement requests. When contributing code, please ensure all existing tests pass and add new tests for new functionality.

## License

This project is open source and available under the MIT License.

## References

- [International ISBN Agency](https://www.isbn-international.org/)
- [ISBN-10 Wikipedia](https://en.wikipedia.org/wiki/International_Standard_Book_Number#ISBN-10_check_digits)
- [ISBN Users' Manual](https://www.isbn-international.org/sites/default/files/ISBN%20Manual%202012%20-corr.pdf)
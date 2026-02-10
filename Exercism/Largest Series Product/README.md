# Largest Series Product - Go Implementation

## Overview
This Go program calculates the largest product of a contiguous series of digits in a given string. It's a solution to the "Largest Series Product" problem, similar to the one found on Exercism.

## Problem Statement
Given a string of digits, find the largest product of a contiguous substring of digits of length `n`.

### Requirements:
- The input string must contain only digits (0-9)
- `span` (n) must be non-negative
- `span` must be less than or equal to the length of the input string
- Return the maximum product as an `int64`
- Handle error cases appropriately

## Code Explanation

### Function Signature
```go
func LargestSeriesProduct(digits string, span int) (int64, error)
```

### Parameters
- `digits`: A string containing digits (should only contain characters '0'-'9')
- `span`: An integer representing the length of the contiguous series to consider

### Return Values
- Returns the maximum product as `int64`
- Returns an error if:
    - `span` is negative
    - `span` is greater than the length of `digits`
    - `digits` contains non-digit characters

### Algorithm
1. **Input Validation**:
    - Check if `span` is negative
    - Check if `span` is greater than the length of `digits`
    - Verify all characters in `digits` are digits (0-9)

2. **Product Calculation**:
    - Convert the string to a slice of runes
    - For each possible starting position `i` (from 0 to `len(digits)-span`)
    - Calculate the product of `span` consecutive digits
    - Store each product in a results array

3. **Find Maximum**:
    - Use Go's `slices.Max()` to find the maximum product
    - Return the maximum value

## Key Features

### Error Handling
The function returns descriptive errors for invalid inputs:
- `"span must not be negative"` for negative span values
- `"span must be smaller than string length"` when span exceeds string length
- `"digits input must only contain digits"` for non-digit characters

### Edge Cases Handled
- Empty string with non-zero span
- String shorter than requested span
- Zero span (note: current implementation doesn't handle span=0 specifically)
- Strings containing all zeros

## Example Usage

```go
package main

import "fmt"

func main() {
    // Example 1: Basic case
    result, err := LargestSeriesProduct("12345", 2)
    // Product of "12" = 1*2 = 2
    // Product of "23" = 2*3 = 6
    // Product of "34" = 3*4 = 12
    // Product of "45" = 4*5 = 20
    // Returns: 20, nil
    
    // Example 2: Error case
    result, err := LargestSeriesProduct("1234a5", 2)
    // Returns: 0, error("digits input must only contain digits")
    
    // Example 3: Large span
    result, err := LargestSeriesProduct("12345", 5)
    // Product of "12345" = 1*2*3*4*5 = 120
    // Returns: 120, nil
}
```

## Test Cases

The `main()` function includes several test cases:

1. **Valid cases**:
    - `"12345", 5` → Error (span equals string length, but code expects span < length)
    - `"99099", 3` → 0 (contains 0 in series)
    - `"0000", 2` → 0
    - Large number with span 6 → 23520
    - `"0123456789", 5` → 15120 (5*6*7*8*9)
    - `"1027839564", 3` → 270

2. **Error cases**:
    - Negative span → error
    - Non-digit characters → error
    - Empty string → error
    - Span larger than string length → error

## Implementation Notes

### Important Considerations
1. **Array Size**: The results array is sized correctly as `len(digits)-span+1` to avoid unnecessary zero values
2. **Character Conversion**: Digits are converted from runes to integers using `int64(digitsRunes[j] - '0')`
3. **Case Sensitivity**: All digits are treated equally (0-9)
4. **Performance**: The algorithm has O(n*m) time complexity where n is string length and m is span

### Potential Improvements
1. Add handling for `span = 0` (typically returns 1)
2. Consider using a sliding window approach for better performance with large spans
3. Add more comprehensive test coverage

## Dependencies
- Go 1.21+ (for `slices` package)
- Standard library packages:
    - `fmt`
    - `slices`
    - `unicode`

## Output
Running the provided `main()` function will output:
```
0 span must be smaller than string length
0 span must not be negative
0 digits input must only contain digits
0 span must be smaller than string length
0 span must be smaller than string length
0
0
23520 <nil>
15120 <nil>
270 <nil>
```


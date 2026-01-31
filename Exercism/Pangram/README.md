# Pangram Checker in Go

A simple and efficient Go implementation to check if a given string is a pangram.

## What is a Pangram?

A pangram is a sentence that contains every letter of the alphabet at least once. The most famous example is:
> "The quick brown fox jumps over the lazy dog."

## Features

- **Case-insensitive**: Automatically handles uppercase and lowercase letters
- **Efficient**: Uses a map for O(n) time complexity
- **Simple API**: Single function with clear interface
- **Unicode-aware**: Properly handles runes (Go's character type)
- **Memory efficient**: Only tracks seen letters

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/go-pangram.git
cd go-pangram

# Or simply copy the function into your project
```

## Usage

```go
package main

import (
    "fmt"
    "strings"
)

func IsPangram(input string) bool {
    alphabet := make(map[rune]bool)

    for _, char := range strings.ToUpper(input) {
        if char >= 'A' && char <= 'Z' {
            alphabet[char] = true
        }
    }

    return len(alphabet) == 26
}

func main() {
    // Test cases
    tests := []struct {
        sentence string
        expected bool
    }{
        {"The quick brown fox jumps over the lazy dog", true},
        {"Pack my box with five dozen liquor jugs", true},
        {"How vexingly quick daft zebras jump!", true},
        {"a quick movement of the enemy will jeopardize five gunboats", false},
        {"Hello, World!", false},
        {"", false},
        {"abcdefghijklmnopqrstuvwxyz", true},
        {"ABCDEFGHIJKLMNOPQRSTUVWXYZ", true},
    }

    for _, test := range tests {
        result := IsPangram(test.sentence)
        status := "✓"
        if result != test.expected {
            status = "✗"
        }
        fmt.Printf("%s %q -> %v (expected: %v)\n", status, test.sentence, result, test.expected)
    }
}
```

## How It Works

1. **Convert to uppercase**: `strings.ToUpper(input)` makes the check case-insensitive
2. **Iterate through characters**: Loop through each rune in the string
3. **Filter letters**: Only process characters between 'A' and 'Z'
4. **Track seen letters**: Use a map to mark which letters have been seen
5. **Check count**: If 26 unique letters have been seen, it's a pangram

## Algorithm Complexity

- **Time Complexity**: O(n) where n is the length of the input string
- **Space Complexity**: O(1) (at most 26 entries in the map)
- **Best Case**: Early return possible with optimization
- **Worst Case**: Need to process entire string

## Alternative Implementations

### Using Bit Manipulation (More Efficient)

```go
func IsPangramBitwise(input string) bool {
    var bitmask uint32 // 32 bits are enough for 26 letters
    
    for _, char := range strings.ToUpper(input) {
        if char >= 'A' && char <= 'Z' {
            index := char - 'A'
            bitmask |= 1 << index
        }
    }
    
    // 0x3FFFFFF = binary 11111111111111111111111111 (26 ones)
    return bitmask == 0x3FFFFFF
}
```

### Using Array Instead of Map

```go
func IsPangramArray(input string) bool {
    var seen [26]bool
    
    for _, char := range strings.ToUpper(input) {
        if char >= 'A' && char <= 'Z' {
            seen[char-'A'] = true
        }
    }
    
    for _, s := range seen {
        if !s {
            return false
        }
    }
    return true
}
```

## Performance Comparison

| Method | Time | Space | Readability |
|--------|------|-------|-------------|
| Map (Current) | O(n) | O(1) | ⭐⭐⭐⭐⭐ |
| Bitwise | O(n) | O(1) | ⭐⭐⭐ |
| Array | O(n) | O(1) | ⭐⭐⭐⭐ |

The current implementation offers the best balance of readability and performance for most use cases.

## Edge Cases Handled

- Empty strings return `false`
- Strings with only punctuation return `false`
- Mixed case works correctly
- Non-English characters are ignored
- Very long strings work efficiently

## Testing

Run the examples:
```bash
go run pangram.go
```

Expected output:
```
false
true
```

## Dependencies

- Go 1.13+ (but compatible with older versions)
- Standard library only

## Contributing

Feel free to:
1. Report issues
2. Suggest improvements
3. Submit pull requests
4. Add more test cases

## License

MIT License 
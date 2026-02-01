# String Reversal in Go

A simple Go implementation of string reversal that correctly handles Unicode characters, including multi-byte runes like emojis and international characters.

## Features

- ✅ **Unicode Support**: Correctly reverses strings containing multi-byte characters (e.g., "Hello, 世界" → "界世 ,olleH")
- ✅ **Efficient**: Uses rune slices for O(n) time complexity
- ✅ **Simple API**: Single function with clean interface
- ✅ **No External Dependencies**: Pure Go implementation

```go
// reverse.go
func Reverse(input string) string {
    runes := []rune(input)
    r := make([]rune, len(runes))
    step := len(runes) - 1
    for _, char := range runes {
        r[step] = char
        step--
    }
    return string(r)
}
```

## Usage

```go
package main

import (
    "fmt"
)

func main() {
    // Basic ASCII string
    fmt.Println(Reverse("Hello"))           // "olleH"
    
    // String with spaces and punctuation
    fmt.Println(Reverse("Hello, World!"))   // "!dlroW ,olleH"
    
    // Unicode string with multi-byte characters
    fmt.Println(Reverse("Hello, 世界"))      // "界世 ,olleH"
    
    // String with emojis
    fmt.Println(Reverse("Go is awesome! 🚀")) // "🚀 !emosewa si oG"
    
    // Arabic text (right-to-left script)
    fmt.Println(Reverse("مرحبا"))           // "ابحم"
    
    // Mixed script
    fmt.Println(Reverse("Hello 你好 नमस्ते")) // "ेत्समन 好你 olleH"
}
```

## How It Works

The function works in three steps:

1. **Convert to Runes**: `[]rune(input)` converts the string to a slice of Unicode code points
2. **Create Reversed Slice**: Allocate a new slice with the same length
3. **Reverse Order**: Copy runes from the end to the beginning

### Comparison with Naive Approaches

| Approach | "Hello" | "Hello, 世界" | Complexity |
|----------|---------|---------------|------------|
| **This Implementation** | "olleH" | "界世 ,olleH" | O(n) |
| Byte-by-Byte Reversal | "olleH" | Incorrect (乱码) | O(n) |
| Built-in Go Methods | Not available | Not available | N/A |

## Testing

The implementation has been tested with various edge cases:

```go
// Test cases
testCases := []struct {
    input    string
    expected string
}{
    {"", ""},                           // Empty string
    {"a", "a"},                         // Single character
    {"ab", "ba"},                       // Two characters
    {"racecar", "racecar"},             // Palindrome
    {"Hello, 世界", "界世 ,olleH"},      // Unicode
    {"👋🌍", "🌍👋"},                    // Emojis
    {"नमस्ते", "ेत्समन"},               // Devanagari script
}
```

## Performance

- **Time Complexity**: O(n) where n is the number of runes in the string
- **Space Complexity**: O(n) for the rune slice allocation
- **Memory**: Creates one additional rune slice of the same size as input

For extremely large strings, consider using in-place reversal:

```go
func ReverseInPlace(input string) string {
    runes := []rune(input)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

## License

MIT License. See LICENSE file for details.

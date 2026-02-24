# Ten Green Bottles - Go Edition

A clean and efficient Go implementation of the classic children's repetitive song "Ten Green Bottles". This package generates the complete lyrics with proper formatting, number-to-word conversion, and grammatical accuracy.

## 📋 Features

- Generate any number of verses from any starting point
- Automatic number-to-word conversion with proper capitalization
- Intelligent grammar handling (bottle/bottles, "no bottles" for zero)
- Clean output formatting with empty lines between verses
- Memory-efficient implementation with pre-allocated slices
- Simple and intuitive API

## 💻 Usage

### Basic Example

```go
package main

import "fmt"

func main() {
    // Generate all 10 verses
    lyrics := Recite(10, 10)
    for _, line := range lyrics {
        fmt.Println(line)
    }
}
```

### Function Reference

#### `Recite(startBottles, takeDown int) []string`

The main function that generates the song lyrics.

| Parameter | Description | Example |
|-----------|-------------|---------|
| `startBottles` | Starting number of bottles (1-10) | `10` for "Ten green bottles" |
| `takeDown` | Number of verses to generate | `3` for three verses |

**Return Value:** A slice of strings, each containing one line of the song.

#### Examples:

```go
// Generate the first 3 verses (10, 9, 8 bottles)
verses := Recite(10, 3)

// Generate verses from 5 down to 3
verses := Recite(5, 3)

// Generate just the last verse (1 bottle)
verses := Recite(1, 1)

// Invalid inputs return empty slice
verses := Recite(0, 5)  // returns []
```

#### `NumToWord(num int) string`

Converts a number (1-10) to its English word representation with proper pluralization and capitalization.

| Input | Output |
|-------|--------|
| 1 | "One green bottle" |
| 2 | "Two green bottles" |
| 3 | "Three green bottles" |
| ... | ... |
| 10 | "Ten green bottles" |

## 📝 Output Format

Each verse follows this pattern:

```
[Number] green bottle(s) hanging on the wall,
[Number] green bottle(s) hanging on the wall,
And if one green bottle should accidentally fall,
There'll be [next number] green bottle(s) hanging on the wall.

[Next verse...]
```

### Example Output (Recite(3, 2)):

```
Three green bottles hanging on the wall,
Three green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be two green bottles hanging on the wall.

Two green bottles hanging on the wall,
Two green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be one green bottle hanging on the wall.
```

### Special Case - Last Verse (1 bottle):

```
One green bottle hanging on the wall,
One green bottle hanging on the wall,
And if one green bottle should accidentally fall,
There'll be no green bottles hanging on the wall.
```

## 🔧 Implementation Details

### Key Design Decisions

1. **Direct Capitalization**: Numbers are directly capitalized in `NumToWord()` for simplicity
2. **Lowercase Conversion**: Uses `strings.ToLower()` for subsequent appearances
3. **Pre-allocated Slices**: Memory-efficient with capacity calculation
4. **Input Validation**: Comprehensive checks for invalid parameters
5. **Clean Formatting**: Proper comma placement and empty line separation

### Code Structure

```
ten-green-bottles/
├── main.go          # Core implementation
├── main_test.go     # Unit tests
└── README.md        # This file
```

## 🧪 Testing

Run the test suite:

```bash
go test -v
```

Example test cases:
- Single verse generation
- Multiple verses generation
- Edge cases (start=1, takeDown=1)
- Invalid input handling
- Grammar verification

## ⚡ Performance

- **Time Complexity**: O(n) where n is the number of verses
- **Space Complexity**: O(n) for storing the lyrics
- **Memory Efficient**: Pre-allocates slice capacity to avoid reallocation

## 🤝 Contributing

Contributions are welcome! Feel free to:

1. Fork the repository
2. Create a feature branch
3. Submit a Pull Request

Please ensure:
- Code is properly formatted (`go fmt`)
- Tests pass (`go test`)
- New features include tests

## 📄 License

MIT License - feel free to use this code in your projects!

## 🎵 Complete Lyrics

For reference, here's the complete song structure:

```
Ten green bottles hanging on the wall,
Ten green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be nine green bottles hanging on the wall.

Nine green bottles hanging on the wall,
... (continues down to 1) ...

One green bottle hanging on the wall,
One green bottle hanging on the wall,
And if one green bottle should accidentally fall,
There'll be no green bottles hanging on the wall.
```

## ✨ Why This Implementation?

- **Clean Code**: Simple, readable, and maintainable
- **Go Idiomatic**: Follows Go best practices and conventions
- **Efficient**: Optimized for both memory and performance
- **Flexible**: Generate any subset of verses you need
- **Educational**: Great example of string manipulation in Go
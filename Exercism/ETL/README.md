# Lexiconia Score Transformer

A Go utility for transforming letter score data structures from grouped format to individual mapping format, designed for the Lexiconia multiplayer word game.

## Overview

This tool converts the scoring system data structure from a grouped format (letters grouped by score) to an individual format (each letter mapped directly to its score). This transformation supports the internationalization efforts of the Lexiconia game by making it easier to add support for new languages.

## Problem Context

Lexiconia is an online multiplayer game where players rearrange letters to create words. Different letters have different point values based on their frequency in a language. The game is expanding from English to other languages, each requiring different point values for letters.

### Original Data Structure (Grouped by Score)
```go
1 point:  ["A", "E", "I", "O", "U", "L", "N", "R", "S", "T"]
2 points: ["D", "G"]
3 points: ["B", "C", "M", "P"]
4 points: ["F", "H", "V", "W", "Y"]
5 points: ["K"]
8 points: ["J", "X"]
10 points: ["Q", "Z"]
```

### New Data Structure (Individual Mapping)
```go
"a": 1
"b": 3
"c": 3
"d": 2
"e": 1
"f": 4
...
"z": 10
```

## Why This Transformation is Needed

1. **Internationalization**: Different languages have different letter frequencies
    - English: 'C' = 3 points (common)
    - Norwegian: 'C' = 10 points (rare)
2. **Flexibility**: Easier to add special characters for different languages
3. **Simplified Code**: Direct lookup instead of searching through groups
4. **Case Normalization**: All letters are converted to lowercase for consistency

### Basic Transformation

```go
package main

import (
    "fmt"
)

func main() {
    // Original grouped scores
    oldEnglishScores := map[int][]string{
        1:  {"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"},
        2:  {"D", "G"},
        3:  {"B", "C", "M", "P"},
        4:  {"F", "H", "V", "W", "Y"},
        5:  {"K"},
        8:  {"J", "X"},
        10: {"Q", "Z"},
    }

    // Transform to individual mapping
    newEnglishScores := Transform(oldEnglishScores)

    // Use the transformed scores
    fmt.Println("Score for 'c':", newEnglishScores["c"]) // Output: 3
    fmt.Println("Score for 'z':", newEnglishScores["z"]) // Output: 10
}
```

### The Transform Function

```go
// Transform converts a map of scores to letter groups into a map of individual letters to scores
// Input: map[int][]string where key is score and value is array of uppercase letters
// Output: map[string]int where key is lowercase letter and value is score
func Transform(in map[int][]string) map[string]int {
    out := make(map[string]int)
    for score, letters := range in {
        for _, char := range letters {
            lowerChar := strings.ToLower(char)
            out[lowerChar] = score
        }
    }
    return out
}
```

## Example Output

### Input (Old Format)
```json
{
    "1": ["A", "E", "I", "O", "U", "L", "N", "R", "S", "T"],
    "2": ["D", "G"],
    "3": ["B", "C", "M", "P"]
}
```

### Output (New Format)
```json
{
    "a": 1,
    "e": 1,
    "i": 1,
    "o": 1,
    "u": 1,
    "l": 1,
    "n": 1,
    "r": 1,
    "s": 1,
    "t": 1,
    "d": 2,
    "g": 2,
    "b": 3,
    "c": 3,
    "m": 3,
    "p": 3
}
```

## Benefits of the New Structure

1. **Language Support**: Easy to add new language configurations
2. **Performance**: O(1) lookup time for letter scores
3. **Maintainability**: Clear, direct mapping without nested structures
4. **Extensibility**: Supports special characters and diacritics
5. **Consistency**: All letters are lowercase for uniform handling

## Adding Support for New Languages

```go
// Norwegian letter scores (example)
norwegianScores := map[int][]string{
    1:  {"A", "E", "I", "O", "U", "N", "R", "S", "T"},
    2:  {"D", "G", "K", "L", "M"},
    3:  {"B", "F", "H", "J", "P", "V"},
    4:  {"Æ", "Ø", "Å", "Y"},
    5:  {"W"},
    6:  {"C"},
    8:  {"X"},
    10: {"Q", "Z"},
}

// Transform for use in game
norwegianLetterScores := Transform(norwegianScores)
```

## Testing

Run the included example:
```bash
go run main.go
```

Expected output:
```
Score for 'c': 3
Score for 'z': 10
```

## Dependencies

- Go 1.16 or higher
- Standard library only (no external dependencies)

## License

This project is part of the Lexiconia game system. Developed for educational and practical purposes in game development internationalization.

Here's a complete `README.md` file for your anagram detection implementation in English:

# Anagram Detection

A Go implementation of an anagram detector that finds words which are anagrams of a given target word.

## Overview

An anagram is a word formed by rearranging the letters of another word. For example, "owns" is an anagram of "snow". A word is not considered an anagram of itself, so "stop" is not an anagram of "stop".

This implementation correctly handles:
- Case-insensitive comparisons (uppercase and lowercase letters are treated as equivalent)
- Non-identical words (a word is not an anagram of itself)
- UTF-8 characters (via Go's `rune` type)
- Preservation of original letter case in results

## Features

- **Case-insensitive comparison**: "PoTS" is recognized as an anagram of "sTOp"
- **Self-word exclusion**: "stop" is not considered an anagram of "stop"
- **UTF-8 support**: Handles non-ASCII characters correctly
- **Order preservation**: Returns anagrams in the same order as they appear in the candidate list

## Usage

```go
package main

import (
    "fmt"
)

func main() {
    target := "stone"
    candidates := []string{"stone", "tones", "banana", "tons", "notes", "Seton"}
    
    result := Detect(target, candidates)
    fmt.Println(result) // Output: [tones notes Seton]
}
```

## API Reference

### `Detect(subject string, candidates []string) []string`

Finds all anagrams of the subject word from the list of candidates.

**Parameters:**
- `subject`: The target word to find anagrams for
- `candidates`: A slice of candidate words to check

**Returns:**
- A slice containing all anagrams found, in the same order as they appear in the candidates list

### Example Results

```go
Detect("stone", []string{"stone", "tones", "banana", "tons", "notes", "Seton"})
// Returns: ["tones", "notes", "Seton"]

Detect("listen", []string{"enlists", "google", "inlets", "banana"})
// Returns: ["inlets"]

Detect("diaper", []string{"hello", "world", "zombies", "pants"})
// Returns: []
```

## Implementation Details

The implementation works by:

1. **Normalizing words**: Converting all letters to lowercase and sorting the characters
2. **Self-check**: Skipping candidates that are identical to the subject (case-insensitive)
3. **Comparison**: Comparing normalized forms of subject and candidate words

### Normalization Process

```go
func NormalizeWords(str string) string {
    lowered := strings.Map(unicode.ToLower, str)  // Convert to lowercase
    runes := []rune(lowered)                      // Convert to runes for UTF-8 support
    sort.Slice(runes, func(i, j int) bool {       // Sort characters
        return runes[i] < runes[j]
    })
    return string(runes)                          // Return sorted string
}
```

## Performance Considerations

- Time complexity: O(n * m log m) where n is the number of candidates and m is the length of the words
- Space complexity: O(m) for the normalization process
- Suitable for typical use cases with reasonable word lengths

## Testing

Run the tests with:

```bash
go test
```

### Test Cases Included

The implementation should handle:

1. **Basic anagrams**: Simple rearrangements
2. **Case variations**: Different letter cases
3. **Non-anagrams**: Words that aren't anagrams
4. **Self-exclusion**: The subject word itself
5. **Empty cases**: Empty subject or candidates
6. **UTF-8 characters**: Non-ASCII characters (e.g., "café", "naïve")

## Edge Cases

- **Empty strings**: Handled correctly
- **Single character words**: "a" can only be anagram of itself (excluded)
- **Multiple spaces**: Not applicable (words are alphabetic only per spec)
- **Special characters**: Not applicable (alphabetic ASCII and UTF-8 letters only)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for new functionality
4. Make your changes
5. Ensure all tests pass
6. Submit a pull request

## License

MIT License - see LICENSE file for details.


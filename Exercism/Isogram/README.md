# Isogram Checker

A Go program for determining whether a word or phrase is an isogram.

## 📖 What is an Isogram?

An **isogram** (also known as a "non-pattern word") is a word or phrase that contains no repeating letters, ignoring spaces, hyphens, and usually case. Examples include:
- "lumberjacks" (each letter appears only once)
- "background" (no repeating letters)
- "downstream" (all letters are unique)

```go
// Copy this function to your code
func IsIsogram(word string) bool {
    word = strings.ToUpper(word)
    m := make(map[rune]bool, len(word))
    
    for _, c := range word {
        if c == ' ' || c == '-' {
            continue
        }
        if _, exist := m[c]; exist {
            return false
        }
        m[c] = true
    }
    return true
}
```

### Basic Examples

```go
program main

import "fmt"

func main() {
    // Test words
    fmt.Println(IsIsogram("lumberjacks"))  // true
    fmt.Println(IsIsogram("background"))   // true
    fmt.Println(IsIsogram("downstream"))   // true
    fmt.Println(IsIsogram("six-year-old")) // true (hyphens ignored)
    fmt.Println(IsIsogram("isograms"))     // false ('s' repeats)
    
    // Phrases with spaces
    fmt.Println(IsIsogram("Emily Jung Schwartzkopf"))  // true
    fmt.Println(IsIsogram("the big dwarf"))            // false ('g' repeats)
}
```

### Interactive Example

```go
func main() {
    words := []string{
        "algorithm",
        "documentary",
        "hydropneumatic",
        "misconjugatedly",
        "uncopyrightable",
        "subdermatoglyphic",
        "hello",
        "bookkeeper",
    }
    
    fmt.Println("Isogram Check Results:")
    fmt.Println("======================")
    for _, word := range words {
        result := IsIsogram(word)
        status := "✓"
        if !result {
            status = "✗"
        }
        fmt.Printf("%s %-20s → %v\n", status, word, result)
    }
}
```

## 🎯 Features

- **Case-insensitive**: Converts all letters to uppercase before checking
- **Ignores special characters**: Spaces and hyphens are ignored
- **Efficient**: Uses a map with pre-allocated capacity for O(n) performance
- **Unicode-aware**: Works with runes, supporting international characters

## 🔧 How It Works

### Algorithm Steps:
1. **Normalize**: Convert the entire word to uppercase
2. **Initialize**: Create a map to track seen letters
3. **Iterate**: Process each character (rune) in the word
4. **Skip**: Ignore spaces (' ') and hyphens ('-')
5. **Check**: If a letter already exists in the map, return `false`
6. **Store**: Add new letters to the map
7. **Return**: If no duplicates found, return `true`

### Code Breakdown:

```go
func IsIsogram(word string) bool {
    // Step 1: Case normalization
    word = strings.ToUpper(word)
    
    // Step 2: Pre-allocate map for efficiency
    m := make(map[rune]bool, len(word))
    
    // Step 3-6: Character iteration and checking
    for _, c := range word {
        if c == ' ' || c == '-' {
            continue  // Skip allowed characters
        }
        if _, exist := m[c]; exist {
            return false  // Duplicate found
        }
        m[c] = true  // Mark as seen
    }
    
    // Step 7: No duplicates found
    return true
}
```

## 🧪 Test Cases

### Valid Isograms
```go
// Single words
"IsIsogram(\"lumberjacks\")"    // true
"IsIsogram(\"background\")"     // true
"IsIsogram(\"downstream\")"     // true

// With hyphens
"IsIsogram(\"six-year-old\")"   // true

// Long isograms
"IsIsogram(\"subdermatoglyphic\")"  // true (17 letters!)
"IsIsogram(\"uncopyrightable\")"    // true (15 letters)
```

### Invalid (Not Isograms)
```go
// Repeating letters
"IsIsogram(\"isograms\")"       // false (s repeats)
"IsIsogram(\"hello\")"          // false (l repeats)
"IsIsogram(\"bookkeeper\")"     // false (o, k, e repeat)

// Case doesn't matter
"IsIsogram(\"Alpha\")"          // false (A and a are same letter)
```

### Edge Cases
```go
// Empty string
"IsIsogram(\"\")"               // true (no repeating letters)

// Single character
"IsIsogram(\"a\")"              // true
"IsIsogram(\"A\")"              // true

// Only spaces/hyphens
"IsIsogram(\"   --  \")"        // true

// Unicode characters
"IsIsogram(\"café\")"           // true
"IsIsogram(\"naïve\")"          // false (i with diaeresis considered unique)
```

## 📊 Performance

- **Time Complexity**: O(n) where n is the number of characters
- **Space Complexity**: O(k) where k is the number of unique letters (at most 26 for English)
- **Optimizations**:
    - Pre-allocated map size reduces reallocations
    - Early exit on first duplicate
    - Uses `rune` for proper Unicode support

## 🔍 Use Cases

1. **Word games**: Useful for Scrabble, Boggle, or custom word puzzles
2. **Password validation**: Ensure no repeating characters
3. **Linguistic analysis**: Study word patterns and structures
4. **Educational tools**: Teach about letter frequency and word properties
5. **Data cleaning**: Find unique-letter words in text corpora

## 🌐 International Support

The function works with Unicode characters, but note:

- **Accented letters**: Treated as unique (e.g., 'é' ≠ 'e')
- **Non-Latin scripts**: Works but may have unexpected results with combined characters
- **Case folding**: Only ASCII case folding (A-Z → a-z)

For full Unicode case folding, you could modify the function:

```go
import "golang.org/x/text/cases"
import "golang.org/x/text/language"

func IsIsogramUnicode(word string) bool {
    c := cases.Fold()
    word = c.String(word)
    // ... rest of the function
}
```

## 🚀 Advanced Usage

### Find All Isograms in a List

```go
func FindIsograms(words []string) []string {
    var isograms []string
    for _, word := range words {
        if IsIsogram(word) {
            isograms = append(isograms, word)
        }
    }
    return isograms
}
```

### Check with Custom Ignored Characters

```go
func IsIsogramCustom(word string, ignoreChars string) bool {
    word = strings.ToUpper(word)
    m := make(map[rune]bool, len(word))
    
    for _, c := range word {
        // Check if character should be ignored
        if strings.ContainsRune(ignoreChars, c) {
            continue
        }
        if _, exist := m[c]; exist {
            return false
        }
        m[c] = true
    }
    return true
}

// Usage:
fmt.Println(IsIsogramCustom("O'Reilly", "'"))  // true (apostrophe ignored)
```

## 📝 Example: Complete Analysis Tool

```go
program main

import (
    "fmt"
    "strings"
)

func IsIsogram(word string) bool {
    word = strings.ToUpper(word)
    m := make(map[rune]bool, len(word))
    
    for _, c := range word {
        if c == ' ' || c == '-' {
            continue
        }
        if _, exist := m[c]; exist {
            return false
        }
        m[c] = true
    }
    return true
}

func AnalyzeWord(word string) {
    isIsogram := IsIsogram(word)
    uniqueLetters := countUniqueLetters(word)
    
    fmt.Printf("Word: %-20s\n", word)
    fmt.Printf("  Isogram: %v\n", isIsogram)
    fmt.Printf("  Length: %d\n", len(word))
    fmt.Printf("  Unique letters: %d\n", uniqueLetters)
    fmt.Println()
}

func countUniqueLetters(word string) int {
    word = strings.ToUpper(word)
    m := make(map[rune]bool)
    
    for _, c := range word {
        if c >= 'A' && c <= 'Z' {
            m[c] = true
        }
    }
    return len(m)
}

func main() {
    words := []string{
        "lumberjacks",
        "background",
        "downstream",
        "six-year-old",
        "bookkeeper",
        "uncopyrightable",
    }
    
    for _, word := range words {
        AnalyzeWord(word)
    }
}
```

## 🤝 Contributing

Feel free to:
1. Add support for more ignored characters
2. Implement case-insensitive Unicode comparison
3. Add benchmarking tests
4. Extend with statistical analysis features

## 📄 License

This code is provided as-is. Feel free to use, modify, and distribute according to your needs.

---

**Note**: This implementation follows the common definition used in word games where spaces and hyphens are ignored. Different contexts may have different rules for what constitutes an isogram.
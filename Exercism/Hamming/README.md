# Hamming Distance Calculator

A Go implementation for calculating the Hamming distance between two strings. The Hamming distance is the number of positions at which the corresponding symbols are different.

## 📖 What is Hamming Distance?

The Hamming distance between two strings of equal length is the number of positions at which the corresponding symbols are different. It's named after Richard Hamming, who introduced the concept in 1950.

**Example:**
```
String 1: GAGCCTACTAACGGGAT
String 2: CATCGTAATGACGGCCT
          ^ ^ ^  ^ ^    ^^
Differences: 7 positions
Hamming Distance: 7
```

## 🚀 Quick Start

### Prerequisites
- Go 1.16 or higher

# Run the example
go run hamming.go
```

## 📋 Usage

### Basic Usage
```go
package main

import (
    "errors"
    "fmt"
)

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("not equal length")
    }
    distance := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            distance++
        }
    }
    return distance, nil
}

func main() {
    strand1 := "GAGCCTACTAACGGGAT"
    strand2 := "CATCGTAATGACGGCCT"
    
    distance, err := Distance(strand1, strand2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Printf("Hamming distance: %d\n", distance)
    }
}
```

### Example Output
```
Hamming distance between
"GAGCCTACTAACGGGAT"
and
"CATCGTAATGACGGCCT"
is: 7
```

## 🧪 Examples

### Example 1: DNA Sequence Comparison
```go
func main() {
    // DNA sequence comparison
    sequences := []struct {
        name   string
        strand1 string
        strand2 string
    }{
        {"Example 1", "GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"},
        {"Example 2", "ACGTACGT", "ACGTACGT"},
        {"Example 3", "ACGTACGT", "TGCAACGT"},
    }
    
    for _, seq := range sequences {
        distance, err := Distance(seq.strand1, seq.strand2)
        if err != nil {
            fmt.Printf("%s: Error - %v\n", seq.name, err)
        } else {
            fmt.Printf("%s: Distance = %d\n", seq.name, distance)
        }
    }
}
```

### Example 2: Error Handling
```go
func main() {
    // Test cases with error handling
    testCases := []struct {
        a, b string
    }{
        {"ABC", "ABCD"},    // Different lengths
        {"", ""},           // Empty strings
        {"A", "A"},         // Same strings
        {"A", "B"},         // Different single char
    }
    
    for _, tc := range testCases {
        distance, err := Distance(tc.a, tc.b)
        fmt.Printf("Distance(%q, %q) = ", tc.a, tc.b)
        if err != nil {
            fmt.Printf("Error: %v\n", err)
        } else {
            fmt.Printf("%d\n", distance)
        }
    }
}
```

## 🔧 Advanced Usage

### Extended Implementation with Additional Features
```go
package main

import (
    "errors"
    "fmt"
    "strings"
)

// Distance calculates the Hamming distance between two strings
func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("strings must be of equal length")
    }
    
    distance := 0
    for i := 0; i < len(a); i++ {
        if a[i] != b[i] {
            distance++
        }
    }
    return distance, nil
}

// CaseInsensitiveDistance calculates Hamming distance ignoring case
func CaseInsensitiveDistance(a, b string) (int, error) {
    return Distance(strings.ToUpper(a), strings.ToUpper(b))
}

// VisualizeDifferences shows where the strings differ
func VisualizeDifferences(a, b string) (string, error) {
    if len(a) != len(b) {
        return "", errors.New("strings must be of equal length")
    }
    
    var builder strings.Builder
    builder.WriteString("String 1: " + a + "\n")
    builder.WriteString("String 2: " + b + "\n")
    builder.WriteString("Differs:  ")
    
    for i := 0; i < len(a); i++ {
        if a[i] == b[i] {
            builder.WriteString(" ")
        } else {
            builder.WriteString("^")
        }
    }
    
    distance, _ := Distance(a, b)
    builder.WriteString(fmt.Sprintf("\nDistance: %d", distance))
    
    return builder.String(), nil
}

func main() {
    strand1 := "GAGCCTACTAACGGGAT"
    strand2 := "CATCGTAATGACGGCCT"
    
    // Get visual representation
    visualization, err := VisualizeDifferences(strand1, strand2)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
    } else {
        fmt.Println(visualization)
    }
    
    // Case insensitive comparison
    dist, _ := CaseInsensitiveDistance("abc", "ABC")
    fmt.Printf("\nCase insensitive distance: %d\n", dist)
}
```

## 🧪 Testing

### Unit Tests
Create `hamming_test.go`:
```go
package main

import (
    "testing"
)

func TestDistance(t *testing.T) {
    tests := []struct {
        name     string
        a        string
        b        string
        expected int
        hasError bool
    }{
        {"empty strings", "", "", 0, false},
        {"single char same", "A", "A", 0, false},
        {"single char different", "A", "G", 1, false},
        {"example from description", "GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT", 7, false},
        {"different lengths", "AG", "AGT", 0, true},
        {"all different", "ACGT", "TGCA", 4, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Distance(tt.a, tt.b)
            
            if tt.hasError {
                if err == nil {
                    t.Errorf("Distance(%q, %q) expected error, got %d", tt.a, tt.b, got)
                }
            } else {
                if err != nil {
                    t.Errorf("Distance(%q, %q) unexpected error: %v", tt.a, tt.b, err)
                }
                if got != tt.expected {
                    t.Errorf("Distance(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.expected)
                }
            }
        })
    }
}
```

Run tests:
```bash
go test -v
```

### Benchmark Tests
```go
func BenchmarkDistance(b *testing.B) {
    strand1 := "GAGCCTACTAACGGGATCATCGTAATGACGGCCT"
    strand2 := "CATCGTAATGACGGCCTGAGCCTACTAACGGGAT"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = Distance(strand1, strand2)
    }
}
```

Run benchmarks:
```bash
go test -bench=.
```

## 📊 Performance

The implementation has:
- **Time Complexity:** O(n) where n is the length of strings
- **Space Complexity:** O(1)
- **Memory Usage:** Minimal - only stores the distance counter

## ⚠️ Limitations

1. **ASCII Only:** This implementation works correctly only with ASCII characters. For Unicode strings (like Arabic, Persian, Chinese), use a rune-based implementation.

2. **Equal Length Required:** Strings must be of equal length.

## 🔄 Alternative Implementations

### Unicode-Compatible Version
```go
func UnicodeDistance(a, b string) (int, error) {
    runesA := []rune(a)
    runesB := []rune(b)
    
    if len(runesA) != len(runesB) {
        return 0, errors.New("strings must have equal number of characters")
    }
    
    distance := 0
    for i := 0; i < len(runesA); i++ {
        if runesA[i] != runesB[i] {
            distance++
        }
    }
    return distance, nil
}
```

## 🎯 Applications

The Hamming distance is useful in:

1. **Bioinformatics**: Comparing DNA/RNA sequences
2. **Error Detection**: In telecommunications and computer networks
3. **Coding Theory**: Measuring code distance
4. **Machine Learning**: As a distance metric
5. **Data Analysis**: Comparing strings or sequences

## 🚨 Important Note

**This implementation only works with ASCII characters!** For international text support (Arabic, Persian, Chinese, etc.), use the Unicode-compatible version shown above.

---

**Happy Coding!** If you find this useful, please consider giving it a star ⭐
# DNA Nucleotide Counter

A Go program for counting nucleotides in DNA strands and validating their sequences.

## Description

This program implements a DNA analysis system that:
1. Defines a custom type for DNA strands
2. Counts the frequency of each nucleotide in a DNA strand
3. Returns an error if the DNA contains invalid nucleotides

## Code Structure

### Data Types

```go
// Histogram maps nucleotides to their frequency counts
type Histogram map[byte]int

// DNA represents a DNA strand as a sequence of nucleotides
type DNA string
```

### Counts Function

```go
func (d DNA) Counts() (Histogram, error)
```

**Parameters:**
- `d`: A DNA strand of type `DNA`

**Returns:**
- `Histogram`: A map of nucleotides to their frequency counts
- `error`: An error if invalid nucleotides are found (otherwise `nil`)

**Valid Nucleotides:**
- `'A'` (Adenine)
- `'C'` (Cytosine)
- `'G'` (Guanine)
- `'T'` (Thymine)

## Usage

```go
package main

import (
    "fmt"
)

func main() {
    // Example 1: Valid DNA strand
    strand := DNA("ACGTACGT")
    histogram, err := strand.Counts()
    
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Results:")
        fmt.Printf("A: %d\n", histogram['A'])
        fmt.Printf("C: %d\n", histogram['C'])
        fmt.Printf("G: %d\n", histogram['G'])
        fmt.Printf("T: %d\n", histogram['T'])
    }
    
    // Example 2: Invalid DNA strand
    invalidStrand := DNA("ACGTXACGT")
    _, err2 := invalidStrand.Counts()
    if err2 != nil {
        fmt.Println("Error for invalid string:", err2)
    }
}
```

## Sample Output

```
Results:
A: 2
C: 2
G: 2
T: 2

Error for invalid string:  invalid nucleotides
```

## Additional Tests

```go
// Test empty string
emptyStrand := DNA("")
result, err := emptyStrand.Counts()
// Output: A:0, C:0, G:0, T:0, err=nil

// Test string with single nucleotide type
singleStrand := DNA("AAAAA")
result, err := singleStrand.Counts()
// Output: A:5, C:0, G:0, T:0, err=nil

// Test string with lowercase letters (invalid)
lowercaseStrand := DNA("acgt")
_, err := lowercaseStrand.Counts()
// Output: err != nil
```

## Algorithm

1. **Initialization**: Create a `Histogram` with zero values for all valid nucleotides
2. **String traversal**: Examine each character in the DNA string
3. **Validation**: Check if each character is a valid nucleotide
4. **Counting**: Increment the counter for valid nucleotides
5. **Error handling**: Return error if invalid nucleotide is found

## Technical Details

- **Case sensitivity**: Only uppercase English letters are valid
- **Time complexity**: Linear O(n) for a string of length n
- **Space complexity**: Constant O(1) for the histogram
- **Safety**: Array bounds checking handled automatically by Go

## Alternative Implementation

The function can also be implemented as:

```go
func (d DNA) Counts() (Histogram, error) {
    h := Histogram{}
    
    // Initialize all keys
    for _, nuc := range []byte{'A', 'C', 'G', 'T'} {
        h[nuc] = 0
    }
    
    for i := 0; i < len(d); i++ {
        nuc := d[i]
        if nuc == 'A' || nuc == 'C' || nuc == 'G' || nuc == 'T' {
            h[nuc]++
        } else {
            return nil, fmt.Errorf("invalid nucleotide '%c' at position %d", nuc, i)
        }
    }
    
    return h, nil
}
```

## Requirements

- Go version 1.13 or higher
- No external dependencies required

## Running the Program

1. Save the file as `dna.go`
2. Run the program:
```bash
go run dna.go
```

3. Compile to executable:
```bash
go build -o dna_counter dna.go
./dna_counter
```

## Applications

- Genetic sequence analysis
- DNA data validation
- Bioinformatics education
- Academic and research projects

## Limitations

- Only supports uppercase English letters
- Does not recognize modified or unknown nucleotides
- May require optimization for very long strings

## License

This sample code is free for educational use.
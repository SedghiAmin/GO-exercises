# DNA to RNA Transcription

A Go program that transcribes DNA sequences into their RNA complements.

## Description

This program implements the biological process of transcription, where DNA sequences are converted into RNA sequences by replacing each nucleotide with its complement according to standard molecular biology rules.

## Transcription Rules

DNA contains four nucleotides: Adenine (A), Cytosine (C), Guanine (G), and Thymine (T).

RNA contains four nucleotides: Adenine (A), Cytosine (C), Guanine (G), and Uracil (U).

The transcription follows these complement rules:

```
DNA -> RNA
G   -> C
C   -> G
T   -> A
A   -> U
```

## Function

### `ToRNA(dna string) string`

Transcribes a DNA sequence into an RNA sequence.

**Parameters:**
- `dna` (string): A DNA sequence consisting of nucleotides A, C, G, T

**Returns:**
- `string`: The transcribed RNA sequence

**Behavior:**
- Valid nucleotides are transcribed according to the rules above
- Invalid characters are preserved as-is in the output
- The function is case-sensitive (only uppercase letters are processed)

## Usage

```go
package main

import "fmt"

func main() {
    // Basic transcription
    dna := "GCTA"
    rna := ToRNA(dna)
    fmt.Printf("DNA: %s -> RNA: %s\n", dna, rna)
    // Output: DNA: GCTA -> RNA: CGAU
    
    // Longer sequence
    dna2 := "ACGTGGTCTTAA"
    rna2 := ToRNA(dna2)
    fmt.Printf("DNA: %s -> RNA: %s\n", dna2, rna2)
    // Output: DNA: ACGTGGTCTTAA -> RNA: UGCACCAGAAUU
    
    // Sequence with invalid characters
    dna3 := "GCTAX"
    rna3 := ToRNA(dna3)
    fmt.Printf("DNA: %s -> RNA: %s\n", dna3, rna3)
    // Output: DNA: GCTAX -> RNA: CGAUX
}
```

## Examples

| DNA Input | RNA Output | Explanation |
|-----------|------------|-------------|
| "G" | "C" | Guanine transcribes to Cytosine |
| "C" | "G" | Cytosine transcribes to Guanine |
| "T" | "A" | Thymine transcribes to Adenine |
| "A" | "U" | Adenine transcribes to Uracil |
| "GCTA" | "CGAU" | Mixed sequence transcription |
| "" | "" | Empty string returns empty string |
| "XYZ" | "XYZ" | Invalid characters are preserved |

## Algorithm

The function works as follows:
1. Creates a byte slice with the same length as the input DNA string
2. Iterates through each character in the DNA string
3. Uses a switch statement to map each valid nucleotide to its RNA complement
4. Preserves any invalid characters as-is
5. Converts the byte slice back to a string and returns it

## Time and Space Complexity

- **Time Complexity**: O(n), where n is the length of the DNA string
- **Space Complexity**: O(n), for storing the RNA result

## Alternative Implementations

### Using a Map (More Flexible)
```go
func ToRNA(dna string) string {
    transcription := map[byte]byte{
        'G': 'C',
        'C': 'G',
        'T': 'A',
        'A': 'U',
    }
    
    rna := make([]byte, len(dna))
    for i := 0; i < len(dna); i++ {
        if complement, exists := transcription[dna[i]]; exists {
            rna[i] = complement
        } else {
            rna[i] = dna[i]
        }
    }
    return string(rna)
}
```

### With Error Handling
```go
func ToRNA(dna string) (string, error) {
    var result strings.Builder
    result.Grow(len(dna))
    
    for i := 0; i < len(dna); i++ {
        switch dna[i] {
        case 'G':
            result.WriteByte('C')
        case 'C':
            result.WriteByte('G')
        case 'T':
            result.WriteByte('A')
        case 'A':
            result.WriteByte('U')
        default:
            return "", fmt.Errorf("invalid nucleotide '%c' at position %d", dna[i], i)
        }
    }
    
    return result.String(), nil
}
```

## Testing

You can test the function with various inputs:

```go
func TestToRNA(t *testing.T) {
    testCases := []struct {
        dna      string
        expected string
    }{
        {"GCTA", "CGAU"},
        {"ACGT", "UGCA"},
        {"", ""},
        {"GGG", "CCC"},
        {"CAT", "GUA"},
    }
    
    for _, tc := range testCases {
        result := ToRNA(tc.dna)
        if result != tc.expected {
            t.Errorf("ToRNA(%q) = %q, want %q", tc.dna, result, tc.expected)
        }
    }
}
```

## Installation and Running

1. Save the code to a file named `dna_to_rna.go`
2. Run the program:
```bash
go run dna_to_rna.go
```

3. Expected output:
```
DNA: GCTA -> RNA: CGAU
DNA: ACGTGGTCTTAA -> RNA: UGCACCAGAAUU
DNA: GCTAX -> RNA: CGAUX
```

## Biological Context

In molecular biology, transcription is the first step of gene expression where a particular segment of DNA is copied into RNA by the enzyme RNA polymerase. This RNA strand (mRNA) then serves as a template for protein synthesis during translation.

## Limitations

- Only processes uppercase letters
- Invalid characters are preserved rather than causing an error
- Does not handle lowercase DNA sequences
- No support for modified nucleotides (like methylated cytosine)

## Requirements

- Go 1.13 or higher
- No external dependencies
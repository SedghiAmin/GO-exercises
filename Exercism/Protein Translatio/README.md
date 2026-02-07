# RNA to Protein Translator

A Go implementation of RNA codon translation to amino acid sequences, simulating the biological process of protein synthesis.

## Overview

This program translates RNA sequences into protein chains by converting 3-letter codons into their corresponding amino acids, following the standard genetic code with stop codon handling.

## The Genetic Code

The program uses the standard RNA codon table:

| Codon | Amino Acid | Codon | Amino Acid | Codon | Amino Acid | Codon | Amino Acid |
|-------|------------|-------|------------|-------|------------|-------|------------|
| AUG | Methionine (Start) | UUU | Phenylalanine | UUA | Leucine | UCU | Serine |
| UUC | Phenylalanine | UUG | Leucine | UCC | Serine | UCA | Serine |
| UCG | Serine | UAU | Tyrosine | UAC | Tyrosine | UGU | Cysteine |
| UGC | Cysteine | UGG | Tryptophan | UAA | **STOP** | UAG | **STOP** |
| UGA | **STOP** | | | | | | |

## Features

- ✅ Translates RNA sequences to protein chains
- ✅ Handles stop codons (UAA, UAG, UGA)
- ✅ Validates codon sequences
- ✅ Error handling for invalid inputs
- ✅ Multiple codon support for same amino acid

## Error Types

The program defines two custom error types:

1. **`ErrStop`**: Returned when a stop codon is encountered
2. **`ErrInvalidBase`**: Returned for invalid codons or malformed RNA sequences

## Functions

### `FromRNA(rna string) ([]string, error)`
Translates an entire RNA sequence into a protein chain.

**Parameters:**
- `rna`: RNA sequence string (must be multiple of 3)

**Returns:**
- Slice of amino acid names
- Error (if any)

**Behavior:**
- Stops translation at first stop codon
- Returns `ErrInvalidBase` for sequences not divisible by 3
- Returns `ErrInvalidBase` for invalid codons

### `FromCodon(codon string) (string, error)`
Translates a single codon to its amino acid.

**Parameters:**
- `codon`: 3-letter codon string

**Returns:**
- Amino acid name (empty string for stop codons)
- Error (if any)

**Behavior:**
- Returns amino acid name for valid codons
- Returns `ErrStop` for stop codons
- Returns `ErrInvalidBase` for invalid codons

## Usage Examples

### Basic Translation
```go
rna := "AUGUUUUCU"
proteins, err := FromRNA(rna)
// Result: ["Methionine", "Phenylalanine", "Serine"]
```

### Translation with Stop Codon
```go
rna := "AUGUUUUAAUCU"
proteins, err := FromRNA(rna)
// Result: ["Methionine", "Phenylalanine"]
// Translation stops at UAA (stop codon)
```

### Invalid Codon
```go
rna := "AUGXYZUCU"
proteins, err := FromRNA(rna)
// Result: Error: invalid codon
```

## Biological Context

### Translation Process
1. **Initiation**: AUG codon signals start of translation
2. **Elongation**: Codons are read in groups of three
3. **Termination**: STOP codon signals end of translation

### Key Concepts
- **Codon**: Group of 3 nucleotides that codes for an amino acid
- **Genetic Code**: Universal mapping of codons to amino acids
- **Redundancy**: Multiple codons can code for the same amino acid
- **Stop Signals**: Three codons signal translation termination

## Code Structure

```go
package main

// Error definitions
var (
    ErrStop        = errors.New("stop codon")
    ErrInvalidBase = errors.New("invalid codon")
)

// Codon to amino acid mapping
var acid map[string]string = map[string]string{
    // ... codon mappings
}

// Translation functions
func FromRNA(rna string) ([]string, error)
func FromCodon(codon string) (string, error)
```

## Testing

The program includes test cases in the main function:

```go
func main() {
    // Test normal translation
    rna1 := "AUGUUUUCU"
    proteins1, _ := FromRNA(rna1)
    // Output: ["Methionine", "Phenylalanine", "Serine"]
    
    // Test stop codon handling
    rna2 := "AUGUUUUAAUCU"
    proteins2, _ := FromRNA(rna2)
    // Output: ["Methionine", "Phenylalanine"]
    
    // Test error handling
    rna3 := "AUGXYZUCU"
    _, err3 := FromRNA(rna3)
    // Output: Error: invalid codon
}
```

## Compilation and Execution

```bash
# Save the code to rna_translator.go
go run rna_translator.go

# Output:
# === Basic Tests ===
# RNA: AUGUUUUCU
# Proteins: [Methionine Phenylalanine Serine]
# 
# RNA: AUGUUUUAAUCU
# Proteins: [Methionine Phenylalanine]
# 
# RNA: AUGXYZUCU
# Error: invalid codon
```

## Applications

1. **Bioinformatics**: Analyzing genetic sequences
2. **Education**: Teaching molecular biology concepts
3. **Research**: Prototyping genetic algorithms
4. **Biotechnology**: Simulating protein synthesis

## Limitations

- Only handles standard RNA codons
- No support for DNA sequences
- No frame-shift mutation handling
- Simple error handling without detailed diagnostics

## Extension Ideas

1. Add support for DNA sequences
2. Implement frame-shift detection
3. Add support for rare codons
4. Create reverse translation (protein to RNA)
5. Add sequence validation and cleaning

## License

This educational implementation is free for academic and learning purposes.
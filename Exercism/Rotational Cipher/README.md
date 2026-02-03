# Rotational Cipher (Caesar Cipher) in Go

A Go implementation of the classic rotational cipher, also known as the Caesar cipher. This cryptographic algorithm shifts letters in the alphabet by a specified key value while preserving non-alphabetic characters.

## Overview

The rotational cipher is one of the simplest and most widely known encryption techniques. It's a type of substitution cipher where each letter in the plaintext is shifted a certain number of places down or up the alphabet.

## Features

- ✅ Encrypts text using rotational cipher algorithm
- ✅ Preserves letter case (uppercase/lowercase)
- ✅ Maintains non-alphabetic characters (numbers, punctuation, spaces)
- ✅ Handles keys from 0 to 26 (full rotation)
- ✅ Includes comprehensive test cases
- ✅ Clean, efficient Go implementation

## Algorithm

For each character in the input string:
- **Uppercase letters (A-Z)**: `'A' + (char - 'A' + key) % 26`
- **Lowercase letters (a-z)**: `'a' + (char - 'a' + key) % 26`
- **Non-alphabetic characters**: Remain unchanged

### Mathematical Formula

The encryption formula for a character `c` with key `k` is:
- If `c` is uppercase: `E(c) = ((c - 'A' + k) mod 26) + 'A'`
- If `c` is lowercase: `E(c) = ((c - 'a' + k) mod 26) + 'a'`

## Installation

```bash
# Clone the repository or copy the code to a Go file
# Save as rotational_cipher.go
```

## Usage

### Basic Usage

```go
package main

import "fmt"

func main() {
    // Encrypt text with key 13 (ROT13)
    encrypted := RotationalCipher("Hello, World!", 13)
    fmt.Println(encrypted) // Output: "Uryyb, Jbeyq!"
    
    // Decrypt by applying the same cipher again
    decrypted := RotationalCipher(encrypted, 13)
    fmt.Println(decrypted) // Output: "Hello, World!"
}
```

### Function Signature

```go
func RotationalCipher(plain string, shiftKey int) string
```

**Parameters:**
- `plain`: The plaintext string to encrypt
- `shiftKey`: The encryption key (0-26)

**Returns:**
- Encrypted ciphertext string

## Examples

### Standard Rotations

| Input | Key | Output | Description |
|-------|-----|--------|-------------|
| `"omg"` | 5 | `"trl"` | Basic lowercase rotation |
| `"c"` | 0 | `"c"` | Zero key (no change) |
| `"Cool"` | 26 | `"Cool"` | Full rotation (no change) |
| `"OMG"` | 5 | `"TRL"` | Uppercase rotation |

### ROT13 (Common Usage)

```go
// ROT13 is its own inverse
RotationalCipher("The quick brown fox", 13)
// Returns: "Gur dhvpx oebja sbk"

RotationalCipher("Gur dhvpx oebja sbk", 13)
// Returns: "The quick brown fox"
```

### With Special Characters

```go
RotationalCipher("Testing 1 2 3 testing!", 4)
// Returns: "Xiwxmrk 1 2 3 xiwxmrk!"

RotationalCipher("Let's eat, Grandma!", 21)
// Returns: "Gzo'n zvo, Bmviyhz!"
```

## Running the Tests

The package includes built-in test cases. To run them:

```bash
# Run the program
go run rotational_cipher.go
```

### Expected Test Output

```
✓ ROT5  omg                                         → trl                                         (Expected: trl)
✓ ROT0  c                                           → c                                           (Expected: c)
✓ ROT26 Cool                                        → Cool                                        (Expected: Cool)
✓ ROT13 The quick brown fox jumps over the lazy dog. → Gur dhvpx oebja sbk whzcf bire gur ynml qbt. (Expected: Gur dhvpx oebja sbk whzcf bire gur ynml qbt.)
✓ ROT13 Gur dhvpx oebja sbk whzcf bire gur ynml qbt. → The quick brown fox jumps over the lazy dog. (Expected: The quick brown fox jumps over the lazy dog.)
✓ ROT5  OMG                                         → TRL                                         (Expected: TRL)
✓ ROT4  Testing 1 2 3 testing!                     → Xiwxmrk 1 2 3 xiwxmrk!                     (Expected: Xiwxmrk 1 2 3 xiwxmrk!)
✓ ROT21 Let's eat, Grandma!                        → Gzo'n zvo, Bmviyhz!                        (Expected: Gzo'n zvo, Bmviyhz!)
```

## Algorithm Details

### How It Works

1. **Input Processing**: The function creates a rune slice with the same length as the input string
2. **Character Iteration**: Each character is processed individually
3. **Case Detection**:
    - Uppercase letters (ASCII 65-90): `'A' <= char <= 'Z'`
    - Lowercase letters (ASCII 97-122): `'a' <= char <= 'z'`
4. **Rotation Calculation**:
    - Calculate position in alphabet (0-25)
    - Add the shift key
    - Apply modulo 26 to wrap around the alphabet
    - Convert back to ASCII character
5. **Character Preservation**: Non-alphabetic characters are copied unchanged

### Edge Cases Handled

- ✅ Key values greater than 26 (automatically wrapped using modulo)
- ✅ Negative keys (would need additional handling if required)
- ✅ Empty strings
- ✅ Strings with only non-alphabetic characters
- ✅ Mixed content (letters, numbers, punctuation, spaces)

## Performance

- **Time Complexity**: O(n) where n is the length of the input string
- **Space Complexity**: O(n) for the result rune slice
- **Memory Efficient**: Uses rune slices for proper Unicode handling

## Extending the Code

### Adding Decryption Function

```go
func RotationalDecipher(cipher string, shiftKey int) string {
    // Decryption is just encryption with the inverse key
    return RotationalCipher(cipher, 26 - (shiftKey % 26))
}
```

### Supporting Negative Keys

```go
func RotationalCipher(plain string, shiftKey int) string {
    // Normalize key to 0-25 range
    shiftKey = shiftKey % 26
    if shiftKey < 0 {
        shiftKey += 26
    }
    
    // ... rest of the implementation
}
```

## Limitations

1. **Only English Alphabet**: Supports only A-Z and a-z characters
2. **No Unicode Support**: Extended Unicode characters are treated as non-alphabetic
3. **Simple Cryptography**: Not suitable for secure encryption (educational purposes only)

## Common Use Cases

- **Educational Tool**: Teaching basic cryptography concepts
- **Puzzle Solving**: Cryptograms and puzzles often use Caesar ciphers
- **ROT13 Encoding**: Commonly used to hide spoilers or offensive content
- **Programming Exercises**: Common interview and coding challenge problem

## Security Note

⚠️ **Warning**: The Caesar cipher provides no real cryptographic security. It can be easily broken by:
- Brute force (only 25 possible keys)
- Frequency analysis
- Known plaintext attacks

Use only for educational purposes or entertainment.

## References

- [Caesar Cipher on Wikipedia](https://en.wikipedia.org/wiki/Caesar_cipher)
- [ROT13 Cipher](https://en.wikipedia.org/wiki/ROT13)
- [Go Language Documentation](https://golang.org/doc/)

## License

This implementation is provided for educational purposes. Feel free to use and modify as needed.
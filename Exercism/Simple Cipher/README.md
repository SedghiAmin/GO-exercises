# Vigenère Cipher Implementation in Go

This package provides a complete implementation of the Vigenère cipher encryption algorithm in Go, including Caesar cipher and shift cipher variations.

## Features

- **Caesar Cipher**: Fixed shift of 3 positions
- **Shift Cipher**: Custom shift values (1-25 or -1 to -25)
- **Vigenère Cipher**: Polyalphabetic substitution with a key
- **Input Validation**: Proper handling of invalid parameters
- **Case Insensitivity**: All inputs are normalized to lowercase
- **Non-letter Handling**: Non-alphabetic characters are ignored


## Usage

### Caesar Cipher

```go
cipher := NewCaesar()
encoded := cipher.Encode("hello world")  // "khoor zruog"
decoded := cipher.Decode("khoor zruog")  // "helloworld"
```

### Shift Cipher

```go
// Positive shift
shift5 := NewShift(5)
encoded := shift5.Encode("hello")  // "mjqqt"
decoded := shift5.Decode("mjqqt")  // "hello"

// Negative shift
shiftMinus3 := NewShift(-3)
encoded := shiftMinus3.Encode("hello")  // "ebiil"
decoded := shiftMinus3.Decode("ebiil")  // "hello"
```

### Vigenère Cipher

```go
vigenere := NewVigenere("abc")
encoded := vigenere.Encode("hello world")  // "hfnos xntme"
decoded := vigenere.Decode("hfnos xntme")  // "helloworld"
```

## API Reference

### Types

- `Cipher`: Interface with `Encode(string) string` and `Decode(string) string` methods
- `shift`: Implements shift cipher
- `vigenere`: Implements Vigenère cipher

### Functions

#### `NewCaesar() Cipher`
Creates a new Caesar cipher with shift value of 3.

#### `NewShift(distance int) Cipher`
Creates a new shift cipher with the specified distance.
- Valid distances: 1 to 25 or -1 to -25
- Returns `nil` for invalid distances

#### `NewVigenere(key string) Cipher`
Creates a new Vigenère cipher with the specified key.
- Key must contain only lowercase letters a-z
- Key cannot consist entirely of the letter 'a'
- Returns `nil` for invalid keys

## Features

- **Automatic Normalization**: All input is converted to lowercase
- **Non-letter Filtering**: Only alphabetic characters are processed
- **Key Wrapping**: Vigenère key repeats cyclically
- **Bidirectional**: Both Encode and Decode operations supported
- **Error Handling**: Returns nil for invalid parameters

## Testing

Run the tests:

```bash
go test -v
```

The test suite includes:
- Caesar cipher with and without symbols
- Shift cipher with positive and negative values
- Vigenère cipher with various keys
- Invalid input validation

## Examples

```go
package main

import (
    "fmt"
    "yourmodule/cipher"
)

func main() {
    // Caesar cipher
    caesar := cipher.NewCaesar()
    fmt.Println(caesar.Encode("Hello, World!"))  // "khoor zruog"
    
    // Shift cipher
    shift := cipher.NewShift(5)
    fmt.Println(shift.Encode("hello"))  // "mjqqt"
    
    // Vigenère cipher
    v := cipher.NewVigenere("gold")
    fmt.Println(v.Encode("attack at dawn"))  // "gvvtgk gd hntk"
}
```

## License

MIT License - see LICENSE file for details


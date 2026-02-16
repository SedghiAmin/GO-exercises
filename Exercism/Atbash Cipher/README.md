# Atbash Cipher Implementation in Go

This repository contains a simple implementation of the Atbash cipher in Go. The Atbash cipher is a ancient encryption system that substitutes each letter with its reverse in the alphabet.

## How It Works

The Atbash cipher maps each letter to its counterpart in the reversed alphabet:
- 'a' becomes 'z'
- 'b' becomes 'y'
- 'c' becomes 'x'
- ... and so on

### Rules
- **Letters**: Converted to lowercase and reversed
- **Numbers**: Remain unchanged
- **Punctuation**: Removed/ignored
- **Output**: Grouped in 5-character chunks for readability

## Code

```go
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var output strings.Builder
	
	// Remove spaces and convert to lowercase
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)
	
	// Apply Atbash cipher
	groupByFive := 0
	for _, c := range s {
		if unicode.IsLetter(c) {
			offset := c - 'a'
			output.WriteRune('z' - offset)
			groupByFive++
		} else if unicode.IsDigit(c) {
			output.WriteRune(c)
			groupByFive++
		}
		// Punctuation is ignored
	}
	
	// Handle empty input
	if output.String() == "" {
		return ""
	}
	
	// Group the result in 5-character chunks
	var result strings.Builder
	for i, c := range output.String() {
		if i > 0 && i%5 == 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(c)
	}
	
	return result.String()
}

func main() {
	// Test examples
	fmt.Println(Atbash("test"))                    // gvhg
	fmt.Println(Atbash("x123 yes"))                 // c123b vh
	fmt.Println(Atbash("gvhg"))                     // test
	fmt.Println(Atbash("hello world"))              // svool dliow
	fmt.Println(Atbash("Testing,1 2 3, testing."))  // gvhgr mt123 gvhgr mt
	fmt.Println(Atbash("The quick brown fox jumps over the lazy dog.")) // gsvjf rxpyi ldmul cqfnk hlevi gsvoz abwlt
}
```

## Algorithm Explanation

1. **Clean Input**: Remove spaces and convert to lowercase
2. **Transform Letters**: For each letter, find its distance from 'a' and subtract from 'z'
    - Formula: `'z' - (c - 'a')`
3. **Preserve Numbers**: Keep digits as-is
4. **Group Output**: Add a space every 5 characters for readability

## Key Features

- Simple and readable implementation
- Handles both letters and numbers correctly
- Removes punctuation automatically
- Groups output in 5-character chunks
- Case-insensitive (all output is lowercase)

## Example Walkthrough

Encoding "test":
1. Clean: "test" (no spaces to remove)
2. Transform:
    - 't' → 'g' (t is 19th letter → 19th from end is 'g')
    - 'e' → 'v'
    - 's' → 'h'
    - 't' → 'g'
3. Result: "gvhg" (no grouping needed for 4 chars)

## Usage

```go
encoded := Atbash("Hello World!")
fmt.Println(encoded) // "svool dliow"

decoded := Atbash("svool dliow")
fmt.Println(decoded) // "helloworld"
```


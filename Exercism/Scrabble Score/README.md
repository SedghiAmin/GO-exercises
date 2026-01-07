# Scrabble Score Calculator

A simple Go function to calculate word scores based on Scrabble letter values.

## 📝 Description

This Go program provides a `Score()` function that calculates the score of a given word according to standard Scrabble letter point values.

## ✨ Features

- Calculates Scrabble scores based on standard letter values
- Handles both uppercase and lowercase letters
- Empty string returns 0
- All non-specified letters default to 1 point (as per standard Scrabble rules)

## 📊 Letter Values

| Points | Letters                     |
|--------|-----------------------------|
| 1      | All other letters (default) |
| 2      | D, G                        |
| 3      | B, C, M, P                  |
| 4      | F, H, V, W, Y               |
| 5      | K                           |
| 8      | J, X                        |
| 10     | Q, Z                        |

## 🚀 Installation

```go
go get github.com/yourusername/scrabble-score
```

Or simply copy the `Score` function into your project.

## 💻 Usage

```go
package main

import (
    "fmt"
)

func main() {
    fmt.Println(Score("hello"))   // Output: 8
    fmt.Println(Score("world"))   // Output: 9
    fmt.Println(Score(""))        // Output: 0
    fmt.Println(Score("quirky"))  // Output: 22
    fmt.Println(Score("cabbage")) // Output: 14
}
```

## 🧪 Testing

Run the included examples or create your own test cases:

```go
// Additional test cases
fmt.Println(Score("quiz"))      // 22
fmt.Println(Score("javascript")) // 24
fmt.Println(Score("go"))        // 3
fmt.Println(Score("ZEBRA"))     // 16 (case-insensitive)
```

## 🔧 Function Signature

```go
func Score(word string) int
```

**Parameters:**
- `word string`: The word to score (case-insensitive)

**Returns:**
- `int`: The Scrabble score of the word

## 📈 Example Breakdown

- **"hello"** = h(1) + e(1) + l(1) + l(1) + o(1) = **5**
- **"world"** = w(4) + o(1) + r(1) + l(1) + d(2) = **9**
- **"quirky"** = q(10) + u(1) + i(1) + r(1) + k(5) + y(4) = **22**

## 🐛 Known Issues

The current implementation has a minor inefficiency:
- Line 9: `word = strings.ToUpper(word)` converts to uppercase
- Line 12: `strings.ToLower(word)` converts back to lowercase
- This could be optimized by using only one case conversion

## 🤝 Contributing

Feel free to submit issues and enhancement requests!

## 📄 License

MIT License - feel free to use this code in your projects.
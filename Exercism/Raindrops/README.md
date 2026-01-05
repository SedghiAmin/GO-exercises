# Raindrops (FizzBuzz Variation)

A Go implementation of the classic coding exercise that converts numbers to raindrop sounds based on their factors.

## Problem Description
Convert a number to a string based on its factors:
- If divisible by 3: add "Pling"
- If divisible by 5: add "Plang"
- If divisible by 7: add "Plong"
- If not divisible by any: return the number as string

## Usage

```go
package main

import "fmt"

func main() {
    fmt.Println(Convert(28))  // Plong
    fmt.Println(Convert(30))  // PlingPlang
    fmt.Println(Convert(34))  // 34
    fmt.Println(Convert(105)) // PlingPlangPlong
}
```

## Examples

| Input | Output | Reason |
|-------|--------|--------|
| 1 | "1" | No factors |
| 3 | "Pling" | Divisible by 3 |
| 5 | "Plang" | Divisible by 5 |
| 7 | "Plong" | Divisible by 7 |
| 6 | "Pling" | Divisible by 3 |
| 10 | "Plang" | Divisible by 5 |
| 14 | "Plong" | Divisible by 7 |
| 15 | "PlingPlang" | Divisible by 3 and 5 |
| 21 | "PlingPlong" | Divisible by 3 and 7 |
| 35 | "PlangPlong" | Divisible by 5 and 7 |
| 105 | "PlingPlangPlong" | Divisible by 3, 5, and 7 |

Simple implementation with clean, readable logic.
# Two Fer

A simple Go program that implements the classic "Two Fer" (Two for) string formatting exercise.

## What is Two Fer?
"Two Fer" is short for "Two for one". Given a name, it returns the string:
```
"One for X, one for me."
```

If no name is provided, it uses "you" as the default.

## Usage

```go
package main

import "fmt"

func main() {
    fmt.Println(ShareWith("Alice"))  // One for Alice, one for me.
    fmt.Println(ShareWith("Bob"))    // One for Bob, one for me.
    fmt.Println(ShareWith(""))       // One for you, one for me.
    fmt.Println(ShareWith(""))       // One for you, one for me.
}
```

## API

### `ShareWith(name string) string`
Returns the formatted string. If `name` is empty, uses "you" as default.

## Examples

| Input | Output |
|-------|--------|
| `""` | `"One for you, one for me."` |
| `"Alice"` | `"One for Alice, one for me."` |
| `"Bob"` | `"One for Bob, one for me."` |


A minimal, zero-dependency Go package for the classic coding exercise.
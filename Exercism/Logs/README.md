# Log Analysis (Go)

This exercise focuses on working with Unicode strings and runes in Go.

## Objectives

- Detect the application type from a log message using Unicode symbols
- Replace specific Unicode characters (runes) in a string
- Count characters correctly using UTF-8 rune counting

## Functions

### `Application(log string) string`
Identifies the application emitting the log based on a Unicode symbol:
- ❗ → recommendation
- 🔍 → search
- ☀ → weather  
  Returns `"default"` if no known symbol is found.

### `Replace(log string, oldRune, newRune rune) string`
Replaces all occurrences of `oldRune` with `newRune` in the given log string.

### `WithinLimit(log string, limit int) bool`
Checks whether the number of characters (runes) in the log is within the given limit.

## Key Concepts

- Unicode and runes
- UTF-8 string handling
- Iterating over strings using `range`

## How to Run

```bash
go run main.go

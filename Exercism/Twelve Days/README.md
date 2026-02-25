# Twelve Days of Christmas Song Generator

A Go implementation that generates the lyrics for the classic Christmas carol "The Twelve Days of Christmas".

## Description

This program generates the complete lyrics of "The Twelve Days of Christmas" song. Each verse builds upon the previous ones, listing all the gifts given on each day of Christmas.

## Features

- Generate individual verses for any day (1-12)
- Generate the complete song with all 12 verses
- Proper formatting with commas and "and" before the final gift
- Input validation for verse numbers

## Functions

### `Verse(i int) string`
Returns the lyrics for a specific day (verse) of the song.
- **Parameter**: `i` - Day number (1-12)
- **Returns**: String containing the complete verse for that day

### `Song() string`
Returns the complete lyrics for all 12 days of Christmas.
- **Returns**: Multi-line string containing all verses, each separated by a newline

## Usage Example

```go
package main

import (
    "fmt"
)

func main() {
    // Print a single verse
    fmt.Println(Verse(1))
    // Output: On the first day of Christmas my true love gave to me: a Partridge in a Pear Tree.
    
    fmt.Println(Verse(2))
    // Output: On the second day of Christmas my true love gave to me: two Turtle Doves, and a Partridge in a Pear Tree.
    
    // Print the complete song
    fmt.Println(Song())
}
```

## Output Format

Each verse follows this pattern:
```
On the [day] day of Christmas my true love gave to me: [gift 1], [gift 2], ..., and [gift N].
```

The gifts are listed in reverse chronological order (from the current day back to the first day), with proper comma placement and "and" before the final gift.

## Running the Program

```bash
go run main.go
```

This will output the complete lyrics of "The Twelve Days of Christmas" song.

## Error Handling

The `Verse` function returns an empty string if called with a day number less than 1.
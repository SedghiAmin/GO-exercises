# Resistor Color

A Go program for working with resistor color codes. This package provides functionality to get color values and lists for electronic resistor color coding.

## Description

This package implements the resistor color code system used in electronics to indicate the values of resistors. Each color corresponds to a specific numerical value.

## Installation

```bash
go get yourmodule/resistorcolor
```

## Usage

```go
package main

import (
    "fmt"
    "yourmodule/resistorcolor"
)

func main() {
    // Get all available colors
    allColors := resistorcolor.Colors()
    fmt.Println("All colors:", allColors)
    
    // Get the numeric value for a specific color
    value := resistorcolor.ColorCode("red")
    fmt.Println("Value for 'red':", value)
    
    value = resistorcolor.ColorCode("blue")
    fmt.Println("Value for 'blue':", value)
}
```

## Functions

### `Colors() []string`
Returns a slice containing all the available color names in the resistor color code system.

**Returns:**
- `[]string`: A slice of color names in the standard order: black, brown, red, orange, yellow, green, blue, violet, grey, white.

### `ColorCode(color string) int`
Returns the numeric value associated with a given color name.

**Parameters:**
- `color` (string): The name of the color (case-sensitive, lowercase)

**Returns:**
- `int`: The numeric value (0-9) associated with the color

## Color Mapping

| Color   | Value |
|---------|-------|
| black   | 0     |
| brown   | 1     |
| red     | 2     |
| orange  | 3     |
| yellow  | 4     |
| green   | 5     |
| blue    | 6     |
| violet  | 7     |
| grey    | 8     |
| white   | 9     |

## Example Output

```bash
All colors: [black brown red orange yellow green blue violet grey white]
Value for 'red': 2
Value for 'blue': 6
```

## Error Handling

The `ColorCode` function will return `0` for any color not found in the map (due to Go's zero value for int). For production use, you might want to add error checking:

```go
func ColorCode(color string) (int, error) {
    if value, ok := colorCodes[color]; ok {
        return value, nil
    }
    return 0, fmt.Errorf("invalid color: %s", color)
}
```

## Testing

To run tests (if you have test files):

```bash
go test ./...
```

## Dependencies

- Go 1.13 or higher (no external dependencies)

## License

[MIT](LICENSE)
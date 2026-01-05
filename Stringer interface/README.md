## Overview

This program provides a simple yet powerful way to work with distance measurements in Go. It supports multiple distance units and provides clean string representations for display purposes.

## Features

- **Multiple Units**: Support for kilometers and miles
- **Type Safety**: Strongly typed distance measurements
- **String Representation**: Clean formatting for display
- **Extensible Design**: Easy to add new distance units

### Creating Distances

```go
// Method 1: Direct struct initialization
d1 := Distant{number: 50.0, unit: Kilometer}

// Method 2: Positional initialization
d2 := Distant{31.07, Mile}

// Method 3: Zero value
var d3 Distant // number: 0.0, unit: 0 (Kilometer)
```

### Working with Units

```go
// Unit constants
fmt.Println(Kilometer) // Output: km
fmt.Println(Mile)      // Output: mi

// Check unit types
if d.unit == Kilometer {
    fmt.Println("Distance is in kilometers")
} else if d.unit == Mile {
    fmt.Println("Distance is in miles")
}
```

## API Reference

### Types

#### `DistanceMeasureUnit`
An integer type representing distance measurement units.

```go
type DistanceMeasureUnit int
```

Constants:
- `Kilometer` (0): Represents kilometers
- `Mile` (1): Represents miles

#### `Distant`
A struct representing a distance with a numerical value and unit.

```go
type Distant struct {
    number float64
    unit   DistanceMeasureUnit
}
```

### Methods

#### `(d DistanceMeasureUnit) String() string`
Returns the string representation of a distance unit.

```go
fmt.Println(Kilometer.String()) // "km"
fmt.Println(Mile.String())      // "mi"
```

#### `(d Distant) String() string`
Returns the formatted string representation of a distance.

```go
d := Distant{100.0, Kilometer}
fmt.Println(d.String()) // "100 km"
```

## Examples

### Example 1: Basic Distance Creation

```go
package main

import "fmt"

func main() {
    distances := []Distant{
        {10.5, Kilometer},
        {5.2, Mile},
        {42.195, Kilometer}, // Marathon distance
        {26.2, Mile},        // Marathon in miles
    }
    
    for _, d := range distances {
        fmt.Printf("Distance: %v\n", d)
    }
}
```

Output:
```
Distance: 10.5 km
Distance: 5.2 mi
Distance: 42.195 km
Distance: 26.2 mi
```

### Example 2: Unit Conversion

```go
package main

import "fmt"

func main() {
    // Simple conversion example
    nycToBoston := Distant{306.0, Kilometer}
    fmt.Printf("NYC to Boston: %v\n", nycToBoston)
    
    // In miles (approximate)
    nycToBostonMiles := Distant{190.0, Mile}
    fmt.Printf("NYC to Boston: %v\n", nycToBostonMiles)
}
```

## Extending the Library

### Adding New Units

To add a new distance unit:

1. Add a new constant to `DistanceMeasureUnit`:

```go
const (
    Kilometer DistanceMeasureUnit = 0
    Mile      DistanceMeasureUnit = 1
    Meter     DistanceMeasureUnit = 2
)
```

2. Update the `String()` method:

```go
func (d DistanceMeasureUnit) String() string {
    units := []string{"km", "mi", "m"}
    if int(d) < len(units) {
        return units[d]
    }
    return "unknown"
}
```

### Adding Conversion Methods

You can extend the library with conversion methods:

```go
func (d Distant) ToKilometers() float64 {
    if d.unit == Mile {
        return d.number * 1.60934
    }
    return d.number
}

func (d Distant) ToMiles() float64 {
    if d.unit == Kilometer {
        return d.number * 0.621371
    }
    return d.number
}
```

## Testing

Run the tests with:

```bash
go test -v
```

Example test:

```go
func TestDistantString(t *testing.T) {
    tests := []struct {
        name     string
        distant  Distant
        expected string
    }{
        {"Kilometer", Distant{100, Kilometer}, "100 km"},
        {"Mile", Distant{62.1, Mile}, "62.1 mi"},
        {"Zero", Distant{0, Kilometer}, "0 km"},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := tt.distant.String(); got != tt.expected {
                t.Errorf("String() = %v, want %v", got, tt.expected)
            }
        })
    }
}
```

## Performance

The library is designed for performance:
- No heap allocations for basic operations
- Minimal memory footprint
- Fast string conversions using pre-allocated arrays

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for your changes
4. Make your changes
5. Ensure all tests pass
6. Submit a pull request

**Note**: This is a basic implementation. For production use, consider adding:
- Unit conversion methods
- Arithmetic operations (addition, subtraction)
- Comparison methods
- JSON serialization/deserialization
- Validation methods
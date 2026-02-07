# Clock Implementation in Go

A simple and robust 24-hour clock implementation in Go that handles time arithmetic with rollover support.

## Features

- ✅ 24-hour clock representation
- ✅ Time addition and subtraction with automatic rollover
- ✅ Handling of negative hours and minutes
- ✅ Proper normalization across day boundaries
- ✅ Clean string formatting (HH:MM)
- ✅ No external dependencies

## Installation

```bash
# Clone the repository
git clone <your-repo-url>
cd go-clock

# Run the program
go run clock.go
```

## Usage

### Creating a Clock

```go
// Create a new clock at 14:30
c := New(14, 30)
fmt.Println(c) // Output: 14:30
```

### Adding Minutes

```go
c := New(10, 30)
c = c.Add(45)  // Add 45 minutes
fmt.Println(c) // Output: 11:15
```

### Subtracting Minutes

```go
c := New(10, 30)
c = c.Subtract(90) // Subtract 90 minutes
fmt.Println(c)     // Output: 09:00
```

### Edge Cases

The clock automatically handles:
- Rollover past midnight
- Negative hours and minutes
- Large time additions/subtractions

```go
fmt.Println(New(23, 59).Add(1))     // 00:00
fmt.Println(New(0, 0).Subtract(1))  // 23:59
fmt.Println(New(25, 0))             // 01:00 (25 hours = 1:00 next day)
fmt.Println(New(-1, 15))            // 23:15 (1 hour before midnight)
```

## API Reference

### Types

```go
type Clock struct {
    hour   int
    minute int
}
```

### Functions

#### `New(h, m int) Clock`
Creates a new Clock instance with the specified hours and minutes.

**Parameters:**
- `h`: Hours (can be negative or greater than 23)
- `m`: Minutes (can be negative or greater than 59)

**Returns:** A normalized Clock instance

#### `(c Clock) Add(m int) Clock`
Adds minutes to the clock.

**Parameters:**
- `m`: Minutes to add (can be negative)

**Returns:** A new Clock instance with the added time

#### `(c Clock) Subtract(m int) Clock`
Subtracts minutes from the clock.

**Parameters:**
- `m`: Minutes to subtract (can be negative)

**Returns:** A new Clock instance with the subtracted time

#### `(c Clock) String() string`
Returns the clock time in "HH:MM" format.

**Returns:** Formatted string representation

## Implementation Details

The clock uses a normalization function that:
1. Converts hours and minutes to total minutes
2. Handles negative values by adding full days (24 hours)
3. Uses modulo operations to keep time within 0-23 hours and 0-59 minutes

### Normalization Logic

```go
func normalize(h, m int) Clock {
    minutes := h*60 + m
    
    // Handle negative times
    for minutes < 0 {
        minutes += 24 * 60
    }
    
    return Clock{
        hour:   (minutes / 60) % 24,
        minute: minutes % 60,
    }
}
```

## Examples

```go
package main

import "fmt"

func main() {
    // Basic usage
    clock := New(14, 45)
    fmt.Printf("Initial time: %s\n", clock)
    
    // Add time
    clock = clock.Add(30)
    fmt.Printf("After adding 30 minutes: %s\n", clock)
    
    // Subtract time
    clock = clock.Subtract(90)
    fmt.Printf("After subtracting 90 minutes: %s\n", clock)
    
    // Edge cases
    fmt.Println("\nEdge Cases:")
    fmt.Printf("23:59 + 1 minute = %s\n", New(23, 59).Add(1))
    fmt.Printf("00:00 - 1 minute = %s\n", New(0, 0).Subtract(1))
    fmt.Printf("25:00 normalized = %s\n", New(25, 0))
    fmt.Printf("-1:15 normalized = %s\n", New(-1, 15))
}
```

## Running Tests

```bash
# Create a test file
cat > clock_test.go << 'EOF'
package main

import (
    "fmt"
    "testing"
)

func TestClock(t *testing.T) {
    tests := []struct {
        h, m int
        want string
    }{
        {14, 45, "14:45"},
        {23, 59, "23:59"},
        {25, 0, "01:00"},
        {-1, 15, "23:15"},
    }
    
    for _, tt := range tests {
        got := New(tt.h, tt.m).String()
        if got != tt.want {
            t.Errorf("New(%d, %d) = %s; want %s", tt.h, tt.m, got, tt.want)
        }
    }
}
EOF

# Run tests
go test -v
```

## Design Decisions

1. **Immutable Design**: All operations return new Clock instances
2. **24-hour Format**: Consistent 24-hour representation
3. **Automatic Normalization**: Times are always valid
4. **Simple API**: Minimal, intuitive methods

## Limitations

- Only supports minute-precision (no seconds)
- 24-hour format only (no AM/PM)
- No timezone support

## Contributing

Feel free to submit issues and enhancement requests!

## License

MIT License
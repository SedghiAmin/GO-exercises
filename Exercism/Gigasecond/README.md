# Gigasecond Calculator in Go

A Go implementation that adds one gigasecond (1,000,000,000 seconds) to any given date and time.

## 📊 What is a Gigasecond?

A **gigasecond** is exactly **1,000,000,000 seconds**, which is approximately:
- **31.69 years**
- **11,574.07 days**
- **16,666,666.67 minutes**
- **277,777.78 hours**

This is a fun programming exercise that demonstrates date/time manipulation in Go.

## 🚀 Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/gigasecond-go.git
cd gigasecond-go

# Run the program
go run gigasecond.go
```

## 📋 Prerequisites

- Go 1.16 or higher
- Basic understanding of Go programming

## 🛠️ Usage

### Basic Function

```go
package main

import (
    "fmt"
    "time"
)

func AddGigasecond(t time.Time) time.Time {
    return t.Add(1000000000 * time.Second)
}
```

### Example Usage

```go
func main() {
    // Example 1: Current time
    now := time.Now()
    future := AddGigasecond(now)
    fmt.Printf("Current time: %v\n", now)
    fmt.Printf("One gigasecond later: %v\n", future)
    
    // Example 2: Specific date
    birthDate := time.Date(1990, time.January, 15, 8, 30, 0, 0, time.UTC)
    gigasecondAnniversary := AddGigasecond(birthDate)
    fmt.Printf("\nBirth date: %v\n", birthDate.Format("January 2, 2006"))
    fmt.Printf("Gigasecond anniversary: %v\n", gigasecondAnniversary.Format("January 2, 2006"))
    
    // Example 3: Unix epoch
    epoch := time.Unix(0, 0)
    epochPlusGigasecond := AddGigasecond(epoch)
    fmt.Printf("\nUnix epoch: %v\n", epoch)
    fmt.Printf("Epoch + 1 gigasecond: %v\n", epochPlusGigasecond)
}
```

### Running Tests

Create a test file `gigasecond_test.go`:

```go
package main

import (
    "testing"
    "time"
)

func TestAddGigasecond(t *testing.T) {
    tests := []struct {
        name string
        input time.Time
        expected time.Time
    }{
        {
            name: "From Unix epoch",
            input: time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
            expected: time.Date(2001, 9, 9, 1, 46, 40, 0, time.UTC),
        },
        {
            name: "From specific date",
            input: time.Date(2015, 1, 24, 22, 0, 0, 0, time.UTC),
            expected: time.Date(2046, 10, 2, 23, 46, 40, 0, time.UTC),
        },
        {
            name: "Leap year date",
            input: time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC),
            expected: time.Date(2051, 11, 7, 1, 46, 40, 0, time.UTC),
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := AddGigasecond(tt.input)
            if !result.Equal(tt.expected) {
                t.Errorf("AddGigasecond(%v) = %v, want %v", tt.input, result, tt.expected)
            }
        })
    }
}

func TestAddGigasecondEdgeCases(t *testing.T) {
    // Test with different timezones
    loc, _ := time.LoadLocation("America/New_York")
    nyTime := time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
    result := AddGigasecond(nyTime)
    
    // Verify that duration is correct regardless of timezone
    duration := result.Sub(nyTime)
    expectedDuration := 1000000000 * time.Second
    
    if duration != expectedDuration {
        t.Errorf("Duration mismatch: got %v, want %v", duration, expectedDuration)
    }
}
```

Run tests with:
```bash
go test -v
```

## 📚 API Reference

### `AddGigasecond(t time.Time) time.Time`

Adds exactly one gigasecond (1,000,000,000 seconds) to the given time.

**Parameters:**
- `t time.Time` - The starting time

**Returns:**
- `time.Time` - The time exactly one gigasecond after the input

**Time Complexity:** O(1)
**Space Complexity:** O(1)

## 🌟 Examples

### Example 1: Calculate Your Gigasecond Birthday

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Enter your birth date
    birthDate := time.Date(1985, time.May, 10, 14, 30, 0, 0, time.UTC)
    
    gigasecondBirthday := AddGigasecond(birthDate)
    
    fmt.Println("🎉 Gigasecond Birthday Calculator 🎉")
    fmt.Printf("Birth date:       %s\n", birthDate.Format("Monday, January 2, 2006 at 15:04:05"))
    fmt.Printf("Gigasecond day:   %s\n", gigasecondBirthday.Format("Monday, January 2, 2006 at 15:04:05"))
    
    // Calculate age at gigasecond
    age := gigasecondBirthday.Year() - birthDate.Year()
    fmt.Printf("You'll be %d years old!\n", age)
}
```

### Example 2: Historical Events

```go
func main() {
    historicalEvents := map[string]time.Time{
        "Moon Landing":        time.Date(1969, time.July, 20, 20, 17, 40, 0, time.UTC),
        "Fall of Berlin Wall": time.Date(1989, time.November, 9, 0, 0, 0, 0, time.UTC),
        "Y2K":                 time.Date(1999, time.December, 31, 23, 59, 59, 0, time.UTC),
        "First iPhone":        time.Date(2007, time.June, 29, 0, 0, 0, 0, time.UTC),
    }
    
    fmt.Println("📅 Historical Events + 1 Gigasecond 📅")
    fmt.Println("======================================")
    
    for event, date := range historicalEvents {
        future := AddGigasecond(date)
        fmt.Printf("%-25s: %s\n", event, date.Format("Jan 2, 2006"))
        fmt.Printf("%-25s  %s\n", "", future.Format("→ Jan 2, 2006"))
        fmt.Println()
    }
}
```

### Example 3: Batch Processing

```go
func ProcessDates(dates []time.Time) []time.Time {
    results := make([]time.Time, len(dates))
    for i, date := range dates {
        results[i] = AddGigasecond(date)
    }
    return results
}

func main() {
    dates := []time.Time{
        time.Now(),
        time.Now().Add(-365 * 24 * time.Hour), // 1 year ago
        time.Now().Add(365 * 24 * time.Hour),  // 1 year from now
    }
    
    futureDates := ProcessDates(dates)
    
    for i, date := range dates {
        fmt.Printf("Original: %s\n", date.Format("2006-01-02"))
        fmt.Printf("Future:   %s\n\n", futureDates[i].Format("2006-01-02"))
    }
}
```

## 🔧 Advanced Features

### With Timezone Support

```go
func AddGigasecondInLocation(t time.Time, loc *time.Location) time.Time {
    // Convert to target location, add gigasecond, return in same location
    return t.In(loc).Add(1000000000 * time.Second)
}

func main() {
    utcTime := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
    
    locations := []string{
        "America/New_York",
        "Europe/London",
        "Asia/Tokyo",
        "Australia/Sydney",
    }
    
    for _, locName := range locations {
        loc, _ := time.LoadLocation(locName)
        result := AddGigasecondInLocation(utcTime, loc)
        fmt.Printf("UTC: %s -> %s: %s\n",
            utcTime.Format("15:04"),
            locName,
            result.Format("2006-01-02 15:04:05 MST"))
    }
}
```

### Performance Benchmark

```go
func BenchmarkAddGigasecond(b *testing.B) {
    baseTime := time.Now()
    for i := 0; i < b.N; i++ {
        _ = AddGigasecond(baseTime)
    }
}
```

## 📊 Interesting Calculations

```go
func main() {
    // Some interesting gigasecond calculations
    fmt.Println("🔢 Interesting Gigasecond Facts 🔢")
    fmt.Println("=================================")
    
    // Unix Epoch + 1 gigasecond
    epoch := time.Unix(0, 0)
    fmt.Printf("Unix Epoch:           %s\n", epoch.Format("Jan 2, 2006"))
    fmt.Printf("+1 gigasecond:        %s\n", AddGigasecond(epoch).Format("Jan 2, 2006"))
    
    // Today
    today := time.Now()
    fmt.Printf("\nToday:                %s\n", today.Format("Jan 2, 2006"))
    fmt.Printf("+1 gigasecond:        %s\n", AddGigasecond(today).Format("Jan 2, 2006"))
    
    // How many gigaseconds old are you?
    birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
    ageInSeconds := time.Since(birthDate).Seconds()
    gigasecondsOld := ageInSeconds / 1000000000
    fmt.Printf("\nIf born on Jan 1, 1990:\n")
    fmt.Printf("You are %.4f gigaseconds old\n", gigasecondsOld)
}
```

## 🧪 Testing with Table-Driven Tests

```go
func TestAddGigasecond_TableDriven(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        expected string
        layout   string
    }{
        {
            name:     "2011-04-25",
            input:    "2011-04-25T00:00:00Z",
            expected: "2043-01-01T01:46:40Z",
            layout:   time.RFC3339,
        },
        {
            name:     "1977-06-13",
            input:    "1977-06-13T00:00:00Z",
            expected: "2009-02-19T01:46:40Z",
            layout:   time.RFC3339,
        },
        {
            name:     "1959-07-19",
            input:    "1959-07-19T00:00:00Z",
            expected: "1991-03-27T01:46:40Z",
            layout:   time.RFC3339,
        },
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            input, _ := time.Parse(tc.layout, tc.input)
            expected, _ := time.Parse(tc.layout, tc.expected)
            
            got := AddGigasecond(input)
            
            if !got.Equal(expected) {
                t.Errorf("AddGigasecond(%v) = %v, want %v", input, got, expected)
            }
        })
    }
}
```

## 📦 Dependencies

This package has no external dependencies beyond the Go standard library:
- `time` - For date/time manipulation
- `fmt` - For formatted I/O

## 📝 License

MIT License - see [LICENSE](LICENSE) file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 🙏 Acknowledgments

- Inspired by [Exercism's Gigasecond exercise](https://exercism.org/tracks/go/exercises/gigasecond)
- Go time package documentation
- The concept of "gigasecond birthdays" in programming communities

## 📚 Further Reading

- [Go time package documentation](https://pkg.go.dev/time)
- [Exercism - Gigasecond](https://exercism.org/tracks/go/exercises/gigasecond)
- [Wikipedia - Gigasecond](https://en.wikipedia.org/wiki/Gigasecond)

---

**Happy Coding!** 🎉 Remember: Your gigasecond birthday is a special milestone that happens only once in a lifetime (literally)!
# Leap Year Checker

A simple Go package to determine whether a given year is a leap year according to the Gregorian calendar rules.

## 📋 Overview

This Go package provides a clean, efficient function to check if a year is a leap year. It implements the standard Gregorian calendar leap year rules in a single, readable function.

## 🚀 Installation

Simply copy the `IsLeapYear` function into your project:

```go
// Copy this function to your code
func IsLeapYear(year int) bool {
    if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
        return true
    }
    return false
}
```

Or import as a package if you structure it as a module.

## 📖 Usage

### Basic Example

```go
package main

import "fmt"

func main() {
    // Test various years
    fmt.Println(IsLeapYear(2024))  // true
    fmt.Println(IsLeapYear(2023))  // false
    fmt.Println(IsLeapYear(1900))  // false
    fmt.Println(IsLeapYear(2000))  // true
    fmt.Println(IsLeapYear(2020))  // true
    fmt.Println(IsLeapYear(2100))  // false
}
```

### Using in Conditional Statements

```go
func main() {
    year := 2024
    
    if IsLeapYear(year) {
        fmt.Printf("%d is a leap year. February has 29 days.\n", year)
    } else {
        fmt.Printf("%d is not a leap year. February has 28 days.\n", year)
    }
}
```

### Bulk Year Checking

```go
func main() {
    years := []int{2019, 2020, 2021, 2022, 2023, 2024}
    
    fmt.Println("Leap years in the list:")
    for _, year := range years {
        if IsLeapYear(year) {
            fmt.Printf("  %d\n", year)
        }
    }
}
```

## 🧠 Leap Year Rules

The function implements the official Gregorian calendar leap year rules:

1. **Divisible by 4?** → Yes, go to step 2. No → Not a leap year.
2. **Divisible by 100?** → Yes, go to step 3. No → Leap year.
3. **Divisible by 400?** → Yes → Leap year. No → Not a leap year.

### Examples:
- **2024**: ✅ Divisible by 4, ❌ not by 100 → **Leap year**
- **1900**: ✅ Divisible by 4, ✅ by 100, ❌ not by 400 → **Not leap year**
- **2000**: ✅ Divisible by 4, ✅ by 100, ✅ by 400 → **Leap year**
- **2023**: ❌ Not divisible by 4 → **Not leap year**

## 📊 Algorithm Breakdown

The function uses a single logical expression:

```go
(year%4 == 0) && (year%100 != 0 || year%400 == 0)
```

This translates to:
- Must be divisible by 4 **AND**
- Either NOT divisible by 100 **OR** divisible by 400

## 🧪 Testing

### Test Cases

```go
func TestIsLeapYear(t *testing.T) {
    testCases := []struct {
        year     int
        expected bool
    }{
        {2000, true},   // Divisible by 400
        {1900, false},  // Divisible by 100 but not 400
        {2024, true},   // Divisible by 4 but not 100
        {2023, false},  // Not divisible by 4
        {2020, true},   // Divisible by 4 but not 100
        {2100, false},  // Divisible by 100 but not 400
        {2400, true},   // Divisible by 400
        {1, false},     // Early year
        {4, true},      // First leap year
    }
    
    for _, tc := range testCases {
        result := IsLeapYear(tc.year)
        if result != tc.expected {
            t.Errorf("IsLeapYear(%d) = %v; want %v", tc.year, result, tc.expected)
        }
    }
}
```

## 🎯 Performance

The function has:
- **Time complexity:** O(1) - Constant time
- **Space complexity:** O(1) - No additional memory needed
- **Operations:** Only 3 modulo operations maximum

## 🔧 Edge Cases Handled

1. **Negative years**: Works correctly (though historically inaccurate)
2. **Year 0**: Handled according to mathematical rules
3. **Large years**: No overflow issues (Go's `int` handles large values)
4. **Year 1**: Correctly returns false

```go
fmt.Println(IsLeapYear(0))     // true (mathematically)
fmt.Println(IsLeapYear(1))     // false
fmt.Println(IsLeapYear(-4))    // true (mathematically)
fmt.Println(IsLeapYear(10000)) // false
```

## 📈 Real-World Applications

1. **Date validation** in form inputs
2. **Calendar applications** displaying correct February days
3. **Scheduling systems** for annual events
4. **Financial applications** for interest calculations
5. **Age calculation** systems

## 🌍 International Considerations

This implements the **Gregorian calendar** rules used by most countries. Note that:
- Some countries adopted the Gregorian calendar at different times
- Historical dates before adoption may be inaccurate
- Alternative calendars (Jewish, Islamic, Chinese) have different leap rules

## 🚨 Common Misconceptions

❌ **"Every 4 years is a leap year"** → Wrong! Century years must also be divisible by 400.

❌ **"Divisible by 4 means leap year"** → Wrong! 1900 was divisible by 4 but not a leap year.

✅ **Correct rule**: Use the complete logical condition in this package.

## 📝 Example: Complete Program

```go
package main

import (
    "fmt"
    "time"
)

func IsLeapYear(year int) bool {
    if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
        return true
    }
    return false
}

func main() {
    currentYear := time.Now().Year()
    
    fmt.Printf("Current year: %d\n", currentYear)
    
    if IsLeapYear(currentYear) {
        fmt.Println("This is a leap year! 🎉")
        fmt.Println("February has 29 days.")
    } else {
        fmt.Println("This is not a leap year.")
        fmt.Println("February has 28 days.")
        
        // Find next leap year
        nextYear := currentYear + 1
        for !IsLeapYear(nextYear) {
            nextYear++
        }
        fmt.Printf("Next leap year: %d\n", nextYear)
    }
}
```

## 🤝 Contributing

Feel free to:
- Add more test cases
- Create benchmarking tests
- Extend with calendar-specific functions
- Add support for other calendar systems

## 📄 License

This code is provided as-is. Feel free to use, modify, and distribute it according to your needs.

---

**Happy coding!** May all your date calculations be accurate! 📅✅
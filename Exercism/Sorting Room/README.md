# Type Description Utility

A Go program that demonstrates type assertions, interfaces, and polymorphism through various description functions.

## 📋 Overview

This program provides a set of functions to describe different types of numbers and boxes containing numbers, showcasing Go's interface system and type handling capabilities.

## 🏗️ Architecture

The program implements two main interface hierarchies:

### 1. **NumberBox System**
- `NumberBox` interface: Requires `Number() int` method
- `numberBoxContaining` struct: Concrete implementation

### 2. **FancyNumberBox System**
- `FancyNumberBox` interface: Requires `Value() string` method
- `FancyNumber` struct: Implementation that can be extracted
- `AnotherFancyNumber` struct: Implementation that returns 0 when extracted

## 📖 API Documentation

### Core Functions

#### `DescribeNumber(f float64) string`
Returns a formatted description of a floating-point number.

```go
DescribeNumber(12.345) // "This is the number 12.3"
```

#### `DescribeNumberBox(nb NumberBox) string`
Returns a description of any type implementing the `NumberBox` interface.

```go
box := numberBoxContaining{12}
DescribeNumberBox(box) // "This is a box containing the number 12.0"
```

#### `ExtractFancyNumber(fnb FancyNumberBox) int`
Extracts the integer value from a `FancyNumber` or returns 0 for other types.

```go
ExtractFancyNumber(FancyNumber{"10"})       // 10
ExtractFancyNumber(AnotherFancyNumber{"4"}) // 0
```

#### `DescribeFancyNumberBox(fnb FancyNumberBox) string`
Describes fancy number boxes, showing special handling for `FancyNumber` types.

```go
DescribeFancyNumberBox(FancyNumber{"10"})       // "This is a fancy box containing the number 10.0"
DescribeFancyNumberBox(AnotherFancyNumber{"4"}) // "This is a fancy box containing the number 0.0"
```

#### `DescribeAnything(i any) string`
Universal function that accepts any type and routes to the appropriate description function.

```go
DescribeAnything(42)                // "This is the number 42.0"
DescribeAnything(3.14)              // "This is the number 3.1"
DescribeAnything("unknown")         // "Return to sender"
```

## 🚀 Quick Start

```go
program main

import (
    "fmt"
)

func main() {
    // Basic number description
    fmt.Println(DescribeNumber(-12.345))
    // Output: This is the number -12.3
    
    // Number box description
    box := numberBoxContaining{12}
    fmt.Println(DescribeNumberBox(box))
    // Output: This is a box containing the number 12.0
    
    // Fancy number operations
    fmt.Println(ExtractFancyNumber(FancyNumber{"10"}))
    // Output: 10
    
    fmt.Println(DescribeFancyNumberBox(AnotherFancyNumber{"4"}))
    // Output: This is a fancy box containing the number 0.0
    
    // Universal descriptor
    fmt.Println(DescribeAnything(42))
    // Output: This is the number 42.0
    
    fmt.Println(DescribeAnything("test"))
    // Output: Return to sender
}
```

## 🧪 Examples

### Example 1: Basic Usage
```go
// Describe various number types
fmt.Println(DescribeNumber(123.456))      // This is the number 123.5
fmt.Println(DescribeNumber(-5.0))         // This is the number -5.0

// Work with boxes
myBox := numberBoxContaining{7}
fmt.Println(DescribeNumberBox(myBox))     // This is a box containing the number 7.0
```

### Example 2: Type Switching
```go
// The DescribeAnything function demonstrates type switching
items := []any{
    42,
    3.14159,
    numberBoxContaining{100},
    FancyNumber{"25"},
    AnotherFancyNumber{"50"},
    "not a number",
    true,
}

for _, item := range items {
    fmt.Println(DescribeAnything(item))
}
```

### Example 3: Error Handling
```go
// FancyNumber with non-numeric value
fn := FancyNumber{"not-a-number"}
fmt.Println(ExtractFancyNumber(fn))       // 0 (due to conversion error)
fmt.Println(DescribeFancyNumberBox(fn))   // "An Error occured while extracting number"
```

## ⚙️ Implementation Details

### Type Assertions
The program demonstrates safe type assertions:

```go
// In ExtractFancyNumber
if _, ok := fnb.(FancyNumber); ok {
    // Only process FancyNumber types
}

// In DescribeAnything
switch i.(type) {
case int:
    // Handle int
case float64:
    // Handle float64
// ... other cases
}
```

### Interface Design
- **NumberBox**: Simple interface with one method
- **FancyNumberBox**: More complex with string-based values
- Both demonstrate different approaches to polymorphism

## 🐛 Error Handling

The code handles potential errors:
- String to integer conversion errors in `ExtractFancyNumber`
- Unknown types in `DescribeAnything` (returns "Return to sender")
- Non-numeric values in fancy boxes

## 📊 Output Formatting

All description functions format numbers with one decimal place using Go's `fmt.Sprintf`:
- `%.1f` format specifier for consistent single decimal precision
- Clear, human-readable descriptions

## 🎯 Use Cases

1. **Educational Tool**: Learn about Go interfaces and type assertions
2. **Debugging Utility**: Describe various data types in a uniform way
3. **API Response Formatting**: Format different numerical data types consistently
4. **Polymorphism Example**: Demonstrate interface-based design patterns

## 🔧 Extending the Package

To add support for new types:

1. Implement the appropriate interface (`NumberBox` or `FancyNumberBox`)
2. Add a new case in `DescribeAnything`'s type switch
3. Create a specific description function if needed

## 📝 Notes

- The program uses `any` (Go 1.18+) for the generic `DescribeAnything` function
- Type assertions are performed safely with the comma-ok idiom
- String conversions handle potential errors gracefully
- All floating-point numbers are formatted to one decimal place

## 🤝 Contributing

This is primarily an educational example. Feel free to:
- Add more number container types
- Implement additional interfaces
- Extend the type switching logic
- Improve error messages and handling
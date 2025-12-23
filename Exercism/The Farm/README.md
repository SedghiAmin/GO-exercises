# Farm Fodder Distribution System

A Go implementation of a fodder calculation system for farm management that ensures equal food distribution among cows while handling various edge cases and errors.

## Overview

This system calculates the appropriate amount of fodder for each cow on a farm, ensuring all cows receive equal portions to maintain herd harmony. It features robust error handling, input validation, and a clean interface-based architecture.

## Features

- **Equal Distribution**: Ensures each cow receives the same amount of food
- **Input Validation**: Validates cow counts with descriptive error messages
- **Error Handling**: Comprehensive error management with custom error types
- **Interface-Based**: Extensible design using Go interfaces
- **Testable**: Clean separation of concerns for easy testing

## Installation

```bash
# Clone the repository (if applicable)
git clone <repository-url>
cd fodder-calculator

# Run the program
go run main.go
```

## Usage

### Basic Calculation

```go
package main

func main() {
    food := Food{}
    result, err := ValidateInputAndDivideFood(food, 5)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Each cow gets: %.2f units of food\n", result)
    }
}
```

### Direct Calculation (Without Validation)

```go
result, err := DivideFood(food, 5)
if err != nil {
    fmt.Println("Error:", err)
}
```

### Individual Validation

```go
err := ValidateNumberOfCows(10)
if err != nil {
    fmt.Println("Validation error:", err)
}
```

## API Reference

### Interfaces

#### `FodderCalculator`
```go
type FodderCalculator interface {
    FodderAmount(int) (float64, error)
    FatteningFactor() (float64, error)
}
```

### Structs

#### `Food`
A sample implementation of `FodderCalculator` for demonstration purposes.
- `FodderAmount(int)`: Returns 50.0 (example value)
- `FatteningFactor()`: Returns 1.5 (example factor)

#### `InvalidCowsError`
Custom error type for invalid cow counts.
```go
type InvalidCowsError struct {
    cows int
    msg  string
}
```

### Functions

#### `DivideFood(fc FodderCalculator, cows int) (float64, error)`
Calculates the amount of food per cow.

**Formula:**
```
foodPerCow = (FodderAmount(cows) × FatteningFactor()) / cows
```

**Parameters:**
- `fc`: Implementation of `FodderCalculator`
- `cows`: Number of cows (assumed to be positive)

**Returns:**
- `float64`: Amount of food per cow
- `error`: Any error from the `FodderCalculator` methods

**Example:**
```go
result, err := DivideFood(food, 5)
// Returns: 15.0, nil (50 × 1.5 / 5 = 15)
```

#### `ValidateNumberOfCows(cows int) error`
Validates the number of cows and returns an appropriate error if invalid.

**Validation Rules:**
- `cows < 0`: Returns `InvalidCowsError` with message "there are no negative cows"
- `cows == 0`: Returns `InvalidCowsError` with message "no cows don't need food"
- `cows > 0`: Returns `nil`

**Error Format:**
```
{cows} cows are invalid: {message}
```

**Example:**
```go
err := ValidateNumberOfCows(-3)
// Error: "-3 cows are invalid: there are no negative cows"
```

#### `ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error)`
Validates the input and then calculates the food per cow.

**Workflow:**
1. Validate cow count using `ValidateNumberOfCows`
2. If valid, call `DivideFood`
3. Return result or first encountered error

**Example:**
```go
result, err := ValidateInputAndDivideFood(food, 5)
// Returns: 15.0, nil

result, err := ValidateInputAndDivideFood(food, -2)
// Returns: 0.0, "-2 cows are invalid: there are no negative cows"
```

## Error Handling

The system provides detailed error messages:

### Invalid Cow Counts
```
-5 cows are invalid: there are no negative cows
0 cows are invalid: no cows don't need food
```

### External Calculator Errors
Errors from the `FodderCalculator` implementation are propagated unchanged.

## Examples

### Successful Calculation
```go
food := Food{}
result, err := ValidateInputAndDivideFood(food, 5)
// result = 15.0, err = nil
```

### Invalid Input (Negative Cows)
```go
result, err := ValidateInputAndDivideFood(food, -2)
// result = 0.0, err = "-2 cows are invalid: there are no negative cows"
```

### Invalid Input (Zero Cows)
```go
result, err := ValidateInputAndDivideFood(food, 0)
// result = 0.0, err = "0 cows are invalid: no cows don't need food"
```

## Extending the System

### Adding a New Fodder Calculator
```go
type CustomCalculator struct{}

func (c CustomCalculator) FodderAmount(cows int) (float64, error) {
    // Custom logic here
    return 75.0, nil
}

func (c CustomCalculator) FatteningFactor() (float64, error) {
    // Custom logic here
    return 1.2, nil
}

// Usage
calculator := CustomCalculator{}
result, err := ValidateInputAndDivideFood(calculator, 10)
```

### Modifying Error Messages
Edit the messages in `ValidateNumberOfCows`:
```go
if cows < 0 {
    return &InvalidCowsError{cows: cows, msg: "Cows cannot be negative"}
}
```

## Testing

Run the program:
```bash
go run main.go
```

Expected output for the provided main function:
```
15
```

## Design Principles

1. **Single Responsibility**: Each function has a clear, specific purpose
2. **Interface Segregation**: Small, focused interfaces
3. **Error Transparency**: Errors are propagated with context
4. **Validation First**: Input validation before processing
5. **Testability**: Functions are pure and easy to test

## Assumptions

1. The `FodderCalculator` implementation (`Food` struct) provides example values
2. Real implementations would calculate values based on cow breed, age, weight, etc.
3. The system assumes cows are the only animals being fed
4. All cows receive equal portions regardless of individual needs

## Performance Considerations

- Minimal memory footprint
- Fast validation with early exit on errors
- No external dependencies
- Efficient floating-point calculations

## Security Considerations

- Input validation prevents division by zero
- Type safety through Go's strong typing system
- No external data sources in this example implementation

## Future Enhancements

1. **Multiple Animal Types**: Support for different animal types
2. **Individual Needs**: Account for cow age, weight, and breed
3. **Seasonal Factors**: Adjust calculations based on season
4. **Database Integration**: Store and retrieve historical data
5. **Configuration**: External configuration for calculation parameters
6. **API Endpoints**: REST API for remote access
7. **Logging**: Structured logging for monitoring
8. **Metrics**: Performance and usage metrics

## Contributing

1. Fork the repository
2. Create a feature branch
3. Implement changes with tests
4. Submit a pull request

## License

This code is provided as an example. Modify and use according to your needs.

## Support

For questions or issues:
1. Review the code comments
2. Check example usage
3. Contact the maintainer


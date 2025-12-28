# Resident Information System

A Go package for managing resident information in a city, with validation and data management capabilities.

## Overview

This program provides a `Resident` struct and associated methods to handle resident information, validate required data, and perform cleanup operations.

## Features

- Create new resident records
- Validate required resident information
- Delete/reset resident data
- Count residents with complete information

## Usage

### Creating a Resident

```go
name := "Matthew Sanabria"
age := 29
address := map[string]string{"street": "Main St."}

resident := NewResident(name, age, address)
// &{Matthew Sanabria 29 map[street:Main St.]}
```

### Checking Required Information

```go
// A resident has required information if:
// 1. Name is not empty
// 2. Address contains a non-empty "street" key

resident.HasRequiredInfo() // returns true or false
```

### Deleting Resident Information

```go
resident.Delete()
// Resets all fields to their zero values:
// Name: "", Age: 0, Address: nil
```

### Counting Valid Residents

```go
residents := []*Resident{resident1, resident2, resident3}
validCount := Count(residents)
// Returns the number of residents with required information
```

## Examples

### Example 1: Basic Usage
```go
name := "Matthew Sanabria"
age := 29
address := map[string]string{"street": "Main St."}

resident1 := NewResident(name, age, address)
fmt.Println(resident1.HasRequiredInfo()) // true
```

### Example 2: Missing Information
```go
resident2 := NewResident("", 30, map[string]string{"street": "Main St."})
fmt.Println(resident2.HasRequiredInfo()) // false (name missing)

resident3 := NewResident("Alice", 25, map[string]string{})
fmt.Println(resident3.HasRequiredInfo()) // false (street missing)

resident4 := NewResident("Bob", 35, map[string]string{"street": ""})
fmt.Println(resident4.HasRequiredInfo()) // false (street empty)
```

### Example 3: Delete Operation
```go
resident := NewResident("John", 40, map[string]string{"street": "Elm St."})
resident.Delete()
fmt.Println(resident) // &{ 0 map[]}
```

### Example 4: Counting Residents
```go
r1 := NewResident("Alice", 25, map[string]string{"street": "Oak St."})
r2 := NewResident("", 30, map[string]string{"street": "Pine St."})
r3 := NewResident("Bob", 35, map[string]string{})

residents := []*Resident{r1, r2, r3}
fmt.Println(Count(residents)) // 1 (only Alice has all required info)
```

## API Reference

### Types

#### `Resident`
```go
type Resident struct {
    Name    string
    Age     int
    Address map[string]string
}
```

### Functions

#### `NewResident(name string, age int, address map[string]string) *Resident`
Creates a new resident instance with the provided information.

#### `Count(residents []*Resident) int`
Returns the number of residents who have provided all required information.

### Methods

#### `(r *Resident) HasRequiredInfo() bool`
Determines if a resident has all required information:
- Name must not be empty
- Address must contain a "street" key with a non-empty value

#### `(r *Resident) Delete()`
Deletes all resident information by resetting fields to their zero values.

## Requirements

- Go 1.16 or higher
- No external dependencies

## Testing

Run the test suite:
```bash
go test -v
```

## Notes

- The `Address` field is a map that should contain at minimum a "street" key
- Age is not required for validation (only Name and street address)
- `Delete()` method completely resets the resident object
- Nil maps are acceptable as empty addresses

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for your changes
4. Ensure all tests pass
5. Submit a pull request

## License

This project is part of Exercism.io Go track exercises.
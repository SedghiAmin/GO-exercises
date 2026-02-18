# Phone Number Cleaner and Formatter in Go

This program provides functions to clean, validate, and format North American-style phone numbers according to standard conventions.

## Features

- **Clean phone numbers** by removing punctuation, spaces, and country codes
- **Validate** phone numbers for correct length and format
- **Extract area codes** from valid phone numbers
- **Format** phone numbers in the standard `(NXX) NXX-XXXX` format

## Usage

### Basic Functions

```go
package main

import (
    "fmt"
    "yourmodule/phonenumber"
)

func main() {
    // Clean a phone number
    num, err := phonenumber.Number("(223) 456-7890")
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(num) // Output: 2234567890
    
    // Extract area code
    areaCode, _ := phonenumber.AreaCode("+1 (223) 456-7890")
    fmt.Println(areaCode) // Output: 223
    
    // Format a phone number
    formatted, _ := phonenumber.Format("2234567890")
    fmt.Println(formatted) // Output: (223) 456-7890
}
```

## Function Documentation

### `Number(phoneNumber string) (string, error)`

Cleans and validates a phone number according to North American standards.

**Rules:**
- Removes punctuation: `(`, `)`, `-`, `.`, and spaces
- Handles optional `+1` country code prefix
- Validates final number must be 10 digits
- Area code (first 3 digits) cannot start with 0 or 1
- Exchange code (digits 4-6) cannot start with 0 or 1
- 11-digit numbers are valid only if they start with 1 (US country code)

**Examples:**
```go
Number("(223) 456-7890")     // Returns: "2234567890", nil
Number("+1 (223) 456-7890")  // Returns: "2234567890", nil
Number("123-456-7890")       // Returns: "", error (invalid area code)
Number("223-056-7890")       // Returns: "", error (invalid exchange code)
```

### `AreaCode(phoneNumber string) (string, error)`

Extracts the 3-digit area code from a valid phone number.

**Examples:**
```go
AreaCode("(223) 456-7890")   // Returns: "223", nil
AreaCode("+1 223-456-7890")  // Returns: "223", nil
```

### `Format(phoneNumber string) (string, error)`

Formats a valid phone number into standard `(NXX) NXX-XXXX` format.

**Examples:**
```go
Format("2234567890")          // Returns: "(223) 456-7890", nil
Format("+1 (223) 456-7890")   // Returns: "(223) 456-7890", nil
```

## Complete Example

The `main()` function demonstrates testing various phone number formats:

```go
func main() {
    test := []string{
        "(223) 456-7890",
        "223.456.7890",
        "223 456   7890   ",
        "123456789",
        "22234567890",
        "12234567890",
        "+1 (223) 456-7890",
        "321234567890",
        "523-abc-7890",
        "523-@:!-7890",
        "(023) 456-7890",
        "(123) 456-7890",
        "(223) 056-7890",
        "(223) 156-7890",
        "1 (023) 456-7890",
        "1 (123) 456-7890",
        "1 (223) 056-7890",
        "1 (223) 156-7890",
    }
    
    for i := range test {
        num, err := Number(test[i])
        area, _ := AreaCode(test[i])
        format, _ := Format(test[i])
        fmt.Printf(
            "input: %v , output => value: %v , AreaCode: %v , Formatted: %v , error: %v \n",
            test[i],
            num,
            area,
            format,
            err,
        )
    }
}
```

### Expected Output

Running the example will produce output similar to:

```
input: (223) 456-7890 , output => value: 2234567890 , AreaCode: 223 , Formatted: (223) 456-7890 , error: <nil> 
input: 223.456.7890 , output => value: 2234567890 , AreaCode: 223 , Formatted: (223) 456-7890 , error: <nil> 
input: 223 456   7890    , output => value: 2234567890 , AreaCode: 223 , Formatted: (223) 456-7890 , error: <nil> 
input: 123456789 , output => value:  , AreaCode:  , Formatted:  , error: invalid count of digits digits 
input: 22234567890 , output => value: 22234567890 , AreaCode: 222 , Formatted: (222) 345-6789 , error: <nil> 
input: 12234567890 , output => value: 2234567890 , AreaCode: 223 , Formatted: (223) 456-7890 , error: <nil> 
input: +1 (223) 456-7890 , output => value: 2234567890 , AreaCode: 223 , Formatted: (223) 456-7890 , error: <nil> 
input: 321234567890 , output => value:  , AreaCode:  , Formatted:  , error: invalid count of digits digits 
input: 523-abc-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid with punctuations 
input: 523-@:!-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid with punctuations 
input: (023) 456-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid area code 
input: (123) 456-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid area code 
input: (223) 056-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid exchange code 
input: (223) 156-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid exchange code 
input: 1 (023) 456-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid area code 
input: 1 (123) 456-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid area code 
input: 1 (223) 056-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid exchange code 
input: 1 (223) 156-7890 , output => value:  , AreaCode:  , Formatted:  , error: invalid exchange code 
```

## Validation Rules

The package enforces the following North American phone number conventions:

1. **Length**: Exactly 10 digits after cleaning
2. **Country Code**: Optional `+1` or `1` prefix (removed automatically)
3. **Area Code**: First digit cannot be 0 or 1 (NXX format)
4. **Exchange Code**: Fourth digit cannot be 0 or 1 (NXX format)
5. **Allowed Characters**: Digits, spaces, parentheses, hyphens, dots, and optional leading `+`

## Error Types

The package returns descriptive errors for various invalid inputs:

- `"invalid country code"` - When + is followed by something other than 1
- `"invalid with punctuations"` - When unsupported characters are found
- `"invalid count of digits digits"` - When cleaned number isn't 10 digits
- `"invalid when 11 digits does not start with a 1"` - Invalid 11-digit format
- `"invalid area code"` - Area code starts with 0 or 1
- `"invalid exchange code"` - Exchange code starts with 0 or 1

## Performance Considerations

- Uses `[]rune` for proper Unicode handling
- Pre-allocates slice capacity for efficiency
- Simple character-by-character parsing for clarity

## Testing

Run the tests with:

```bash
go test -v
```

Example test cases cover:
- Valid numbers with various formats
- Numbers with country codes
- Invalid lengths
- Invalid characters
- Invalid area codes
- Invalid exchange codes

## Contributing

Contributions are welcome! Please ensure:
- All tests pass
- Code follows Go standards (`go fmt`)
- New features include appropriate tests

## License

This code is provided as an educational example and is free to use, modify, and distribute.
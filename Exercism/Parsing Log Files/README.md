# Log File Parser

A Go program for parsing and processing log files with various validation, extraction, and transformation capabilities.

## Features

### 1. **Line Validation**
- Validates if a log line starts with a standard log level tag
- Supported tags: `[TRC]`, `[DBG]`, `[INF]`, `[WRN]`, `[ERR]`, `[FTL]`
- Lines must begin with one of these tags to be considered valid

### 2. **Log Line Splitting**
- Splits log lines using custom field separators
- Separators are strings that start with `<`, end with `>`, and contain any combination of `~`, `*`, `=`, `-` in between
- Examples: `<*>`, `<~~~>`, `<--->`, `<=>`

### 3. **Password Detection**
- Detects quoted passwords in log lines
- Case-insensitive matching for the word "password"
- Only counts passwords surrounded by quotation marks
- Accounts for additional content between quotation marks

### 4. **End-of-Line Text Removal**
- Removes "end-of-line" markers followed by numbers from log lines
- Useful for cleaning up log files with debugging markers

### 5. **User Tagging**
- Identifies lines containing user references
- Extracts usernames following "User" with one or more spaces
- Tags such lines with `[USR]` prefix followed by the username

## Installation

```bash
go get your-module-path
```

## Usage

```go
package main

import (
    "fmt"
    "your-module-path/parsinglogfiles"
)

func main() {
    // Validate log lines
    fmt.Println(parsinglogfiles.IsValidLine("[ERR] A good error here")) // true
    fmt.Println(parsinglogfiles.IsValidLine("Any old [ERR] text"))      // false
    
    // Split log lines with custom separators
    result := parsinglogfiles.SplitLogLine("section 1<*>section 2<~~~>section 3")
    // []string{"section 1", "section 2", "section 3"}
    
    // Count quoted passwords
    lines := []string{
        `"passWord"`,
        `[INF] The message "Please reset your password" was ignored`,
    }
    count := parsinglogfiles.CountQuotedPasswords(lines) // 2
    
    // Remove end-of-line markers
    cleaned := parsinglogfiles.RemoveEndOfLineText(
        "[INF] end-of-line23033 Network Failure end-of-line27"
    ) // "[INF]  Network Failure "
    
    // Tag lines with usernames
    tagged := parsinglogfiles.TagWithUserName([]string{
        "[WRN] User James123 has exceeded storage space.",
        "[INF] Users can login again after 23:00.",
    })
    // []string{
    //     "[USR] James123 [WRN] User James123 has exceeded storage space.",
    //     "[INF] Users can login again after 23:00.",
    // }
}
```

## API Reference

### `IsValidLine(text string) bool`
Returns `true` if the line starts with a valid log level tag, `false` otherwise.

### `SplitLogLine(text string) []string`
Splits the input text using custom separators and returns the fields as a slice of strings.

### `CountQuotedPasswords(lines []string) int`
Returns the number of lines containing the word "password" (case-insensitive) surrounded by quotation marks.

### `RemoveEndOfLineText(text string) string`
Removes all occurrences of "end-of-line" followed by digits from the input string.

### `TagWithUserName(lines []string) []string`
Processes log lines, tagging those containing user references with `[USR]` prefix.

## Regular Expressions Used

1. **Line Validation**: `^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*`
2. **Line Splitting**: `<[*-=~]*>`
3. **Password Detection**: `(?i)"[^"]*?password[^"]*?"`
4. **End-of-Line Removal**: `end-of-line\d+`
5. **User Extraction**: `(User\s+)([a-zA-Z0-9_]+)`

## Examples

### Valid Log Lines
```go
IsValidLine("[DBG] Debug information")      // true
IsValidLine("[ERR] System error occurred")  // true
IsValidLine("[XYZ] Invalid tag")           // false
```

### Field Splitting
```go
SplitLogLine("a<->b<~~>c")        // []string{"a", "b", "c"}
SplitLogLine("test<>empty")       // []string{"test", "empty"}
SplitLogLine("no separators")     // []string{"no separators"}
```

### Password Detection
```go
CountQuotedPasswords([]string{
    `"password"`,                            // counted
    `"new PASSWORD required"`,               // counted
    `password without quotes`,               // not counted
    `"some text" and then password`,         // not counted
    `The admin said "change password" now`,  // counted
}) // returns 2
```

## Dependencies

- Go 1.16 or higher
- Standard library only (no external dependencies)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Add tests for your changes
4. Make your changes
5. Ensure all tests pass
6. Submit a pull request


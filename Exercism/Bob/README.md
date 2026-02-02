# Hey - Bob's Automated Response System

A simple Go project that provides appropriate responses based on how the user is speaking.

## 📋 Overview

This project implements a `Hey` function that takes a string (remark or question) as input and returns an appropriate response based on its content. This is inspired by a classic programming exercise and demonstrates basic string manipulation and conditional logic in Go.

## 🚀 Features

- Detects regular questions
- Detects shouting (all uppercase sentences)
- Detects shouted questions
- Handles empty input gracefully
- Properly processes strings with numbers and special characters
- Unicode-aware character processing

## 📁 Project Structure

```
/
├── main.go          # Main implementation with test cases
└── README.md        # This documentation file
```

## 🛠️ Installation & Usage

### Prerequisites
- Go 1.16 or higher installed

### Running the Program

```bash
# Clone or create the project
git clone <repository-url>
cd <project-directory>

# Run the program
go run main.go
```

## 📝 Code Implementation

### Core Function: `Hey(remark string) string`

The main function that processes input and returns responses:

```go
func Hey(remark string) string {
    remark = strings.TrimSpace(remark)
    input := []rune(remark)
    
    if remark == "" {
        return "Fine. Be that way!"
    } else if remark == strings.ToUpper(remark) && input[len(input)-1] == '?' && haveletter(remark) {
        return "Calm down, I know what I'm doing!"
    } else if remark == strings.ToUpper(remark) && haveletter(remark) {
        return "Whoa, chill out!"
    } else if input[len(input)-1] == '?' {
        return "Sure."
    }
    return "Whatever."
}
```

### Helper Function: `haveletter(remark string) bool`

Checks if a string contains at least one letter:

```go
func haveletter(remark string) bool {
    input := []rune(remark)
    for _, char := range input {
        if unicode.IsLetter(char) {
            return true
        }
    }
    return false
}
```

## 📊 Response Logic

| Input Condition | Response | Example |
|----------------|----------|---------|
| Empty string or whitespace only | `"Fine. Be that way!"` | `""`, `"   "`, `"\t\n"` |
| Question ending with '?' | `"Sure."` | `"How are you?"`, `"Is this a test?"` |
| Shouting (all uppercase letters) | `"Whoa, chill out!"` | `"HELLO"`, `"STOP"` |
| Shouted question | `"Calm down, I know what I'm doing!"` | `"WHAT ARE YOU DOING?"` |
| Numbers/symbols only (no letters) | `"Whatever."` | `"1, 2, 3"`, `"%@#$"` |
| Mixed case statements | `"Whatever."` | `"Hello world"` |

## 🧪 Test Cases

The included test cases demonstrate various scenarios:

```go
testCases := []string{
    "",                     // Empty - Fine. Be that way!
    "   ",                  // Whitespace - Fine. Be that way!
    "How are you?",         // Regular question - Sure.
    "HELLO",                // Shouting - Whoa, chill out!
    "WHAT ARE YOU DOING?",  // Shouted question - Calm down...
    "1, 2, 3",              // Numbers only - Whatever.
    "1, 2, 3?",             // Number question - Sure.
    "HELLO 123?",           // Shouted numbers question - Calm down...
    "hello?",               // Lowercase question - Sure.
}
```

## 🎯 Sample Output

```
Quick tests for Hey():
----------------------------------------
Hey("") = "Fine. Be that way!"
Hey("   ") = "Fine. Be that way!"
Hey("How are you?") = "Sure."
Hey("HELLO") = "Whoa, chill out!"
Hey("WHAT ARE YOU DOING?") = "Calm down, I know what I'm doing!"
Hey("1, 2, 3") = "Whatever."
Hey("1, 2, 3?") = "Sure."
Hey("HELLO 123?") = "Calm down, I know what I'm doing!"
Hey("hello?") = "Sure."
```

## 🔧 Key Implementation Details

1. **String Trimming**: Uses `strings.TrimSpace()` to handle leading/trailing whitespace
2. **Rune Conversion**: Converts strings to `[]rune` for proper Unicode character handling
3. **Letter Detection**: Uses `unicode.IsLetter()` to identify alphabetic characters
4. **Case Detection**: Uses `strings.ToUpper()` comparison to detect shouting
5. **Question Detection**: Checks if the last character is '?'

## 📚 Learning Points

This project demonstrates:
- String manipulation in Go
- Conditional logic with multiple conditions
- Unicode-aware string processing
- Helper function design
- Test-driven development approach

## 🚀 Getting Started for Development

1. Fork or clone the repository
2. Modify the `Hey` function to add new features
3. Add more test cases to the `main` function
4. Run tests: `go run main.go`
5. Consider adding unit tests in a separate `_test.go` file

## 🤝 Contributing

Feel free to:
- Add more test cases
- Improve the response logic
- Add support for more edge cases
- Implement additional language features
- Create proper unit tests

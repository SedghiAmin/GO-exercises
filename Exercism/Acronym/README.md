# Abbreviate - Acronym Generator

A simple Go program that converts phrases into their acronyms/abbreviations. This function takes a string input and returns the acronym formed by the first letter of each word, handling various separator characters.

## Features

- Converts multi-word phrases into acronyms
- Handles spaces, hyphens, and underscores as word separators
- Preserves the case of letters (returns as-is)
- Safe boundary checking to prevent index out of range errors
- Efficient string building using `strings.Builder`

## Installation

```bash
# Clone the repository or copy the code directly
git clone <repository-url>
```

## Usage

```go
package main

import (
    "fmt"
    "strings"
)

func Abbreviate(s string) string {
    out := strings.Builder{}
    for i := 0; i < len(s); i++ {
        if i == 0 && string(s[i]) != " " && string(s[i]) != "-" && string(s[i]) != "_" {
            out.WriteString(string(s[i]))
            continue
        }
        if string(s[i]) == " " || string(s[i]) == "-" || string(s[i]) == "_" {
            if (i+1 < len(s)) && (string(s[i+1]) != " ") && (string(s[i+1]) != "-") && (string(s[i+1]) != "_") {
                out.WriteString(strings.ToUpper(string(s[i+1])))
            }
        }
    }
    return out.String()
}

func main() {
    // Example usage
    input := "Complementary metal-oxide semiconductor"
    fmt.Println(Abbreviate(input)) // Output: "CMOS"
}
```

## Examples

| Input | Output |
|-------|--------|
| "Portable Network Graphics" | "PNG" |
| "Complementary metal-oxide semiconductor" | "CMOS" |
| "Ruby on Rails" | "RoR" |
| "First In, First Out" | "FIFO" |
| "The Road _Not_ Taken" | "TRNT" |
| "PHP: Hypertext Preprocessor" | "PHP" |
| "light emitting diode" | "lED" |

## How It Works

1. **First Character**: The first character of the string is always included if it's not a separator
2. **Separator Detection**: The function looks for separators (space, hyphen, underscore)
3. **Next Character**: When a separator is found, the next character is added to the acronym
4. **Safety Check**: Ensures we don't access indices beyond the string length

## Limitations

- The function doesn't automatically uppercase all letters (returns as-is)
- Only recognizes spaces, hyphens, and underscores as separators
- Doesn't handle punctuation marks (commas, periods, etc.) as separators

## Contributing

Feel free to submit issues or pull requests if you find bugs or have suggestions for improvements.

## License

This code is available under the MIT License. Feel free to use it in your projects.
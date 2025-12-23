# Airport Robot Greeter

## Description
A Go implementation of a multilingual airport greeting robot that can welcome visitors in different languages using Go interfaces.

## Features
- **Multilingual Support**: Greet visitors in German, Italian, and Portuguese
- **Extensible Design**: Easy to add new languages by implementing the `Greeter` interface
- **Clean Interface**: Simple and intuitive API for greeting functionality

## Code Structure

### Interface Definition
```go
type Greeter interface {
    LanguageName() string
    Greet(string) string
}
```

### Core Function
```go
func SayHello(name string, greeter Greeter) string
```
Generates a formatted greeting message combining the language name and personalized greeting.

### Available Languages

#### German
- Language: German
- Greeting: "Hallo {name}!"
- Implementation: `German` struct

#### Italian
- Language: Italian
- Greeting: "Ciao {name}!"
- Implementation: `Italian` struct

#### Portuguese
- Language: Portuguese
- Greeting: "Olá {name}!"
- Implementation: `Portuguese` struct

## Usage Example

```go
package main

import (
    "fmt"
)

func main() {
    // Create greeters for different languages
    germanGreeter := German{}
    italianGreeter := Italian{}
    portugueseGreeter := Portuguese{}
    
    // Greet visitors
    fmt.Println(SayHello("Dietrich", germanGreeter))
    fmt.Println(SayHello("Giovanni", italianGreeter))
    fmt.Println(SayHello("Carlos", portugueseGreeter))
}
```

**Output:**
```
I can speak German: Hallo Dietrich!
I can speak Italian: Ciao Giovanni!
I can speak Portuguese: Olá Carlos!
```

## How to Add a New Language

1. Create a new struct for the language:
```go
type French struct{}
```

2. Implement the `Greeter` interface:
```go
func (f French) LanguageName() string {
    return "French"
}

func (f French) Greet(name string) string {
    return "Bonjour " + name + "!"
}
```

3. Use it with `SayHello`:
```go
frenchGreeter := French{}
fmt.Println(SayHello("Pierre", frenchGreeter))
```

## Design Principles

1. **Open/Closed Principle**: Open for extension (new languages), closed for modification
2. **Interface Segregation**: Small, focused interface with clear responsibilities
3. **Dependency Inversion**: High-level `SayHello` function depends on abstraction (`Greeter` interface), not concrete implementations

## Testing
Run the example program to see the robot in action:
```bash
go run main.go
```

## Future Enhancements
- Add more languages (Spanish, Japanese, Arabic, etc.)
- Include localization options (formal/informal greetings)
- Add time-based greetings (Good morning/afternoon/evening)

## License
This project is part of an Exercism Go exercise and follows their licensing terms.
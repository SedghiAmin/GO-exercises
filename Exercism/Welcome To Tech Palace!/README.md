# Tech Palace Marketing System

A Go program for generating and managing marketing messages for Tech Palace, a technology retail store. This package provides utilities for creating welcome messages, adding decorative borders, and cleaning up old marketing content.

## Overview

The Tech Palace Marketing System is designed to help create consistent and formatted marketing messages. It includes three main functions for message manipulation that follow Tech Palace's branding guidelines.

## Features

- Generate personalized welcome messages
- Add decorative borders to messages
- Clean up old marketing messages by removing formatting
- Simple and intuitive API

## Usage

### Welcome Messages

Create personalized welcome messages for customers:

```go
customer := "Judy"
message := WelcomeMessage(customer)
// Returns: "Welcome to the Tech Palace, JUDY"
```

### Adding Borders

Add decorative star borders to messages:

```go
welcome := "Welcome!"
bordered := AddBorder(welcome, 10)
// Returns:
// **********
// Welcome!
// **********
```

### Cleaning Messages

Clean up old marketing messages by removing stars and extra whitespace:

```go
oldMessage := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

cleaned := CleanupMessage(oldMessage)
// Returns: "BUY NOW, SAVE 10%"
```

## API Reference

### `WelcomeMessage(customer string) string`

Creates a welcome message for the specified customer.

**Parameters:**
- `customer` (string): The customer's name

**Returns:**
- Formatted welcome message with the customer's name in uppercase

**Example:**
```go
WelcomeMessage("Alice") // "Welcome to the Tech Palace, ALICE"
```

### `AddBorder(welcomeMsg string, numStarsPerLine int) string`

Adds a border of stars around a welcome message.

**Parameters:**
- `welcomeMsg` (string): The message to border
- `numStarsPerLine` (int): Number of stars in the border lines

**Returns:**
- The message surrounded by star borders

**Example:**
```go
AddBorder("Hello", 5)
// *****
// Hello
// *****
```

### `CleanupMessage(oldMsg string) string`

Cleans up old marketing messages by removing all stars and trimming whitespace.

**Parameters:**
- `oldMsg` (string): The old marketing message to clean

**Returns:**
- Cleaned message without stars or extra whitespace

**Example:**
```go
CleanupMessage("***Special Offer***") // "Special Offer"
```

## Examples

### Complete Workflow

```go
package main

import (
    "fmt"
    "techpalace"
)

func main() {
    // Create welcome message
    welcome := WelcomeMessage("John Doe")
    fmt.Println(welcome)
    // Output: Welcome to the Tech Palace, JOHN DOE
    
    // Add border to message
    bordered := AddBorder(welcome, 15)
    fmt.Println(bordered)
    // Output:
    // ***************
    // Welcome to the Tech Palace, JOHN DOE
    // ***************
    
    // Clean up an old message
    oldAd := `
    ********************
    *   SALE ENDS SOON *
    ********************
    `
    cleaned := CleanupMessage(oldAd)
    fmt.Println(cleaned)
    // Output: SALE ENDS SOON
}
```

### Integration Example

```go
func CreatePromotionalBanner(customerName string, promotion string, borderWidth int) string {
    // Create welcome message
    welcome := WelcomeMessage(customerName)
    
    // Create promotion message with border
    promoWithBorder := AddBorder(promotion, borderWidth)
    
    // Combine messages
    return welcome + "\n\n" + promoWithBorder
}

func CleanOldCampaign(messages []string) []string {
    var cleaned []string
    for _, msg := range messages {
        cleaned = append(cleaned, CleanupMessage(msg))
    }
    return cleaned
}
```

## Requirements

- Go 1.16 or higher
- No external dependencies

## Best Practices

1. **Customer Names**: Always sanitize customer names before passing to WelcomeMessage
2. **Border Sizes**: Choose border sizes appropriate to message length (recommended: message length + 2)
3. **Message Cleanup**: Use CleanupMessage before storing or comparing marketing messages
4. **Performance**: For bulk operations, consider batching message processing

## License

This project is part of Exercism.io's Go track exercises and follows their licensing terms.

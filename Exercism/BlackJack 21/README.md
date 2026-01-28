# Blackjack Decision Maker

A Go implementation of basic Blackjack strategy for the first turn decision-making.

## Description

This program implements the basic Blackjack strategy for determining the optimal first move (Stand, Hit, Split, or automatically Win) based on the player's first two cards and the dealer's up card.

## Functions

### `ParseCard(card string) int`
Converts a card name to its corresponding Blackjack value.

**Parameters:**
- `card` (string): The name of the card (e.g., "ace", "king", "five")

**Returns:**
- `int`: The numerical value of the card according to Blackjack rules:
    - Number cards (1-9): Their face value
    - Face cards (jack, queen, king): 10
    - Ace: 11
    - Invalid cards: 0

### `FirstTurn(card1, card2, dealerCard string) string`
Determines the optimal first move in Blackjack based on basic strategy.

**Parameters:**
- `card1` (string): Player's first card
- `card2` (string): Player's second card
- `dealerCard` (string): Dealer's face-up card

**Returns:**
- `string`: A single character representing the decision:
    - `"P"` = Split
    - `"H"` = Hit
    - `"S"` = Stand
    - `"W"` = Automatic Win (Blackjack)

## Decision Logic

The function follows these rules in order:

1. **Split Aces**: If both player cards are aces (`"ace"`), return `"P"` (Split)
2. **Blackjack**: If player has 21 (Blackjack) AND dealer doesn't have a 10-value card or ace, return `"W"` (Win)
3. **Stand on 17+**: If player total is 17 or higher, return `"S"` (Stand)
4. **Hit on 12-16 vs 7+**: If player total is 12-16 AND dealer has 7 or higher, return `"H"` (Hit)
5. **Stand on 12-16 vs <7**: If player total is 12-16 AND dealer has less than 7, return `"S"` (Stand)
6. **Hit on <12**: If player total is less than 12, return `"H"` (Hit)
7. **Default**: Return `"!!!!"` for any unhandled case

## Usage Example

```go
package main

import (
    "fmt"
)

func main() {
    // Test cases
    fmt.Println(FirstTurn("ace", "ace", "jack"))    // Output: P (Split aces)
    fmt.Println(FirstTurn("ace", "king", "ace"))    // Output: S (Stand on 21 vs dealer ace)
    fmt.Println(FirstTurn("five", "queen", "ace"))  // Output: H (Hit 15 vs dealer ace)
    
    // More examples
    fmt.Println(FirstTurn("ten", "ace", "five"))    // W (Blackjack)
    fmt.Println(FirstTurn("seven", "seven", "ten")) // H (14 vs 10)
    fmt.Println(FirstTurn("six", "six", "six"))     // H (12 vs 6)
}
```

## Card Values Reference

| Card Name | Value |
|-----------|-------|
| "one"     | 1     |
| "two"     | 2     |
| "three"   | 3     |
| "four"    | 4     |
| "five"    | 5     |
| "six"     | 6     |
| "seven"   | 7     |
| "eight"   | 8     |
| "nine"    | 9     |
| "ten"     | 10    |
| "jack"    | 10    |
| "queen"   | 10    |
| "king"    | 10    |
| "ace"     | 11    |

## Sample Output

```
P
S
H
```

## Testing

You can test the program with various scenarios:

```go
// Blackjack scenarios
fmt.Println("Blackjack vs low card:", FirstTurn("ace", "king", "six"))  // W
fmt.Println("Blackjack vs 10/ace:", FirstTurn("ace", "king", "ten"))    // S

// Split scenarios
fmt.Println("Split aces:", FirstTurn("ace", "ace", "seven"))            // P

// Hard totals
fmt.Println("Hard 16 vs 7:", FirstTurn("ten", "six", "seven"))          // H
fmt.Println("Hard 16 vs 6:", FirstTurn("ten", "six", "six"))            // S
fmt.Println("Hard 12 vs 3:", FirstTurn("seven", "five", "three"))       // H
```

## Limitations

- This implements only **basic strategy** for the **first turn**
- Does not account for card counting or advanced strategies
- Assumes standard Blackjack rules (single deck, dealer stands on soft 17, etc.)
- The `"!!!!"` return value indicates an unhandled case that should be addressed

## Dependencies

- Go 1.13 or higher
- No external packages required
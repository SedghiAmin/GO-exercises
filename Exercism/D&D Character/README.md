# Dungeons & Dragons Character Generator

A Go implementation of a random character generator for Dungeons & Dragons (D&D) following the standard rules for generating ability scores and calculating hit points.

## Overview

This program generates a D&D character with six ability scores (Strength, Dexterity, Constitution, Intelligence, Wisdom, and Charisma) determined randomly according to the standard D&D rules, then calculates the character's initial hit points.

## Rules Implemented

### Ability Score Generation
For each of the six abilities:
1. Roll four 6-sided dice (4d6)
2. Discard the lowest roll
3. Sum the three highest rolls
4. Assign the result to the ability

### Constitution Modifier Calculation
The Constitution modifier is calculated as:
```
Modifier = floor((Constitution - 10) / 2)
```
Where `floor()` means rounding down to the nearest integer.

### Hit Points Calculation
Initial hit points are calculated as:
```
Hit Points = 10 + Constitution Modifier
```

## Code Structure

### Data Structure
```go
type Character struct {
    Strength     int
    Dexterity    int
    Constitution int
    Intelligence int
    Wisdom       int
    Charisma     int
    Hitpoints    int
}
```

### Functions

#### `Modifier(score int) int`
Calculates the ability modifier for a given ability score using the formula: `floor((score - 10) / 2)`.

#### `Ability() int`
Generates a random ability score by:
1. Rolling four 6-sided dice (values 1-6)
2. Keeping track of the smallest roll
3. Returning the sum of the three highest rolls

#### `GenerateCharacter() Character`
Creates a new character with random scores for all six abilities and calculates the hit points based on the Constitution score.

## Usage

Run the program:
```bash
go run main.go
```

Example output:
```
Strength: 14
Dexterity: 11
Constitution: 16 (Modifier: 3)
Intelligence: 9
Wisdom: 13
Charisma: 15
Hitpoints: 13
```

## Algorithm Details

### Random Number Generation
- The program uses Go's `math/rand` package for random number generation
- The random seed is initialized with the current Unix timestamp (`time.Now().UnixNano()`) to ensure different results each run
- Dice rolls are simulated using `rand.Intn(6) + 1` to get values between 1 and 6

### Ability Score Range
- Minimum possible score: 3 (if all four dice show 1)
- Maximum possible score: 18 (if all four dice show 6)
- Average expected score: ~12-13

### Edge Cases Handled
- The `Modifier()` function correctly handles negative scores using `math.Floor()` for proper rounding down
- Constitution scores below 10 result in negative modifiers, correctly reducing hit points
- The `Ability()` function efficiently finds and discards the lowest roll without full sorting

## Dependencies
- Standard Go libraries only
- No external dependencies required

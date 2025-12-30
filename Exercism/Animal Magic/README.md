# Animal Magic Wizardry

A Go Program for generating magical wizardry data including dice rolls, wand energy, and shuffled animal sequences.

## Features

- **Roll a Magical Die**: Generate random numbers for wizard games
- **Generate Wand Energy**: Create random wand energy levels
- **Shuffle Animals**: Randomize the order of magical animals

## Functions

### RollADie()
```go
func RollADie() int
```
Returns a random integer between 1 and 20 (inclusive), simulating a 20-sided magical die roll.

### GenerateWandEnergy()
```go
func GenerateWandEnergy() float64
```
Returns a random floating-point number between 0.0 (inclusive) and 12.0 (exclusive), representing wand energy levels.

### ShuffleAnimals()
```go
func ShuffleAnimals() []string
```
Returns a slice containing eight animal strings in random order:
- ant, beaver, cat, dog, elephant, fox, giraffe, hedgehog

## Usage Example

```go
package main

import "fmt"

func main() {
    // Roll a magical die
    diceRoll := RollADie()
    fmt.Printf("You rolled: %d\n", diceRoll)

    // Generate wand energy
    energy := GenerateWandEnergy()
    fmt.Printf("Wand energy: %.2f\n", energy)

    // Get shuffled animals
    animals := ShuffleAnimals()
    fmt.Printf("Magical animals: %v\n", animals)
}
```

## Technical Details

### Random Number Generation
- Uses Go's `math/rand` package with proper seeding
- Seeding is done using current time for true randomness
- Thread-safe operations within individual function calls

### Animal Shuffling
- Implements Fisher-Yates shuffle algorithm via `rand.Shuffle()`
- Returns a new shuffled slice each time
- Maintains all original elements in random order

## Dependencies
- Go standard library only (`math/rand`, `time`)

The test suite validates:
- Die rolls stay within 1-20 range
- Wand energy stays within 0.0-12.0 range
- Animal shuffling produces random orders
- All original animals are preserved

## Use Cases
- Wizard role-playing games
- Random data generation for testing
- Educational examples of Go's random functions
- Game development prototypes

## Notes
- Results are pseudo-random but sufficient for most applications
- For cryptographic randomness, consider using `crypto/rand` instead
- The animal list is fixed and cannot be customized
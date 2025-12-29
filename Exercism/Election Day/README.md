# Votes Manager

A simple Go program that provides vote counting functionality for election systems. This package demonstrates pointer operations, struct usage, and map manipulations in Go.

## Features

- **Vote Counter**: Create and manage integer vote counters using pointers
- **Election Results**: Track candidate names and their vote counts
- **Vote Operations**: Increment and decrement vote counts safely
- **Nil Safety**: Handle nil pointers gracefully in vote counting

## Functions

### Vote Counter Functions
- `NewVoteCounter(initialVotes int) *int` - Creates a new vote counter
- `VoteCount(counter *int) int` - Safely extracts votes from a counter (handles nil)
- `IncrementVoteCount(counter *int, increment int)` - Increments vote count

### Election Result Functions
- `NewElectionResult(candidateName string, votes int) *ElectionResult` - Creates election result
- `DisplayResult(result *ElectionResult) string` - Formats result for display
- `DecrementVotesOfCandidate(results map[string]int, candidate string)` - Decrements candidate's votes

## Usage Example

```go
// Create a vote counter
counter := NewVoteCounter(3)
votes := VoteCount(counter) // 3

// Increment votes
IncrementVoteCount(counter, 2)

// Create and display election result
result := NewElectionResult("Alice", 42)
display := DisplayResult(result) // "Alice (42)"

// Decrement votes in a map
results := map[string]int{"Bob": 10}
DecrementVotesOfCandidate(results, "Bob") // Bob now has 9 votes
```

## Notes

- The program demonstrates proper pointer usage in Go
- All functions handle edge cases (like nil pointers)
- Election results are managed as structs with automatic pointer dereferencing
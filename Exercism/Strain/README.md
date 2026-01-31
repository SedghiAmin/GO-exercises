# Go Collection Filter Operations

This program provides generic implementations of `Keep` and `Discard` operations for filtering collections in Go.

## Overview

The `Keep` and `Discard` functions allow you to filter slices based on a predicate function. These are fundamental functional programming operations implemented using Go's generics feature.

## Functions

### `Keep[T any](arr []T, f func(T) bool) []T`
Returns a new slice containing only the elements from the input slice for which the predicate function `f` returns `true`.

### `Discard[T any](arr []T, f func(T) bool) []T`
Returns a new slice containing only the elements from the input slice for which the predicate function `f` returns `false`.

## Key Features

- **Type-safe**: Uses Go generics to work with any data type
- **Immutable**: Original slices are not modified
- **Complementary**: `Keep` and `Discard` are inverses of each other
- **Simple API**: Easy to understand and use

## Installation

```go
go get github.com/yourusername/strain
```

Or simply copy the functions into your project.

## Usage Examples

### Basic Example with Numbers

```go
numbers := []int{1, 2, 3, 4, 5, 6}
isEven := func(n int) bool { return n%2 == 0 }

evens := Keep(numbers, isEven)    // [2, 4, 6]
odds := Discard(numbers, isEven)  // [1, 3, 5]
```

### Working with Strings

```go
words := []string{"apple", "banana", "cherry"}
startsWithA := func(s string) bool { return strings.HasPrefix(s, "a") }

aWords := Keep(words, startsWithA)      // ["apple"]
nonAWords := Discard(words, startsWithA) // ["banana", "cherry"]
```

### Custom Types

```go
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Ali", 25},
    {"Reza", 17},
    {"Sara", 30},
}

isAdult := func(p Person) bool { return p.Age >= 18 }

adults := Keep(people, isAdult)    // [Ali(25), Sara(30)]
minors := Discard(people, isAdult) // [Reza(17)]
```

## Mathematical Property

For any collection `C` and predicate `P`:
```
Keep(C, P) ∪ Discard(C, P) = C
Keep(C, P) ∩ Discard(C, P) = ∅
```

## Performance Considerations

- Time Complexity: O(n) where n is the length of the input slice
- Space Complexity: O(k) where k is the number of elements matching the predicate
- Uses `append()` which handles memory allocation efficiently

## Edge Cases

- Empty slices return empty slices
- All elements matching predicate: `Discard` returns empty slice
- No elements matching predicate: `Keep` returns empty slice
- Nil slices are handled safely

## Testing

Run the examples:
```bash
go run main.go
```

Expected output:
```
Numbers: [1 2 3 4 5 6]
Even Numbers (Keep): [2 4 6]
Odd Numbers (Discard): [1 3 5]

Words: [apple banana cherry date elderberry]
With Words 'a' (Keep): [apple]
Without Words 'a' (Discard): [banana cherry date elderberry]

adult individuals (Keep):
  Ali (age: 25)
  Sara (age: 30)
People under 18 years old (Discard):
  Reza (age: 17)
  Mona (age: 16)
```

## Dependencies

- Go 1.18+ (for generics support)
- Standard library only

## Related Concepts

These functions are similar to:
- `filter` in functional programming languages
- LINQ's `Where` in C#
- List comprehensions in Python

## Contributing

Feel free to submit issues and pull requests for improvements.

## License

MIT
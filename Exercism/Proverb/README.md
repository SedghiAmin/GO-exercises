# English Proverb Generator

A Go program that generates the classic English proverb "For Want of a Nail" based on a sequence of items.

## Description

This program creates a chain proverb where each item's loss leads to the next item's loss, following the traditional structure of the "For Want of a Nail" proverb.

## The Proverb Pattern

The proverb follows this pattern:
```
For want of a [item1] the [item2] was lost.
For want of a [item2] the [item3] was lost.
...
And all for the want of a [item1].
```

## Code Structure

### Main Function: `Proverb(rhyme []string) []string`

**Parameters:**
- `rhyme`: A slice of strings representing the sequence of items in the proverb

**Returns:**
- A slice of strings containing the complete proverb

**Logic:**
1. If the input slice is empty, returns an empty slice
2. Creates lines for each consecutive pair: `"For want of a X the Y was lost."`
3. Adds the final line: `"And all for the want of a [first item]."`

### Example Usage

```go
func main() {
    result := Proverb([]string{"nail", "shoe", "horse"})
    
    for i, line := range result {
        fmt.Printf("Line %d: %s\n", i+1, line)
    }
}
```

**Output:**
```
Line 1: For want of a nail the shoe was lost.
Line 2: For want of a shoe the horse was lost.
Line 3: And all for the want of a nail.
```

## More Examples

### Standard Proverb
```go
Proverb([]string{"nail", "shoe", "horse", "rider", "message", "battle", "kingdom"})
```

**Output:**
```
1. For want of a nail the shoe was lost.
2. For want of a shoe the horse was lost.
3. For want of a horse the rider was lost.
4. For want of a rider the message was lost.
5. For want of a message the battle was lost.
6. For want of a battle the kingdom was lost.
7. And all for the want of a nail.
```

### Short Version
```go
Proverb([]string{"drop", "bucket"})
```

**Output:**
```
1. For want of a drop the bucket was lost.
2. And all for the want of a drop.
```

### Single Item
```go
Proverb([]string{"key"})
```

**Output:**
```
1. And all for the want of a key.
```

### Empty Input
```go
Proverb([]string{})
```

**Output:**
```
[] (empty slice)
```

## The Original Proverb

The program is based on the traditional English proverb:
> "For want of a nail the shoe was lost.
> For want of a shoe the horse was lost.
> For want of a horse the rider was lost.
> For want of a rider the message was lost.
> For want of a message the battle was lost.
> For want of a battle the kingdom was lost.
> And all for the want of a horseshoe nail."

## Compilation and Execution

```bash
# Run directly
go run proverb.go

# Or compile and run
go build -o proverb proverb.go
./proverb
```

## Algorithm Explanation

1. **Input Validation**: Checks if the input slice is empty
2. **Memory Allocation**: Creates a new slice with the same length as input
3. **Chain Generation**: Iterates through consecutive pairs to create loss statements
4. **Final Line**: Adds the concluding sentence referencing the first item

## Edge Cases Handled

- ✅ Empty input slice returns empty slice
- ✅ Single item returns only the final line
- ✅ Two items creates one chain link + final line
- ✅ Multiple items creates full chain

## Use Cases

- **Educational**: Teaching Go slices and string manipulation
- **Literary**: Generating chain-style proverbs
- **Creative**: Custom proverb creation for stories or games
- **Testing**: Example for testing slice operations

## Dependencies

- Standard Go library only
- No external packages required

## License

Free to use for educational and personal projects.
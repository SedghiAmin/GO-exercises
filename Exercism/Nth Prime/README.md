# Prime Number Finder

A Go package that provides functionality to find the nth prime number efficiently.

## Usage

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/yourusername/prime-finder"
)

func main() {
    // Find the 6th prime number
    prime, err := prime.Nth(6)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("The 6th prime number is: %d\n", prime) // Output: 13
}
```

## Function Documentation

### `Nth(n int) (int, error)`

Returns the nth prime number (where n starts from 1).

**Parameters:**
- `n`: The position of the prime number to find (1-indexed)

**Returns:**
- `int`: The nth prime number
- `error`: Error if n < 1

**Examples:**

```go
prime, _ := Nth(1)  // returns 2, nil
prime, _ := Nth(2)  // returns 3, nil
prime, _ := Nth(3)  // returns 5, nil
prime, _ := Nth(4)  // returns 7, nil
prime, _ := Nth(5)  // returns 11, nil
prime, _ := Nth(6)  // returns 13, nil
prime, _ := Nth(10) // returns 29, nil
```

## Algorithm

The function uses a straightforward trial division algorithm:
1. Handles edge cases (n < 1 returns error, n = 1 returns 2)
2. Iterates through odd numbers starting from 3
3. For each number, checks divisibility by all smaller numbers
4. When a prime is found, increments the counter until reaching the nth prime

## Performance Characteristics

- **Time Complexity:** O(n² log n) - For finding the nth prime, each number up to approximately n log n is checked for primality by testing divisibility up to itself
- **Space Complexity:** O(1) - Uses only constant extra space

## Example Output

```go
package main

import (
    "fmt"
    "prime"
)

func main() {
    for i := 1; i <= 10; i++ {
        prime, _ := prime.Nth(i)
        fmt.Printf("Prime #%d: %d\n", i, prime)
    }
}
```

Output:
```
Prime #1: 2
Prime #2: 3
Prime #3: 5
Prime #4: 7
Prime #5: 11
Prime #6: 13
Prime #7: 17
Prime #8: 19
Prime #9: 23
Prime #10: 29
```

## Error Handling

The function returns an error when:
- `n < 1`: Returns "incorrect number" error

```go
prime, err := Nth(0)
if err != nil {
    fmt.Println(err) // Output: incorrect number
}
```

## Testing

Run tests with:

```bash
go test -v
```

Example test:

```go
func TestNth(t *testing.T) {
    tests := []struct {
        n       int
        want    int
        wantErr bool
    }{
        {1, 2, false},
        {2, 3, false},
        {3, 5, false},
        {4, 7, false},
        {5, 11, false},
        {6, 13, false},
        {0, 0, true},
        {-1, 0, true},
    }
    
    for _, tt := range tests {
        got, err := Nth(tt.n)
        if (err != nil) != tt.wantErr {
            t.Errorf("Nth(%d) error = %v, wantErr %v", tt.n, err, tt.wantErr)
            continue
        }
        if got != tt.want {
            t.Errorf("Nth(%d) = %d, want %d", tt.n, got, tt.want)
        }
    }
}
```

## License

MIT License - see LICENSE file for details

# Sieve of Eratosthenes in Go

This program provides an implementation of the ancient **Sieve of Eratosthenes** algorithm for finding all prime numbers up to a given limit.

## Algorithm Overview

The Sieve of Eratosthenes is an efficient algorithm for finding all prime numbers up to a specified integer. It works by iteratively marking the multiples of each prime number, starting from 2.

### How It Works

1. Start with a list of numbers from 2 to the limit
2. Take the first unmarked number (2) - it's prime
3. Mark all multiples of 2 as composite (not prime)
4. Move to the next unmarked number (3) - it's prime
5. Mark all multiples of 3 as composite
6. Continue this process until you've processed all numbers up to the limit

## Code Implementation

```go
package main

import "fmt"

func Sieve(limit int) []int {
    marked := make(map[int]bool, limit)
    prime := make([]int, 0)
    
    for i := 2; i <= limit; i++ {
        if !marked[i] {
            prime = append(prime, i)
            // Start marking from i*i as smaller multiples are already marked
            for j := i * i; j <= limit; j += i {
                marked[j] = true
            }
        }
    }
    return prime
}

func main() {
    fmt.Println(Sieve(1000))
}
```

## Key Optimizations

### 1. Starting from i²
The algorithm starts marking multiples from `i * i` because all smaller multiples (2i, 3i, ..., (i-1)i) have already been marked by previous primes.

### 2. Map-based Marking
Using a `map[int]bool` for marking composites provides:
- O(1) lookup time
- Memory efficiency for sparse marking
- Clear semantics (true = composite number)

## Usage Example

```go
// Find primes up to 30
primes := Sieve(30)
fmt.Println(primes) // Output: [2 3 5 7 11 13 17 19 23 29]

// Find primes up to 10
primes = Sieve(10)
fmt.Println(primes) // Output: [2 3 5 7]
```

## Complexity Analysis

- **Time Complexity**: O(n log log n) - nearly linear
- **Space Complexity**: O(n) - for storing composite markers

## Performance Characteristics

The implementation is optimized for:
- **Readability**: Clear, idiomatic Go code
- **Correctness**: Properly implements the classic algorithm
- **Efficiency**: Uses the i² optimization to reduce operations

## Limitations

- For very large limits (> 10 million), the map-based approach may consume significant memory
- For production use with extremely large datasets, consider using a boolean slice instead of a map for better memory efficiency

## Alternative Implementation

For better memory efficiency with large limits, consider using a boolean slice:

```go
func SieveOptimized(limit int) []int {
    if limit < 2 {
        return []int{}
    }
    
    isComposite := make([]bool, limit+1)
    primes := make([]int, 0)
    
    for i := 2; i <= limit; i++ {
        if !isComposite[i] {
            primes = append(primes, i)
            if i*i <= limit {
                for j := i * i; j <= limit; j += i {
                    isComposite[j] = true
                }
            }
        }
    }
    return primes
}
```

## Testing

To test the implementation:

```go
func TestSieve(t *testing.T) {
    testCases := []struct {
        limit    int
        expected []int
    }{
        {1, []int{}},
        {2, []int{2}},
        {10, []int{2, 3, 5, 7}},
        {30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
    }
    
    for _, tc := range testCases {
        result := Sieve(tc.limit)
        if !reflect.DeepEqual(result, tc.expected) {
            t.Errorf("Sieve(%d) = %v; expected %v", tc.limit, result, tc.expected)
        }
    }
}
```

## License

This code is provided as a learning example and is free to use, modify, and distribute.
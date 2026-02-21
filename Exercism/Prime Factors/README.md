## Code Analysis for Prime Factors Calculation

### Line by Line Explanation:

```go
package main

import "fmt"

func Factors(n int64) []int64 {
    // Create an empty slice to store prime factors
    out := make([]int64, 0)
    
    // Start i from 2 (the smallest prime number)
    var i int64
    for i = 2; i <= n; i++ {
        // While n is divisible by i:
        // 1. Add i to the factors slice
        // 2. Divide n by i
        for n%i == 0 {
            out = append(out, i)
            n = n / i
        }
    }
    return out
}

func main() {
    // Test cases
    fmt.Println(Factors(10))  // Output: [2 5]
    fmt.Println(Factors(20))  // Output: [2 2 5]
    fmt.Println(Factors(60))  // Output: [2 2 3 5]
}
```

### How the Algorithm Works:

**Example with n = 60:**

| Step | i | n (before) | n % i == 0? | Action | n (after) | Factors |
|------|---|------------|-------------|--------|-----------|---------|
| 1 | 2 | 60 | Yes | Add 2, divide by 2 | 30 | [2] |
| 2 | 2 | 30 | Yes | Add 2, divide by 2 | 15 | [2, 2] |
| 3 | 2 | 15 | No | Increment i to 3 | 15 | [2, 2] |
| 4 | 3 | 15 | Yes | Add 3, divide by 3 | 5 | [2, 2, 3] |
| 5 | 3 | 5 | No | Increment i to 4 | 5 | [2, 2, 3] |
| 6 | 4 | 5 | No | Increment i to 5 | 5 | [2, 2, 3] |
| 7 | 5 | 5 | Yes | Add 5, divide by 5 | 1 | [2, 2, 3, 5] |
| 8 | 5 | 1 | Loop ends (i ≤ n is false) | - | - | [2, 2, 3, 5] |

### Key Points:

1. **Two nested loops:**
    - Outer loop: tries different divisors (i) from 2 upwards
    - Inner loop: repeatedly divides by the same divisor while possible

2. **Why this works:**
    - When we find a divisor, we divide the number completely by it
    - This ensures we only get prime factors (composite numbers won't divide the remaining number)
    - Example: After extracting all 2's from 60 (getting 15), 4 won't divide 15

3. **Efficiency:**
    - The algorithm automatically skips testing composite numbers
    - After removing factor 2 from 60, n becomes 15, so 4 never gets tested on numbers that could be divisible by 4

### Sample Output:
```
[2 5]
[2 2 5]
[2 2 3 5]
```
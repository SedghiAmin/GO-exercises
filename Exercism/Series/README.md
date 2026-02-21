## Code Analysis: Series Exercise

### Complete Code with Explanation:

```go
package main

import "fmt"

// All returns all contiguous substrings of length n from string s
// Example: All(3, "49142") returns ["491", "914", "142"]
func All(n int, s string) []string {
    // Create empty slice to store results
    out := make([]string, 0)
    
    // Loop through the string, taking n characters at a time
    // i <= len(s)-n ensures we don't go beyond the string length
    for i := 0; i <= len(s)-n; i++ {
        // Extract substring from i to i+n and append to result
        out = append(out, s[i:i+n])
    }
    return out
}

// UnsafeFirst returns the first substring of length n
// Called "unsafe" because it panics if n > len(s)
// Example: UnsafeFirst(3, "49142") returns "491"
func UnsafeFirst(n int, s string) string {
    return s[0:n]  // Direct slice operation - unsafe if n > len(s)
}

func main() {
    // Test the functions
    fmt.Println(All(3, "49142"))     // Output: [491 914 142]
    fmt.Println(UnsafeFirst(3, "49142")) // Output: 491
}
```

### Output:
```
[491 914 142]
491
```

### How It Works:

**For All(3, "49142"):**
| i | Substring | Added to result |
|---|-----------|-----------------|
| 0 | "491" | ["491"] |
| 1 | "914" | ["491", "914"] |
| 2 | "142" | ["491", "914", "142"] |

### Edge Cases and Improvements:

The code works correctly for valid inputs, but here are potential improvements:

```go
// Improved version with error checking
func SafeAll(n int, s string) []string {
    // Handle invalid inputs
    if n < 1 || n > len(s) {
        return []string{}  // Return empty slice for invalid input
    }
    
    // Pre-allocate slice with exact capacity for better performance
    result := make([]string, 0, len(s)-n+1)
    
    for i := 0; i <= len(s)-n; i++ {
        result = append(result, s[i:i+n])
    }
    return result
}

// Safe version of First with ok flag (from bonus exercise)
func First(n int, s string) (string, bool) {
    if n < 1 || n > len(s) {
        return "", false
    }
    return s[0:n], true
}
```

### Why "UnsafeFirst"?
The function is named "UnsafeFirst" because:
- If you call `UnsafeFirst(10, "49142")`, it will cause a **panic** (runtime error)
- In Go, slicing beyond string length causes panic
- The bonus exercise adds a `First` function that returns an `ok` boolean to handle this safely

### Testing Different Scenarios:
```go
func main() {
    // Normal cases
    fmt.Println(All(3, "49142"))     // [491 914 142]
    fmt.Println(All(4, "49142"))     // [4914 9142]
    
    // Edge cases
    fmt.Println(All(5, "49142"))     // [49142]
    fmt.Println(All(6, "49142"))     // [] (empty)
    
    // UnsafeFirst examples
    fmt.Println(UnsafeFirst(3, "49142")) // "491"
    // fmt.Println(UnsafeFirst(10, "49142")) // This would PANIC!
}
```
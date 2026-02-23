# IntList - Functional List Operations for Integers

A Go program that provides a functional programming style interface for integer slices with common list operations like `map`, `filter`, `fold`, and more.

## Overview

`IntList` is a custom type based on `[]int` that implements various functional programming patterns. It allows you to work with integer lists in a declarative way, similar to languages that support functional programming natively.

## Installation

```bash
go get github.com/yourusername/intlist
```

## Usage

```go
package main

import "fmt"

func main() {
    list := IntList{1, 2, 3, 4, 5}
    
    // Map: square each element
    squares := list.Map(func(x int) int {
        return x * x
    })
    fmt.Println(squares) // [1 4 9 16 25]
    
    // Filter: keep only even numbers
    evens := list.Filter(func(x int) bool {
        return x%2 == 0
    })
    fmt.Println(evens) // [2 4]
    
    // Foldl: sum all elements
    sum := list.Foldl(func(acc, x int) int {
        return acc + x
    }, 0)
    fmt.Println(sum) // 15
}
```

## API Reference

### Type Definition

```go
type IntList []int
```

### Methods

#### `Map(fn func(int) int) IntList`

Applies a function to each element of the list and returns a new list with the results.

```go
list := IntList{1, 2, 3}
doubled := list.Map(func(x int) int { return x * 2 })
// Result: [2, 4, 6]
```

#### `Filter(fn func(int) bool) IntList`

Returns a new list containing only the elements that satisfy the predicate function.

```go
list := IntList{1, 2, 3, 4, 5}
evens := list.Filter(func(x int) bool { return x%2 == 0 })
// Result: [2, 4]
```

#### `Foldl(fn func(int, int) int, initial int) int`

Reduces the list from left to right using the accumulator function.

```go
list := IntList{1, 2, 3, 4}
sum := list.Foldl(func(acc, x int) int { return acc + x }, 0)
// ((((0 + 1) + 2) + 3) + 4) = 10
```

#### `Foldr(fn func(int, int) int, initial int) int`

Reduces the list from right to left using the accumulator function.

```go
list := IntList{1, 2, 3, 4}
// For addition, Foldr gives same result as Foldl
sum := list.Foldr(func(x, acc int) int { return x + acc }, 0)
// (1 + (2 + (3 + (4 + 0)))) = 10
```

#### `Reverse() IntList`

Returns a new list with elements in reverse order.

```go
list := IntList{1, 2, 3}
reversed := list.Reverse()
// Result: [3, 2, 1]
```

#### `Append(lst IntList) IntList`

Returns a new list with the elements of the original list followed by the elements of the provided list.

```go
list1 := IntList{1, 2, 3}
list2 := IntList{4, 5, 6}
combined := list1.Append(list2)
// Result: [1, 2, 3, 4, 5, 6]
```

#### `Concat(lists []IntList) IntList`

Concatenates multiple lists together.

```go
list := IntList{1, 2, 3}
others := []IntList{{4, 5}, {6, 7}, {8, 9}}
result := list.Concat(others)
// Result: [1, 2, 3, 4, 5, 6, 7, 8, 9]
```

#### `Length() int`

Returns the number of elements in the list.

```go
list := IntList{1, 2, 3, 4, 5}
length := list.Length()
// Result: 5
```

## Complete Example

```go
package main

import "fmt"

type IntList []int

func main() {
    source := IntList{1, 2, 3}
    ext := IntList{4, 5, 6}
    lists := []IntList{{5, 6, 3}, {9, 4, 1}}
    
    // Append example
    fmt.Println(source.Append(ext)) // [1 2 3 4 5 6]
    
    // Concat example
    fmt.Println(source.Concat(lists)) // [1 2 3 5 6 3 9 4 1]
    
    // Filter example
    isEven := func(x int) bool { return x%2 == 0 }
    fmt.Println(source.Filter(isEven)) // [2]
    
    // Map example
    square := func(x int) int { return x * x }
    fmt.Println(source.Map(square)) // [1 4 9]
    
    // Fold examples
    add := func(x, y int) int { return x + y }
    fmt.Println(source.Foldl(add, 0)) // 6
    fmt.Println(source.Foldr(add, 0)) // 6
    
    // Reverse example
    fmt.Println(source.Reverse()) // [3 2 1]
}
```


# Generic Linked List in Go
A simple, type-safe generic linked list implementation in Go.

# Overview
This package provides a generic singly-linked list that can store values of any type using Go's generics.

# Features
✅ Type-safe generic implementation

✅ Add elements to the end of the list

✅ Convert list to slice for easy iteration

✅ Simple and clean API

# Usage
``` go
package main

import "fmt"

func main() {
    // Create a new list with integer values
    intList := &List[int]{Node: 10}
    intList.Add(20)
    intList.Add(30)
    
    // Convert to slice and print
    fmt.Println(intList.toSlice()) // [10 20 30]
    
    // Works with any type
    stringList := &List[string]{Node: "hello"}
    stringList.Add("world")
    fmt.Println(stringList.toSlice()) // ["hello" "world"]
}
```

# API Reference
List[T any]
Generic linked list structure.

# Methods
Add(node T) *List[T]
Adds a new node to the end of the list.

# Parameters:

node T: Value to add

# Returns:
[]T: Slice containing all list elements in order

# Requirements
Go 1.18+ (for generics support)

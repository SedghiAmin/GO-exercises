# Go Fibonacci with Memoization

This is a simple and efficient implementation of the Fibonacci sequence in **Go**, using **memoization** to avoid redundant recursive calls.

# Description

The Fibonacci sequence is defined as:

F(0) = 0
F(1) = 1
F(n) = F(n−1) + F(n−2)

A naive recursive implementation recalculates many values multiple times, resulting in exponential time complexity.  
By using **memoization**, we cache previously computed results in a map and achieve linear time complexity — while keeping the simplicity of recursion.

# Code Overview

```go
package main

import "fmt"

func fibonacci() func(int) int {
	cache := map[int]int{
		0: 0,
		1: 1,
	}

	var fib func(int) int
	fib = func(n int) int {
		if val, ok := cache[n]; ok {
			return val
		}
		cache[n] = fib(n-1) + fib(n-2)
		return cache[n]
	}

	return fib
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}

# How It Works
The fibonacci() function returns a closure fib that:

Stores computed Fibonacci values in a map[int]int cache.

Recursively computes F(n) only if it hasn’t been calculated before.

The first time fib(n) is called, it computes and stores the result.

Subsequent calls retrieve results directly from the cache.

# Example Output

0
1
1
2
3
5
8
13
21
34

# Performance
Implementation	Time Complexity	Space Complexity
Naive recursion	O(2ⁿ)	O(n)
With memoization	O(n)	O(n)

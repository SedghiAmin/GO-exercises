package main

import (
	"fmt"
	"strings"
)

func Keep[T any](arr []T, f func(T) bool) []T {
	var r []T

	for _, v := range arr {
		if f(v) {
			r = append(r, v)
		}
	}

	return r
}

func Discard[T any](arr []T, f func(T) bool) []T {
	var r []T
	for _, v := range arr {
		if !f(v) {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	//Example 1: Working with numbers
	numbers := []int{1, 2, 3, 4, 5, 6}

	isEven := func(n int) bool { return n%2 == 0 }

	evens := Keep(numbers, isEven)
	odds := Discard(numbers, isEven)

	fmt.Println("Numbers:", numbers)
	fmt.Println("Even Numbers (Keep):", evens)  // [2, 4, 6]
	fmt.Println("Odd Numbers (Discard):", odds) // [1, 3, 5]

	// Example 2: Working with strings
	words := []string{"apple", "banana", "cherry", "date", "elderberry"}

	startsWithA := func(s string) bool { return strings.HasPrefix(s, "a") }

	aWords := Keep(words, startsWithA)
	nonAWords := Discard(words, startsWithA)

	fmt.Println("\nWords:", words)
	fmt.Println("With Words 'a' (Keep):", aWords)          // ["apple"]
	fmt.Println("Without Words 'a' (Discard):", nonAWords) // ["banana", "cherry", "date", "elderberry"]

	// Example 3: Working with arbitrary structures
	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Ali", 25},
		{"Reza", 17},
		{"Sara", 30},
		{"Mona", 16},
	}

	isAdult := func(p Person) bool { return p.Age >= 18 }

	adults := Keep(people, isAdult)
	minors := Discard(people, isAdult)

	fmt.Println("\nadult individuals (Keep):")
	for _, p := range adults {
		fmt.Printf("  %s (age: %d)\n", p.Name, p.Age)
	}

	fmt.Println("People under 18 years old (Discard):")
	for _, p := range minors {
		fmt.Printf("  %s (age: %d)\n", p.Name, p.Age)
	}
}

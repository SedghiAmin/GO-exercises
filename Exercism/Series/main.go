package main

import "fmt"

func All(n int, s string) []string {
	out := make([]string, 0)
	for i := 0; i <= len(s)-3; i++ {
		out = append(out, s[i:i+n])
	}
	return out
}

func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

func main() {
	fmt.Println(All(3, "49142"))
	fmt.Println(UnsafeFirst(3, "49142"))
}

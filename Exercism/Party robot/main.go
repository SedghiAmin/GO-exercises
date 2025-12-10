package main

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func main() {
	fmt.Println(Welcome("Amin"))
}

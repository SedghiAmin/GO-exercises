package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return " Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func main() {
	fmt.Println(WelcomeMessage("Judy"))
	// => Welcome to the Tech Palace, JUDY
}

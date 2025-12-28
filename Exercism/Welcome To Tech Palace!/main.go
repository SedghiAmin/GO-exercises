package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return " Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
	msg = strings.TrimSpace(msg)
	return msg
}

func main() {
	fmt.Println(WelcomeMessage("Judy"))
	// => Welcome to the Tech Palace, JUDY

	fmt.Println(AddBorder("Welcome!", 10))
	/* => output:
	***********
	Welcome!
	***********
	*/

	message := `
**************************
*    BUY NOW, SAVE 10%   *
**************************
`

	fmt.Println(CleanupMessage(message))
	// => BUY NOW, SAVE 10%
}

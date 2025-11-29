package greetings //Declare greetings package to collect related functions.

import(
	"fmt"
)

//Implement Hello function to return the greeting
func Hello (name string) string{
	msg:= fmt.Sprintf("Hi %v. Welcome to Go!", name)
	return msg
}

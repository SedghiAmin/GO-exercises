package greetings //Declare greetings package to collect related functions.

import (
	"errors"
	"fmt"
)

//Implement Hello function to return the greeting
func Hello (name string) (string, error){

	if name == ""{
		return "", errors.New("empty name")
	}

	msg:= fmt.Sprintf("Hi %v. Welcome to Go!", name)

	return msg, nil
}

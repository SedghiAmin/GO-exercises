package greetings //Declare greetings package to collect related functions.

import (
	"errors"
	"fmt"
	"math/rand"
)

//Implement Hello function to return the greeting
func Hello (name string) (string, error){

	if name == ""{
		return "", errors.New("empty name")
	}

	msg:= fmt.Sprintf(randomGreeting(), name)

	return msg, nil
}

func randomGreeting() string{
	formats:= []string{
		"Hi, %v. Welcome to GO!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
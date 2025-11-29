package greetings

import(
	"fmt"
)

func Hello (name string) string{
	msg:= fmt.Sprintf("Hi %v. Welcome to Go!", name)
	return msg
}

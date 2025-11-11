package main

import (
	"fmt"
	"time"
)

type MyError struct{
	when time.Time
	what string
}

func (e *MyError) Error() string{
	return fmt.Sprintf("at %v: %v", e.when, e.what)
}

func run () error{
	return &MyError{
		time.Now(),
		"It didnt work !",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

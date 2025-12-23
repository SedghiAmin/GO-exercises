package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, greeter Greeter) string {
	language := greeter.LanguageName()
	msg := greeter.Greet(name)
	return fmt.Sprintf("I can speak %v: %v ", language, msg)
}

type GermanGreeter string

func (g GermanGreeter) LanguageName() string {
	return "German"
}

func (g GermanGreeter) Greet(msg string) string {
	return "Hello " + msg + "!"
}

func main() {
	var germanGreeter GermanGreeter = ""
	fmt.Println(SayHello("Dietrich", germanGreeter))
}

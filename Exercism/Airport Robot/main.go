package main

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(name string, greeter Greeter) string {
	language := greeter.LanguageName()
	msg := greeter.Greet(name)
	return fmt.Sprintf("I can speak %s: %s", language, msg)
}

type German struct{}

func (g German) LanguageName() string {
	return "German"
}

func (g German) Greet(name string) string {
	return "Hallo " + name + "!"
}

type Italian struct {
}

func (g Italian) LanguageName() string {
	return "Italian"
}

func (g Italian) Greet(name string) string {
	return "Ciao " + name + "!"
}

type Portuguese struct {
}

func (g Portuguese) LanguageName() string {
	return "Portuguese"
}
func (g Portuguese) Greet(name string) string {

	return "Olá " + name + "!"
}

func main() {
	german := German{}
	fmt.Println(SayHello("Dietrich", german))
}

package main

import (
	"fmt"
	"log"

	"github.com/SedghiAmin/GO-exercises/Modules/greetings"
)

func main(){
	names:= []string{"Amin", "Ehsan", "Nilo"}
	messages, err:= greetings.Hellos(names)

	if err != nil{
		log.SetPrefix("greating: ")
		log.SetFlags(0)
		log.Fatal(err)
	}

	fmt.Println(messages)
}

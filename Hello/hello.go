package main

import (
	"fmt"
	"log"

	"github.com/SedghiAmin/GO-exercises/Modules/greetings"
)

func main(){
	
	message, err:= greetings.Hello("")

	if err != nil{
		log.SetPrefix("greating: ")
		log.SetFlags(0)
		log.Fatal(err)
	}
	
	fmt.Println(message)
}

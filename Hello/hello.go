package main

import (
	"fmt"

	"github.com/SedghiAmin/GO-exercises/Modules/greetings"
)

func main(){
	message:= greetings.Hello("Amin")
	fmt.Println(message)
}

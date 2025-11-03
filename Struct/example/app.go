package main

import (
	"fmt"
)

type Cars struct{
	name string
	model string
	color string
}

func main(){
	fmt.Println(Cars{"Benz", "C200", "Black"});
	
}
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
	var car = &Cars{"Benz", "C200", "Black"}
	fmt.Println(*car);
	(*car).color = "blue"
	fmt.Println("color's car changed to:", (*car).color)
}
package main

import (
	"fmt"
	"math/rand"
)
var c, basic, java bool
func add(x , y int) int{
	return x + y;
}

func swap(y string) (string , bool){
	java= true
	return y , java
}

func main() {
	fmt.Println("random number between 0 and 199: ", add(rand.Intn(100), rand.Intn(100)))
	a, b := swap("Amin")
	fmt.Println(a, b)
}
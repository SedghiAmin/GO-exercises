package main

import(
	"fmt"
)

func main(){

	var i , j = 12, 38
	fmt.Println(
		i , j,
	)
	var pi = &i
	var pj = &j

	*pi = 15
	*pj = 9
	fmt.Println(
		i , j,
	)

}
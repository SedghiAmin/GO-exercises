package main

import(
	"fmt"
)

func main(){ 

	var arr [3]string 

	arr[0] = "Ali"
	arr[1] = "Amin"
	arr[2] = "Ehsan"

	fmt.Println(arr)

	age := [5]int{65, 39, 38, 2, 94}

	fmt.Println(age[2:3])
}
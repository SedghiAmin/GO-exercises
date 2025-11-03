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

	age := []int{65, 39, 38, 2, 94}

	slice1 := age[1:3]
	slice2 := age[2:4]
	
	fmt.Println(age)
	fmt.Println(slice1 , slice2)

	slice2[0] = 0
	fmt.Println(age)
	fmt.Println(slice1 , slice2)

	type persons []struct{
		name string
		age int
	}

	person := persons{
		{"amin" , 39},
		{"ali" , 65},
	}

	fmt.Println(person)

}
package main

import "fmt"

func main() {
	var name string
	var age int
	var height float32

	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		return
	}

	fmt.Print("Enter your age: ")
	_, err = fmt.Scan(&age)
	if err != nil {
		return
	}

	fmt.Print("Enter your height: ")
	_, err = fmt.Scan(&height)
	if err != nil {
		return
	}

	//Print a blank line
	fmt.Println()

	//Show the details you typed
	fmt.Printf("Name is %s.     \n", name)
	fmt.Printf("Age is %d.      \n", age)
	fmt.Printf("Height is %.1f. \n", height)
}

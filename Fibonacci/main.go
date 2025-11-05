package main

import(
	"fmt"
)

func fibonacci(x int) int {
	
	if x <= 1 {
		return x
	}

	return fibonacci(x-1) + fibonacci(x-2)	

}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}
}
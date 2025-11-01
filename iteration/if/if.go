package main

import(
	"fmt"
	"math"
)

func sqrt(x float64) string{
	if x < 0 {
		return  fmt.Sprint(sqrt(-x) , "i") 
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(9), sqrt(-9))	
}

package main

import(
	"fmt"
	"math"
)

func pow(x, n ,lim float64) float64{
	if v := math.Pow(x , n); v < lim {
		return v
	}
	return  lim
}

func main() {
	fmt.Println(
		pow(5,2,27),
		pow(6,6,30),
	)
}
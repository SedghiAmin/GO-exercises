package main

import(
	"fmt"
	"math"
)

const DELTA = 0.0000001
const INITIAL_Z = 100.0

func sqrt(x float64) (float64){

	z := INITIAL_Z

	step := z - ((z * z - x) / (2 * z))

	for math.Abs(step - z) > DELTA{
		z = step
		step = z - ((z * z - x) / (2 * z))
	}

	return  z
}

func main(){
	
	fmt.Println(
		sqrt(9),
		math.Sqrt(9),
	)

}
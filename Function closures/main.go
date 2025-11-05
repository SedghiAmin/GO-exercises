package main

import(
	"fmt"
)

func adder() func(y int) int{
	sum :=0
	return func(x int) int{
		sum += x
		return sum
	}
}

func main(){
	first , sec := adder(), adder()

	for i:=range 10 {
		fmt.Println(
			first(i),
			sec((i * i) -i),
		)
	}
}
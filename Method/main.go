package main

import(
	"fmt"
)

type Dimention struct{
	x int
	y int
}

func (d Dimention) app() float64{
	return (float64(d.x) * float64(d.y))
}

func main(){
	d := Dimention{5 , 8}
	fmt.Println(
		d.app(),
	)
}
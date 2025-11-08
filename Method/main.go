package main

import(
	"fmt"
)

type Dimention struct{
	X int
	Y int
}

func (d Dimention) App() float64{
	return (float64(d.X) * float64(d.Y))
}

func (d *Dimention) Scale(s int){
	d.X = d.X * s
	d.Y = d.Y * s
}

func main(){
	d := Dimention{5 , 8}
	d.Scale(10)
	fmt.Println(
		d.App(),
	)
}
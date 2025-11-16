package main

import(
	"fmt"
)

func sum(s []int, c chan int){
	sum:= 0
	for _, v:= range s{
		sum += v
	}

	c <- sum	
}

func main(){
	slc:= []int{12, 5, 60, -5, 10, 3}
	c:= make(chan int)
	go sum(slc[: len(slc) / 2], c)
	go sum(slc[len(slc) / 2 :], c)

	x , y := <-c , <-c

	fmt.Println(x , y , x+y)
}
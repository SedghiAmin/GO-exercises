package main

import (
	"fmt"
)

type Dimantion struct{
	Weight int
	Height int
}

func main(){
	/*
	var m map[string]Dimantion
	m = make(map[string]Dimantion)
	*/
	m := make(map[string]Dimantion)
	m["window"] = Dimantion{Weight: 100, Height: 200}

	fmt.Println(m["window"]);
}
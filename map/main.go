package main

import (
	"fmt"
)

type Dimantion struct{
	Weight int
	Height int
}

func main(){
	var m = map[string]Dimantion{
		"window" : {100, 200},
		"wall" : {400, 700},
	}

	fmt.Println(m);
}
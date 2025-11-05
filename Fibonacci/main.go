package main

import(
	"fmt"
)

func fibonacci() func(int) int{

	cache := map[int] int{
		0 : 0,
		1 : 1,
	}

	var fibo func(int) int
	fibo = func(x int) int{
		if val , ok:= cache[x]; ok{
			return val
		}
		cache[x] = fibo(x-1) + fibo(x-2)
		return cache[x]
	}

	return fibo

}

func main(){
	f := fibonacci()
	for i:=0; i<10; i++{
		fmt.Println(
			f(i),
		)
	}
}
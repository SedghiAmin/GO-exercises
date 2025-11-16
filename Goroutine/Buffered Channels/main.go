package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)
	
	// Sender
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("I want send product:%d \n", i)
			ch <- i 
			fmt.Printf("product:%d sent \n", i)
		}
		close(ch)
	}()
	
	// Reciver
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		msg := <-ch
		fmt.Printf("Recived product:%d\n", msg)
	}
}
package main

import (
	"fmt"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	return t.Add(1000000000 * time.Second)
}

func main() {
	fmt.Println(AddGigasecond(time.Now()))
}

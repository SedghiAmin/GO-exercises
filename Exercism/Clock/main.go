package main

import "fmt"

type Clock struct {
	hour   int
	minute int
}

func normalize(h, m int) Clock {
	minutes := h*60 + m
	for minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{
		hour:   (minutes / 60) % 24,
		minute: minutes % 60,
	}
}

func New(h, m int) Clock {

	return normalize(h, m)
}

func (c Clock) Add(m int) Clock {
	return normalize(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return normalize(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func main() {
	c := New(12, 44)
	fmt.Println(c.Add(16))       // 13:00
	fmt.Println(c.Subtract(104)) // 11:00

	fmt.Println(New(23, 59).Add(1))    // 00:00
	fmt.Println(New(0, 0).Subtract(1)) // 23:59
	fmt.Println(New(25, 0))            // 01:00
	fmt.Println(New(-1, 15))           // 23:15
	fmt.Println(New(0, -1))            // 23:59
}

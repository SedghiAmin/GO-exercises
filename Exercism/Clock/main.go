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

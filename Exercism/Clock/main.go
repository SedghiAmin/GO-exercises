package main

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

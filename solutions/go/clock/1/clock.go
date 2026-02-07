package clock
import "fmt"
// Define the Clock type here.

type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	minutes := h*60 + m
	for minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{
		hour:   (minutes / 60) % 24,
		minute: minutes % 60,
	}
}

func (c Clock) Add(m int) Clock {
	minutes := c.hour*60 + c.minute + m
	return Clock{
		hour:   (minutes / 60) % 24,
		minute: minutes % 60,
	}
}

func (c Clock) Subtract(m int) Clock {
	minutes := c.hour*60 + c.minute - m
    for minutes < 0 {
		minutes += 24 * 60
	}
	return Clock{
		hour:   (minutes / 60) % 24,
		minute: minutes % 60,
	}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

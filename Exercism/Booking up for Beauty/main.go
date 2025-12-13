package main

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 01, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 02, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	hour := t.Hour()
	if hour >= 12 && hour < 18 {
		return true
	}
	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	weekDay := t.Weekday()
	month := t.Month()
	day := t.Day()
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", weekDay, month, day, year, hour, minute)
}

func main() {
	fmt.Println(Schedule("7/25/2019 13:45:00"))
	fmt.Println(HasPassed("July 25, 2019 13:45:00"))
	fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
	fmt.Println(Description("7/25/2019 13:45:00"))
}

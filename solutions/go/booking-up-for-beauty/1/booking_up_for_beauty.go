package booking

import( 
    "time"
    "fmt"
    )

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout:= "1/2/2006 15:04:05"
	t, _:= time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
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
	return fmt.Sprintf("You have an appointment on %s, at %s.",
		t.Format("Monday, January 2, 2006"), t.Format("15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}

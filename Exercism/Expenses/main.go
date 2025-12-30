package main

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Day1Records only returns true for records that are from day 1
func Day1Records(r Record) bool {
	return r.Day == 1
}

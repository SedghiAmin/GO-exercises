package main

import "fmt"

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

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	out := make([]Record, 0)
	for _, record := range in {
		if predicate(record) {
			out = append(out, record)
		}
	}
	return out
}

func main() {
	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
	}

	fmt.Printf("%+v", Filter(records, Day1Records))
}

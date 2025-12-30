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

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false

	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	total := 0.0
	for _, record := range in {
		if record.Day >= p.From && record.Day <= p.To {
			total += record.Amount
		}
	}
	return total
}

func main() {
	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	period := DaysPeriod{From: 1, To: 15}

	fmt.Printf("%+v", Filter(records, ByDaysPeriod(period)))
	fmt.Println("")
	fmt.Printf("%+v", Filter(records, ByCategory("groceries")))

	p1 := DaysPeriod{From: 1, To: 15}
	p2 := DaysPeriod{From: 16, To: 30}

	fmt.Println("")
	fmt.Printf("%+v", TotalByPeriod(records, p1))
	// => 16
	fmt.Println("")
	fmt.Printf("%+v", TotalByPeriod(records, p2))
	// => 50
}

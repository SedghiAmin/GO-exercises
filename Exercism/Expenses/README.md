# Expenses Analyzer

A Go program for analyzing and filtering expense records. This package provides utilities to filter expenses by date periods and categories, calculate totals, and generate financial reports.

## Features

- **Filter expenses** by date periods using `ByDaysPeriod`
- **Filter expenses** by category using `ByCategory`
- **Calculate total expenses** for a specific period with `TotalByPeriod`
- **Get category-specific expenses** within a time period using `CategoryExpenses`
- **Generic filtering** capabilities through the `Filter` function

## Data Structures

### Record
Represents an individual expense record:
```go
type Record struct {
    Day      int
    Amount   float64
    Category string
}
```

### DaysPeriod
Defines a range of days (inclusive):
```go
type DaysPeriod struct {
    From int
    To   int
}
```

## Functions

### Filter
```go
func Filter(in []Record, predicate func(Record) bool) []Record
```
Filters records based on a predicate function. Returns only records that satisfy the predicate.

### ByDaysPeriod
```go
func ByDaysPeriod(p DaysPeriod) func(Record) bool
```
Returns a predicate function that checks if a record falls within the specified date period.

### ByCategory
```go
func ByCategory(c string) func(Record) bool
```
Returns a predicate function that checks if a record belongs to the specified category.

### TotalByPeriod
```go
func TotalByPeriod(in []Record, p DaysPeriod) float64
```
Calculates the total amount of expenses within a given period.

### CategoryExpenses
```go
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error)
```
Calculates total expenses for a specific category within a period. Returns an error if the category doesn't exist in any records.

## Usage Example

```go
records := []Record{
    {Day: 1, Amount: 15, Category: "groceries"},
    {Day: 11, Amount: 300, Category: "utility-bills"},
    {Day: 12, Amount: 28, Category: "groceries"},
}

period := DaysPeriod{From: 1, To: 15}

// Filter records by period
filtered := Filter(records, ByDaysPeriod(period))

// Calculate total for period
total := TotalByPeriod(records, period)

// Get category expenses
amount, err := CategoryExpenses(records, period, "groceries")
```

## Error Handling

- `CategoryExpenses` returns an error when the specified category doesn't exist in any records
- All other functions handle edge cases gracefully (empty slices, zero-length periods, etc.)


## Design Notes

- Uses function composition for flexible filtering
- Pure functions with no side effects
- Clear separation of concerns
- Comprehensive error handling where appropriate
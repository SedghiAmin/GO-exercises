package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, ok := units[unit]
	if !ok {
		return false
	}

	value, _ := bill[item]
	bill[item] = value + units[unit]

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, ok := bill[item]
	if !ok {
		return false
	}
	_, ok = units[unit]
	if !ok {
		return false
	}
	bill[item] = value - units[unit]
	if bill[item] < 0 {
		bill[item] = value + units[unit]
		return false
	}
	if bill[item] == 0 {
		delete(bill, item)
		return true
	}
	return true
}

func main() {
	units := Units()
	fmt.Println(units)
	bill := NewBill()
	fmt.Println(bill)
	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok) // Output: true (since dozen is a valid unit)
	ok = RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok) // Output: true
}

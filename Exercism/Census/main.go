package main

import "fmt"

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{name, age, address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	switch {
	case r.Name == "" || r.Address == nil:
		return false
	}
	return true
}

func main() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	fmt.Println(NewResident(name, age, address))
	// => &{Matthew Sanabria 29 map[street:Main St.]}

	r := Resident{"", age, address}
	fmt.Println(r.HasRequiredInfo())
	// => false
}

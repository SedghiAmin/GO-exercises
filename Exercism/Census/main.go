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

func main() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	fmt.Println(NewResident(name, age, address))
	// => &{Matthew Sanabria 29 map[street:Main St.]}
}

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
	case r.Name == "" || r.Address["street"] == "":
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	sum := 0
	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			sum += 1
		}
	}
	return sum
}

func main() {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	resident1 := NewResident(name, age, address)
	fmt.Println(resident1)
	// => &{Matthew Sanabria 29 map[street:Main St.]}

	/*r := Resident{"", age, address}
	fmt.Println(r.HasRequiredInfo())
	// => false*/

	/*r.Delete()
	fmt.Println(r)
	// Output: &{ 0 map[]}*/

}

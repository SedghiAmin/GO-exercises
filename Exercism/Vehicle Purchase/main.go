package main

import (
	"fmt"
)

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	choice := option2
	if option1 < option2 {
		choice = option1
	}
	return fmt.Sprintf(`%s is clearly the better choice.`, choice)
}

func main() {
	println(NeedsLicense("car"))
	// => true

	println(NeedsLicense("bike"))
	// => false

	println(NeedsLicense("truck"))
	// => true

	println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
	// => "Toyota Corolla is clearly the better choice."

	println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	// => "Volkswagen Beetle is clearly the better choice."
}

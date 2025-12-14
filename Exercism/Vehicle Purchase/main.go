package main

import ()

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	}
	return false
}

func main() {
	println(NeedsLicense("car"))
	// => true

	println(NeedsLicense("bike"))
	// => false

	println(NeedsLicense("truck"))
	// => true
}

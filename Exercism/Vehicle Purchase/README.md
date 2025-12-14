# Vehicle Purchase Utilities
A Go package providing utility functions for vehicle purchase decisions, license requirements, and resell price calculations.

# Overview
This package offers three main functions to assist with vehicle-related decisions:

Determining if a vehicle type requires a license

Choosing between two vehicle options

Calculating resell prices based on vehicle age

# Functions
NeedsLicense(kind string) bool
Determines whether a driver's license is required for a specific type of vehicle.

# Parameters:

kind: Type of vehicle (string)

Returns: true if a license is needed, false otherwise

# Logic:

Only "car" and "truck" require a license

All other vehicle types (bike, scooter, etc.) don't require a license

# Examples:

``` go
NeedsLicense("car")    // returns true
NeedsLicense("truck")  // returns true
NeedsLicense("bike")   // returns false
NeedsLicense("boat")   // returns false
ChooseVehicle(option1, option2 string) string
```

Recommends a vehicle between two options based on lexicographical (alphabetical) order.

# Parameters:

option1: First vehicle option (string)

option2: Second vehicle option (string)

Returns: A formatted string recommending the vehicle that comes first alphabetically

# Examples:

``` go
ChooseVehicle("Wuling Hongguang", "Toyota Corolla")
// Returns: "Toyota Corolla is clearly the better choice."

ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf")
// Returns: "Volkswagen Beetle is clearly the better choice."
CalculateResellPrice(originalPrice, age float64) float64
Calculates the resell price of a vehicle based on its age and original price.
```

# Parameters:

originalPrice: Original purchase price of the vehicle (float64)

age: Age of the vehicle in years (float64)

Returns: Resell price as float64

Pricing Rules:

Age < 3 years: 80% of original price

3 ≤ Age < 10 years: 70% of original price

Age ≥ 10 years: 50% of original price

# Examples:

``` go
CalculateResellPrice(1000, 1)   // returns 800
CalculateResellPrice(1000, 5)   // returns 700
CalculateResellPrice(1000, 15)  // returns 500
```
# Dependencies
Go 1.16 or higher

Standard library only (no external dependencies)

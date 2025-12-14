# Car Production Calculator
A Go package for calculating car production metrics and costs in a manufacturing assembly line.

# Functions
CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64
Calculates the number of working cars produced per hour based on the production rate and success rate.

# Parameters:

productionRate: Total cars produced per hour

successRate: Percentage of successfully produced cars (0-100)

Returns: Number of working cars per hour as a float64

# Example:

```go
CalculateWorkingCarsPerHour(1547, 90) // Returns 1392.3
CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int
```
Calculates the number of working cars produced per minute based on the production rate and success rate.

# Parameters:

productionRate: Total cars produced per hour

successRate: Percentage of successfully produced cars (0-100)

Returns: Number of working cars per minute as an integer (rounded down)

# Example:

``` go
CalculateWorkingCarsPerMinute(1105, 90) // Returns 16
CalculateCost(carsCount int) uint
```
Calculates the total production cost for a given number of cars. The cost structure is:

10 cars cost $95,000 ($9,500 per car)

Individual cars cost $10,000 each

# Parameters:

carsCount: Total number of cars to produce

Returns: Total production cost as an unsigned integer

Example:

```go
CalculateCost(21) // Returns 200,000
```
// (2 groups of 10 = $190,000 + 1 individual = $10,000)
# Usage
```go
package main

import (
"fmt"
)

func main() {
// Calculate hourly production
hourly := CalculateWorkingCarsPerHour(1547, 90)
fmt.Println(hourly) // Output: 1392.3

    // Calculate minute production
    minute := CalculateWorkingCarsPerMinute(1105, 90)
    fmt.Println(minute) // Output: 16
    
    // Calculate production cost
    cost := CalculateCost(21)
    fmt.Println(cost) // Output: 200000
}
```

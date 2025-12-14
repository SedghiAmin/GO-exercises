package main

import (
	"fmt"
)

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))

}

// Package weather provides weather forecasting functionality.
package weather

var (
    // CurrentCondition holds the current weather condition.
    CurrentCondition string
    // CurrentLocation holds the current city name.
    CurrentLocation string
)

// Forecast returns a weather forecast for the given city and condition.
func Forecast(city, condition string) string {
    CurrentLocation, CurrentCondition = city, condition
    return CurrentLocation + " - current weather condition: " + CurrentCondition
}
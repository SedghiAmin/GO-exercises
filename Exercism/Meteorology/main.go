package main

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string {
	return fmt.Sprintf("%d %s", t.degree, t.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string {
	return fmt.Sprintf("%d %v", s.magnitude, s.unit)
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (m MeteorologyData) String() string {
	return fmt.Sprintf("%s: %d %s, Wind %s at %d %s, %d%% Humidity", m.location, m.temperature.degree, m.temperature.unit, m.windDirection, m.windSpeed.magnitude, m.windSpeed.unit, m.humidity)
}

func main() {
	// تست ۱: Athens
	athensData := MeteorologyData{
		location:      "Athens",
		temperature:   Temperature{21, Celsius},
		windDirection: "N",
		windSpeed:     Speed{16, KmPerHour},
		humidity:      63,
	}

	result := athensData.String()
	expected := "Athens: 21 °C, Wind N at 16 km/h, 63% Humidity"

	fmt.Printf("Result:   %q\n", result)
	fmt.Printf("Expected: %q\n", expected)
	fmt.Printf("Match:    %v\n", result == expected)

	// تست ۲: San Francisco
	sfData := MeteorologyData{
		location:      "San Francisco",
		temperature:   Temperature{57, Fahrenheit},
		windDirection: "NW",
		windSpeed:     Speed{19, MilesPerHour},
		humidity:      60,
	}

	fmt.Println("\nSan Francisco test:")
	fmt.Println(sfData.String())
}

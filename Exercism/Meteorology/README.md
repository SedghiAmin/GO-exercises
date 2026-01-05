# Meteorology Package

A Go program for weather data representation with formatted output.

## Quick Start

```go
package main

import "fmt"

func main() {
    data := MeteorologyData{
        location:      "Athens",
        temperature:   Temperature{21, Celsius},
        windDirection: "N",
        windSpeed:     Speed{16, KmPerHour},
        humidity:      63,
    }
    
    fmt.Println(data) // Athens: 21 °C, Wind N at 16 km/h, 63% Humidity
}
```

## Types

### Temperature
```go
temp := Temperature{25, Celsius}
fmt.Println(temp) // 25 °C
```

### Speed
```go
wind := Speed{30, KmPerHour}
fmt.Println(wind) // 30 km/h
```

### Complete Weather Data
```go
weather := MeteorologyData{
    location:    "Tokyo",
    temperature: Temperature{18, Celsius},
    windSpeed:   Speed{12, KmPerHour},
    humidity:    65,
}
```

## Format
All types implement `String()` for clean display:
- Temperature: `"25 °C"`
- Speed: `"30 km/h"`
- Full data: `"Location: 25 °C, Wind NE at 30 km/h, 65% Humidity"`

Simple, type-safe weather data handling in Go.
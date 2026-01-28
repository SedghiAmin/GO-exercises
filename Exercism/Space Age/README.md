# Space Age

A simple Go program for calculating a person's age on different planets of the solar system.

## Description

The `Age` function converts a person's age given in seconds on Earth into the equivalent age on a specific target planet. The conversion is based on the orbital period (the length of a year) of that planet relative to Earth.

## Usage

```go
package main

import (
    "fmt"
    "yourmodule/space"
)

func main() {
    // Example: A person who has lived for 1,000,000,000 seconds on Earth
    ageOnEarth := space.Age(1000000000, "Earth")
    ageOnMars := space.Age(1000000000, "Mars")
    
    fmt.Printf("Age on Earth: %.2f years\n", ageOnEarth)
    fmt.Printf("Age on Mars: %.2f years\n", ageOnMars)
}
```

## Parameters

*   `seconds` (type: `float64`): The number of seconds the person has lived on Earth.
*   `planet` (type: `Planet` which is a `string`): The name of the target planet for age calculation.
    *   Valid values:
        *   `"Mercury"`
        *   `"Venus"`
        *   `"Earth"`
        *   `"Mars"`
        *   `"Jupiter"`
        *   `"Saturn"`
        *   `"Uranus"`
        *   `"Neptune"`

## Return Value

The function returns a `float64` representing the person's age in orbital years on the specified planet.

**Note:** If a planet name other than the valid ones above is provided, the function returns `-1.00`.

## Prerequisites

*   Go (version 1.13 or higher recommended)

## Installation & Running

1.  Copy the code file (e.g., `space.go`) and this `README.md` into your project directory.
2.  Ensure your package is properly set up (e.g., `module` in `go.mod`).
3.  Call the function in your code as shown in the example above.

## Calculation Explanation

The calculation is based on the **Earth orbital period** (31557600 seconds). For each planet, this value is divided by the ratio of that planet's orbital period to Earth's:
*   `Age on Planet = (input seconds) / (Earth orbital seconds) / (planet's orbital ratio)`

## Example Output

```bash
# For input Age(1000000000, "Earth")
Age on Earth: 31.69 years

# For input Age(1000000000, "Mars")
Age on Mars: 16.85 years
```
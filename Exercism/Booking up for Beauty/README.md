# Appointment Date Utilities
A Go package providing utility functions for parsing, checking, and formatting appointment dates.

# Overview
This package offers a set of date/time utilities specifically designed for handling appointment scheduling and checking. It handles various date formats and provides common operations needed for appointment management systems.

# Functions
Schedule(date string) time.Time
Parses a date string in "M/D/YYYY HH:MM:SS" format and returns a time.Time object.

# Example:

```go
Schedule("7/25/2019 13:45:00")
HasPassed(date string) bool
```
Checks if a given appointment date (in "January 2, 2006 15:04:05" format) has already passed relative to the current time.

# Example:

``` go
HasPassed("July 25, 2019 13:45:00")
IsAfternoonAppointment(date string) bool
```
Determines if an appointment (in "Monday, January 2, 2006 15:04:05" format) is scheduled for the afternoon (12:00-17:59).

# Example:

``` go
IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00")
Description(date string) string
```
Formats an appointment date into a human-readable description string.

# Example Output:

text
"You have an appointment on Thursday, July 25, 2019, at 13:45."
AnniversaryDate() time.Time
Returns this year's anniversary date (September 15) at midnight UTC.

# Usage
```go
package main

import (
"fmt"
// Import the package
)

func main() {
// Parse a scheduled date
appointment := Schedule("7/25/2019 13:45:00")
fmt.Println(appointment)

    // Check if appointment has passed
    fmt.Println(HasPassed("July 25, 2019 13:45:00"))
    
    // Check if it's an afternoon appointment
    fmt.Println(IsAfternoonAppointment("Thursday, July 25, 2019 13:45:00"))
    
    // Get appointment description
    fmt.Println(Description("7/25/2019 13:45:00"))
    
    // Get this year's anniversary date
    fmt.Println(AnniversaryDate())
}
```
# Purpose
This package simplifies common date operations for appointment scheduling systems, providing consistent parsing and formatting across different date representations commonly used in appointment management applications.


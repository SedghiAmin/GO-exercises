# Time Parser with Custom Format
A Go utility that converts human-readable date format patterns into Go's time layout format for parsing date strings.

# Overview
This package provides a ParseWithFormat function that allows you to parse date strings using intuitive format patterns (like "ddd, mm/dd/yyyy, HH:MM") instead of Go's numeric reference layout ("Mon, 01/02/2006, 15:04").

# Key Features
Human-readable format patterns: Use yyyy for year, mm for month, dd for day, etc.

Simple API: Just provide the date string and your format pattern.

Common format support: Handles day names, 12/24 hour time, and date components.

# Format Specifiers

Pattern	Represents	Go Equivalent
yyyy	4-digit year	2006
yy	    2-digit year	06
mm	    Month (01-12)	01
dd	    Day of month	02
ddd	    Abbreviated weekday	Mon
dddd	Full weekday	Monday
HH	    Hour (00-23)	15
MM	    Minute (00-59)	04
SS	    Second (00-59)	05

#Example
```go
t, _ := ParseWithFormat("Tue, 09/22/1995, 13:00", "ddd, mm/dd/yyyy, HH:MM")
fmt.Println(t) // Output: 1995-09-22 13:00:00 +0000 UTC
```

This makes date parsing in Go more intuitive and readable, especially for developers familiar with format patterns from other programming languages.


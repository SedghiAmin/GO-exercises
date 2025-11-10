# IP Address Stringer Utility
A simple Go utility for beautifully formatted IP address display using Go interfaces.

# Features
Automatic IP address display in x.x.x.x format

Custom stringer implementation for IPAddr type

Support for hostname to IP address mapping

Clean and extensible code

# Code Example
```go
package main

import "fmt"

// IPAddr custom type for representing IP addresses
type IPAddr [4]byte

// String method implements fmt.Stringer interface
func (ip IPAddr) String() string {
    return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
        "local":     {192, 168, 1, 1},
        "cloud":     {10, 0, 0, 1},
    }
    
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
```
# Sample Output
text
loopback: 127.0.0.1
googleDNS: 8.8.8.8
local: 192.168.1.1
cloud: 10.0.0.1

# Key Benefits
Clean Code: Uses standard Go interfaces

Extensible: Easy to add new types by implementing String()

Practical: Provides readable and standard IP address display

Educational: Great for understanding Go interfaces

# Learning Objectives
This project is perfect for learning:

Implementing standard Go interfaces

Working with custom types in Go

Using maps and data structures

Formatting output with fmt package

Method receivers in Go

# Advanced Usage
Adding Custom IP Ranges
```go
// Extend with custom networks
var privateNetworks = map[string]IPAddr{
    "router":    {192, 168, 1, 1},
    "nas":       {192, 168, 1, 10},
    "printer":   {192, 168, 1, 20},
}
```
# Batch Processing
```go
func printAllIPs(hosts map[string]IPAddr) {
    for host, ip := range hosts {
        fmt.Printf("Host: %-15s IP: %s\n", host, ip)
    }
}
```
# Getting Started
Prerequisites
Go 1.16 or higher

Basic understanding of Go syntax

# Run with verbose output
go run -v main.go

# Build for specific platform
GOOS=linux GOARCH=amd64 go build
📊 Performance
The implementation is optimized for:

Minimal memory allocation

Fast string conversion

Efficient map iterations

# Future Enhancements
IPv6 support

Network validation

Subnet calculations

DNS lookup integration

JSON serialization

# License
This project is licensed under the MIT License - see the LICENSE file for details.

# Contact & Support
GitHub: Amin Sedghi

Email: sedghi.amin@gmail.com

Issues: GitHub Issues

# Acknowledgments
Go standard library documentation

Go community best practices

Contributors and testers

Happy Coding! 🎉
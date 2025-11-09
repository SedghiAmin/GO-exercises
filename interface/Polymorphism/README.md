# Shape Interface Demo
A sample Go project demonstrating the concept of interfaces and polymorphism.

# Overview
This project defines a Shape interface that is implemented by two different geometric types (Rectangle and Circle). This example demonstrates how to use interfaces in Go to create modular and flexible code.

# Structure
Interface Definition
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

# Code Components
Main Functions
Area(): Calculates the area of the shape

Perimeter(): Calculates the perimeter/circumference of the shape

PrintShapeInfo(): Utility function to display shape information

main(): Demonstrates usage with a slice of different shapes

# Usage
Prerequisites
Go 1.16 or higher

Running the Program
bash

# Clone or download the file
go run main.go

# Key Concepts Demonstrated
Interfaces: Defining contracts that types can implement

Polymorphism: Treating different types uniformly through interfaces

Method Implementation: Implementing interface methods on struct types

Slice of Interfaces: Storing different types in a single collection


# Learning Outcomes
This example helps understand:

How to define and implement interfaces in Go

The implicit interface implementation in Go

Using interfaces for polymorphic behavior

Writing flexible and extensible code with interfaces

# Extending the Project
You can easily add new shapes by:

Creating a new struct type

Implementing the Area() and Perimeter() methods

Adding instances to the shapes slice

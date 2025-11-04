# Student Management System

A simple Go program for managing student information including personal details and educational records.

## Description

This program demonstrates basic Go concepts including:
- Struct definitions and composition
- Slice operations
- Iterating over collections
- Formatted output

## Features

- Store student personal information (name, age)
- Store student educational information (degree, grade)
- Display all students in a formatted list
- Add new students dynamically
- Track total number of students

## Data Structures

### Person
Stores personal information:
- `Name` (string): Student's full name
- `Age` (int): Student's age

### Education
Stores educational information:
- `Degree` (string): Educational degree/level
- `Grade` (float64): Academic grade/GPA

### Student
Combines personal and educational information:
- `PersonalInfo` (Person): Personal details
- `EductionInfo` (Education): Educational details

## Usage

### Prerequisites
- Go 1.x or higher installed on your system

### Running the Program

```bash
go run app.go
```

### Expected Output

```
List of students; 2 student 
==================
Student 1:
  Full Name: Amin Sedghi
  Age: 39
  Degree: bachlor
  Grade: 19.00
------------------
Student 2:
  Full Name: Ahmad Nazari
  Age: 32
  Degree: bachlor
  Grade: 14.50
------------------
One student added, all of students are 3 person.
```

## Code Example

```go
// Creating a new student
newStudent := Student{
    PersonalInfo: Person{Name: "John Doe", Age: 25},
    EductionInfo: Education{Degree: "Bachelor", Grade: 16.5},
}

// Adding to the list
students = append(students, newStudent)
```

## Learning Objectives

This code is useful for learning:
- Go struct composition
- Working with slices
- Range loops
- String formatting with `fmt.Printf`
- Basic CRUD operations (Create, Read)


## License

This is a learning/educational project and is free to use and modify.

## Author

Amin Sedghi
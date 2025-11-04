package main

import(
	"fmt"
)

type Person struct{
	Name string
	Age int
}

type Education struct{
	Degree string
	Grade float64
}

type Student struct{
	PersonalInfo Person
	EductionInfo Education
}

func main(){
	var students = []Student{
		{
			PersonalInfo: Person{Name: "Amin Sedghi", Age: 39,},
			EductionInfo : Education{Degree: "bachlor",	Grade: 19,},
		},
		{
			PersonalInfo: Person{Name: "Ahmad Nazari", Age: 32,},
			EductionInfo : Education{Degree: "bachlor",	Grade: 14.5,},
		},
	}

	// نمایش اطلاعات دانش‌آموزان
	fmt.Printf("List of students; %d student \n", len(students))
	fmt.Println("==================")
	for i, student := range students {
		fmt.Printf("Student %d:\n", i+1)
		fmt.Printf("  Full Name: %s\n", student.PersonalInfo.Name)
		fmt.Printf("  Age: %d\n", student.PersonalInfo.Age)
		fmt.Printf("  Degree: %s\n", student.EductionInfo.Degree)
		fmt.Printf("  Grade: %.2f\n", student.EductionInfo.Grade)
		fmt.Println("------------------")
	}

	newStudent := Student{
		PersonalInfo: Person{Name: "morteza askari", Age: 25},
		EductionInfo: Education{Degree: "Diplom", Grade: 12.68},
	}

	students = append(students, newStudent)
	fmt.Printf("One student added, all of students are %d person.\n", len(students))
}
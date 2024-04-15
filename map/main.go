package main

import "fmt"

func main() {
	studentGrades := make(map[string]int)

	studentGrades["John"] = 90
	studentGrades["Jane"] = 80
	studentGrades["Doe"] = 70

	for student, grade := range studentGrades {
		fmt.Println(student, grade)
	}

	// Accessing a value
	fmt.Println(studentGrades["John"])

	// Deleting a value
	delete(studentGrades, "John")
	fmt.Println(studentGrades)

	// Check if a key exists
	grade, exists := studentGrades["John"]
	if exists {
		fmt.Println("Grade of John is", grade)
	} else {
		fmt.Println("John's grade does not exist")
	}
	
	myMarks := map[string]int{
		"Maths": 90,
		"Science": 80,
		"English": 70,
	}

	fmt.Println(myMarks)

}

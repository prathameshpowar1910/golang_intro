package main

import "fmt"

func main() {
	numbers := []int{39, 12, 93, 29, 59}
	fmt.Println(numbers)
	fmt.Println("Length of numbers slice is: ", len(numbers))
	numbers = append(numbers, 100)
	fmt.Println(numbers)

	name := []string{}
	name = append(name, "Hello")
	fmt.Println(name)

	marks := make([]int, 4, 10)
	marks[0] = 39
	marks[1] = 12
	marks[2] = 93
	fmt.Println(marks)
	fmt.Println("Length of marks slice is: ", len(marks))
	fmt.Println("Capacity of marks slice is: ", cap(marks))


}

//? slices are dynamic arrays. They are similar to arrays but their size can be changed at runtime.
//? Slices are more powerful than arrays. They are more flexible and convenient to use.
//? Slices are reference types. When you assign a slice to another slice, both of them point to the same underlying array.
//? Slices are created using the make() function. The make() function allocates a zeroed array and returns a slice that refers to that array.
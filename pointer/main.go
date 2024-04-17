package main

import "fmt"

func modifyValueByReference(ptr *int) {
	*ptr = *ptr + 10
}

func main() {

	// a ponter is a variable that stores the memory address of another variable.
	// In Go, pointers are used to store the memory address of another value.
	// The type *T is a pointer to a T value. Its zero value is nil.
	var p *int // The & operator generates a pointer to its operand.
	i := 42
	p = &i
	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
	//? default value of pointer is nil
	var num int = 11
	var ptr *int = &num
	fmt.Println("Value of num is: ", num)
	fmt.Println("Value of num is: ", *ptr)
	fmt.Println("Address of num is: ", &num)
	fmt.Println("Address of num is: ", ptr)

	//? passing pointer to a function
	value := 10
	fmt.Println("Value before modification: ", value)
	modifyValueByReference(&value)
	fmt.Println("Value after modification: ", value)
}

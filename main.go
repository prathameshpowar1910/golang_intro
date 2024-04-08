package main

import (
	"fmt"
	"mylearning/myutil"
)

func PublicFunction() {
	fmt.Println("Public Function")
}

func privateFunction() {
	fmt.Println("Private Function")
}

func main() {
	fmt.Println("Hello, World!")
	myutil.PrintMessage("Hello, World! from myutil")

	//? Variable Declaration
	var name string = "Prathamesh"
	var surname = "Powar"
	fmt.Println("Hello,", name, surname)

	var age int = 21
	var birthYear = 2003
	fmt.Println("Age:", age)
	fmt.Println("Birth Year:", birthYear)

	var dimension float64 = 3.14
	fmt.Println("Dimension:", dimension)

	var isTrue bool = true
	fmt.Println("Is True:", isTrue)

	var a, b int = 1, 2
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	const pi = 3.142
	fmt.Println("Pi:", pi)
	//! pi = 3.14 Error: cannot assign to pi

	person := "Prathamesh"
	fmt.Println("Person:", person)

	var Public = "Public"   //? Public is exported
	var private = "Private" //? private is not exported

	fmt.Println("Public:", Public)
	fmt.Println("Private:", private)

	//* this same is applicable for functions, variables, constants, etc.
	PublicFunction()
	privateFunction()

	//? Multiple variable declaration
	var (
		firstName = "Prathamesh"
		lastName  = "Powar"
	)
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)

	//? Multiple variable declaration with type
	var (
		age1 int = 21
		age2 int = 22
	)
	fmt.Println("Age 1:", age1)
	fmt.Println("Age 2:", age2)
}

package main

import "fmt"

func main() {
	age := 35
	name := "John"
	height := 180.5

	fmt.Printf("Name: %s Age: %d Height: %.1f\n", name, age, height)
	fmt.Println("Name:", name, "Age:", age, "Height:", height)
	fmt.Printf("Type of name: %T\n", name)
}

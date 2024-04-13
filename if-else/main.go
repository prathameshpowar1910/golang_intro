package main

import "fmt"

func main() {
	x := 10
	if x == 0 {
		fmt.Println("x is zero")
	} else if x > 0 {
		fmt.Println("x is positive")
	} else {
		fmt.Println("x is negative")
	}

	y := 20
	if x == 10 && y == 20 {
		fmt.Println("x is 10 and y is 20")
	} else {
		fmt.Println("x is not 10 or y is not 20")
	}
}
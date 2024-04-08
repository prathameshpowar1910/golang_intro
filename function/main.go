package main

import "fmt"

func simpleFunction() {
	fmt.Println("This is a simple function")
}

func add(a,b int) int {
	return a+b
}

func multiply(a,b int) (res int) {
	res = a*b
	return
}

func main() {
	fmt.Println("Learning GoLang: Functions")
	simpleFunction()
	fmt.Println(add(1,2))
	data := multiply(2,3)
	fmt.Println("multiplication of 2 and 3 is:", data)
}
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 11
	fmt.Println("Value of num is: ", num)
	fmt.Printf("Type of num is %T", num)
	// num = num + 1.23 //? cannot use num + 1.23 (type float64) as type int in assignment
	var data float64 = float64(num)
	data = data + 1.23 //? no error
	fmt.Println("Value of data is: ", data)
	fmt.Printf("Type of data is %T", data)
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Value of str is: ", str)
	fmt.Printf("Type of str is %T", str)
	number_string := "123"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("Value of number_int is: ", number_int)
	fmt.Printf("Type of number_int is %T", number_int)

	num_str := "3.14"
	num_float, _ := strconv.ParseFloat(num_str, 64)
	fmt.Println("Value of num_float is: ", num_float)
	fmt.Printf("Type of num_float is %T", num_float)
}
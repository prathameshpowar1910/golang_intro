package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		// panic("Division by zero")
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	ans1, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Division is: ", ans1)
	
	ans2, _ := divide(10, 4)
	fmt.Println("Division is: ", ans2)
}

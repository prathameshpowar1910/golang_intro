package main

import "fmt"

func main() {
	// for loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	// while loop
	j := 0
	for j < 5 {
		println(j)
		j++
	}

	// infinite loop
	k := 0
	for {
		println(k)
		k++
		if k == 5 {
			break
		}
	}

	// range loop
	numbers := []int{1, 2, 3, 4, 5}
	for index, number := range numbers {
		println(index, number)
	}

	data := "Hello, World!"
	for index, letter := range data {
		fmt.Printf("Index of data is %d and letter is %c\n", index, letter)
	}
}

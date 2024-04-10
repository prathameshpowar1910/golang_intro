package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "Hello"
	name[1] = "World"
	name[2] = "How"
	name[3] = "Are"
	name[4] = "You"

	fmt.Println(name)

	var numbers = [5]int{39,12,93,29,59}
	fmt.Println(numbers)

	var length = len(numbers)
	fmt.Println("Length of numbers array is: ", length)

	fmt.Println("First element of numbers array is: ", numbers[0])

	var names[5]string
	fmt.Println("Empty names array: ", names)
	fmt.Printf("Names with qutotaion: %q", names)
}
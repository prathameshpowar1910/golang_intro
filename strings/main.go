package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts) // [apple orange banana]

	str := "one two three four five two one four"
	count := strings.Count(str, "two")
	fmt.Println(count) // 2

	str = "    Hello, World!    "
	str = strings.TrimSpace(str)
	fmt.Println(str) // Hello, World!

	str = "Hello, World!"
	str = strings.Replace(str, "World", "Go", 1)
	fmt.Println(str) // Hello, Go!

	str = "Hello, World!"
	str = strings.ToUpper(str)
	fmt.Println(str) // HELLO, WORLD!

	firstName := "John"
	lastName := "Doe"
	fullName := strings.Join([]string{firstName, lastName, "chaman"}, " + ")
	fmt.Println(fullName) // John + Doe + chaman
}

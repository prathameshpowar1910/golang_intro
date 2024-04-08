package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name:")
	//? using fmt package to get user input
	// var name string
	// fmt.Scanln(&name)
	// fmt.Printf("Hello, %s!\n", name)
	
	//? using bufio package to get user input
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s!\n", name)
}

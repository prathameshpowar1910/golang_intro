package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Hello ended")
}

func sayHi() {
	fmt.Println("Hi")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hi ended")
}

func main() {
	fmt.Println("Main execution started")
	go sayHello() // This will run in a separate goroutine
	go sayHi()

	time.Sleep(1000 * time.Millisecond) // This is to wait for the goroutine to finish
}

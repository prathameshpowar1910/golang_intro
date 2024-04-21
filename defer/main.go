package main

import "fmt"

func add(a,b int) int {
	return a+b
}

func main() {
	//? defer keyword is used to delay the execution of a statement until the surrounding function returns.
	//? The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	//? Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	//? Deferred functions may read and assign to the returning function's named return values.
	//? The defer statement is often used to close a file, so the file is closed as soon as the surrounding function returns.
	//? The file is closed after the surrounding function returns, regardless of the return path taken.
	//? The defer statement is also used to unlock a mutex, so the mutex is unlocked as soon as the surrounding function returns.
	//? The defer statement is also used to decrement a counter, so the counter is decremented as soon as the surrounding function returns.
	//? The defer statement is also used to print a log message, so the log message is printed as soon as the surrounding function returns.
	//? The defer statement is also used to recover from a panic, so the panic is recovered as soon as the surrounding function returns.
	//? The defer statement is also used to measure the time taken by a function, so the time is measured as soon as the surrounding function returns.
	//? The defer statement is also used to release a resource, so the resource is released as soon as the surrounding function returns.
	//? The defer statement is also used to clean up a resource, so the resource is cleaned up as soon as the surrounding function returns.

	//? Example 1

	fmt.Println("start")
	data := add(5,6)
	defer fmt.Println("data is ",data)
	defer fmt.Println("middle")
	fmt.Println("end")
}
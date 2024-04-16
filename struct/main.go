package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email   string
	PhoneNo string
}

type Address struct {
	HouseNo string
	Street  string
	City    string
	State   string
}

type Employee struct {
	Person_Details Person
	Person_contact Contact
	Person_address Address
	EmployeeID     int
}

//* Another way to define struct
// type Employee2 struct {
// 	Person
// 	Contact
// 	Address
// 	EmployeeID int
// }

func main() {
	//* first method
	var user1 Person
	fmt.Println(user1)

	user1.FirstName = "John"
	user1.LastName = "Doe"
	user1.Age = 30

	fmt.Println(user1)

	//* second method
	user2 := Person{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       25,
	}

	fmt.Println(user2)

	//* new keyword

	var user3 = new(Person)
	user3.FirstName = "Tom"
	user3.LastName = "Jerry"
	user3.Age = 20

	fmt.Println(user3)

	employee1 := Employee{
		Person_Details: Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		},
		Person_contact: Contact{
			Email:   "i@gmail.com",
			PhoneNo: "1234567890",
		},
		Person_address: Address{
			HouseNo: "123",
			Street:  "Main Street",
			City:    "City",
			State:   "State",
		},
	}

	fmt.Println(employee1)
}

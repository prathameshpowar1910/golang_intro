package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	person := Person{Name: "John", Age: 30, IsAdult: true}

	//convert person struct to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON data: ", string(jsonData))

	//decoding or unmarshalling JSON data to struct
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		panic(err)
	}
	fmt.Println("Name: ", newPerson)
}

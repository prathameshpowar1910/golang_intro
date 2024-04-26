package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("Error: status code not OK")
	}

	//? old way to read data from response
	// data, err := ioutil.ReadAll(res.Body) // ReadAll reads from res.Body until an error or EOF and returns the data it read
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Data :", string(data))

	//? usin unmarshal
	// var todos Todo
	// err = json.Unmarshal(data, &todos)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Data :", todos)

	//? new way to read data from response
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		panic(err)
	}
	// fmt.Println("Data :", todo)

	fmt.Println("User ID :", todo.UserId)
	fmt.Println("ID :", todo.Id)
	fmt.Println("Title :", todo.Title)
	fmt.Println("Completed :", todo.Completed)
}

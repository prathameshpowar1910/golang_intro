package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
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

func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Id:        1,
		Title:     "New Todo",
		Completed: false,
	}

	//convert struct to json data
	jsonData, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}

	//convert json data to string
	jsonString := string(jsonData)
	fmt.Println("JSON Data :", jsonString)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	myURL := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
}

func preformUpdateRequest() {
	todo := Todo{
		UserId:    23,
		Id:        1,
		Title:     "Changed Todo",
		Completed: false,
	}

	//convert struct to json data
	jsonData, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}

	//convert json data to string
	jsonString := string(jsonData)
	fmt.Println("JSON Data :", jsonString)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myURL, jsonReader)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Status Code :", res.Status)
}

func main() {
	//?perform get request
	// performGetRequest()

	//?perform post request
	// performPostRequest()

	//?perform update request
	preformUpdateRequest()
}

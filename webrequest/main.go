package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	result, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	defer result.Body.Close()
	fmt.Printf("Type of response is %T \n", result)

	//read the response
	data, er := ioutil.ReadAll(result.Body)
	if er != nil {
		log.Fatalf("Error: %v", er)
		return
	}

	fmt.Println("response is ", string(data))
}

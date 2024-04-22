package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	log.Fatal("Error is :",err)
	// 	return
	// } 

	// defer file.Close()
	// content := "Hello, World!"
	// byte, error := io.WriteString(file, content)
	
	// fmt.Println("Bytes written are: ", byte)

	// if error != nil {
	// 	log.Fatal("Error is :",err)
	// 	return
	// }
	// fmt.Println("File written successfully")

	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	log.Fatal("Error is :",err)
	// 	return
	// }
	// defer file.Close()
	// //create a buffer to read the file content
	// buffer := make([]byte,1024)

	// //read the file content into the buffer
	// for {
	// 	byteRead, err := file.Read(buffer)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal("Error is :",err)
	// 		return
	// 	}
	// 	fmt.Print(string(buffer[:byteRead]))

	// 	//process the read content
	// 	fmt.Println(string(buffer[:byteRead]))
	// }

	content,err := ioutil.ReadFile("example.txt") // this is deprecated as it reads the whole file at once and memory is not efficient
	if err != nil {
		log.Fatal("Error is :",err)
		return
	}
	fmt.Println("File content is: ",string(content))

	contents,error := os.ReadFile("example.txt") // this is the new way to read the file
	if error != nil {
		log.Fatal("Error is :",error)
		return
	}
	fmt.Println("File content is: ",string(contents))
}
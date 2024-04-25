package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://example.com:8080/path/to/something?name=abc&age=20#fragement"
	fmt.Printf("My URL is %s\n", myUrl)
	fmt.Printf("Type of myUrl is %T\n", myUrl)

	parsedUrl, err := url.Parse(myUrl)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Scheme: %s\n", parsedUrl.Scheme)
	fmt.Printf("Host: %s\n", parsedUrl.Host)
	fmt.Printf("Path: %s\n", parsedUrl.Path)
	fmt.Printf("RawQuery: %s\n", parsedUrl.RawQuery)
	fmt.Printf("Fragment: %s\n", parsedUrl.Fragment)
	fmt.Printf("User: %s\n", parsedUrl.User)
	fmt.Printf("Opaque: %s\n", parsedUrl.Opaque)
	fmt.Printf("ForceQuery: %t\n", parsedUrl.ForceQuery)
	fmt.Printf("RawPath: %s\n", parsedUrl.RawPath)
	fmt.Printf("RequestURI: %s\n", parsedUrl.RequestURI())
	fmt.Printf("String: %s\n", parsedUrl.String())
	fmt.Printf("Hostname: %s\n", parsedUrl.Hostname())
	fmt.Printf("Port: %s\n", parsedUrl.Port())
	fmt.Printf("IsAbs: %t\n", parsedUrl.IsAbs())

	//modify the URL
	parsedUrl.Path = "/newpath"
	parsedUrl.RawQuery = "name=def&age=30"

	newUrl := parsedUrl.String()
	fmt.Printf("New URL: %s\n", newUrl)

}
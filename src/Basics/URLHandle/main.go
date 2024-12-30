package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("LEarning URL")

	MyURL := "https://example.com:8080/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of URL: %T\n", MyURL)

	ParseURL, err := url.Parse(MyURL)
	if err != nil {
		fmt.Println("Can't parse url: ", err)
		return
	}

	fmt.Printf("Type of URL: %T\n", ParseURL)

	fmt.Println("Scheme: ", ParseURL.Scheme)
	fmt.Println("Scheme: ", ParseURL.Host)
	fmt.Println("Scheme: ", ParseURL.Path)
	fmt.Println("Scheme: ", ParseURL.RawQuery)

	//Modifying URL with the help of components
	ParseURL.Host = "learn"
	ParseURL.RawQuery = "Example=1&Ans=2"

	//Constructing a URL string from a URL object
	newURL := ParseURL.String()

	fmt.Println("New URL: ", newURL)
}

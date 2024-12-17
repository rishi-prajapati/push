package main

import (
	"fmt"
	"net/url"
)
const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=dskfjkj343gh"


func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	result,_ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qprams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qprams)

	fmt.Println(qprams["coursename"])

	for _, val := range qprams{
		fmt.Println("Param is : ", val)
	}
	
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "loc.dev",
		Path: "/tutcss",
		RawPath: "user=rishi",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}

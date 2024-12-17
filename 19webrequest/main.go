package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://facebook.com"

func main() {
	fmt.Println("LCO web request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is sof type: %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	fmt.Println(string(databytes))
}

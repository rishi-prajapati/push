package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"
	
	file, err := os.Create("./mylcogofile.txt")

	if err != nil{
		panic(err)
	}

	length, err := io.WriteString(file,content)
	if err != nil{
		panic(err)
	}
	fmt.Println("Lenght is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filname string){
	databyte, err := ioutil.ReadFile(filname)
	if err != nil{
		panic(err)
	}
	fmt.Println("Text data indie the file \n", string(databyte))

}
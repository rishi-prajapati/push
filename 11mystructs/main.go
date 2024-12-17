package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	rishi := User{"Rishi", "prajapatirishi99@gmail.com",true,22}
	fmt.Println(rishi)
	fmt.Printf("rishi details are: %+v\n", rishi)
	fmt.Printf("Name is %v and email is %v", rishi.Name,rishi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

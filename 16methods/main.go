package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	rishi := User{"Rishi", "rishi@go.dev",true,22}
	fmt.Println(rishi)
	fmt.Printf("rishi details are: %+v\n", rishi)
	fmt.Printf("Name is %v and email is %v\n", rishi.Name,rishi.Email)
	rishi.GetStatus()
	rishi.NewMail()
	fmt.Printf("Name is %v and email is %v\n", rishi.Name,rishi.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int

}

func (u User) GetStatus()  {
	fmt.Println("Is user active: ", u.Status)
}

func(u User) NewMail(){
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ",u.Email)
}


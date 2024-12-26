package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct")

	// No Inheritance no super no parent

	user := User{"Rakesh", "Rakesh@go.dev", 30, true}

	fmt.Printf("User is:- %v \n", user)

	fmt.Printf("User is:- %+v \n", user)

	fmt.Printf("User name is:- %v and email is:- %v \n", user.Name, user.Email)
}

type User struct {
	Name   string
	Email  string
	Age    int
	status bool
}

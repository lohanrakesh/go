package main

import "fmt"

func main() {
	fmt.Println("Welcome to go methods")

	user := User{"Rakesh", "rakesh@go.dev", 34, true}

	fmt.Printf("User is:- %+v", user)

	user.GetStatus()
	user.UpdateEmail()

	fmt.Printf("User is:- %+v", user)
}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) GetStatus() {
	fmt.Println("user status is: ", u.Status)
}

func (u User) UpdateEmail() {
	u.Email = "test@go.dev"
	fmt.Println("user updated email is: ", u.Email)
}

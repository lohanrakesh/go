package main

import "fmt"

func main() {
	fmt.Println("Welcome if else")

	loginCount := 24

	message := "Login Message"

	if loginCount > 10 {
		message = "Authenticated user"
	} else if loginCount < 10 {
		message = "Guest user"
	} else {
		message = "10 login count"
	}

	fmt.Println(message)

	if 9%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

}

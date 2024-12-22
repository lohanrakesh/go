package main

import "fmt"

const LoginType = "User" // Public

func main() {
	var username string = "Rakesh"
	fmt.Println(username)
	fmt.Printf("Username type is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("IsLoggedIn type is %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("SmallInt type is %T \n", smallInt)

	var smallFloat float32 = 255.1234567890
	fmt.Println(smallFloat)
	fmt.Printf("SmallFloat type is %T \n", smallFloat)

	//Default value and some alias

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("Default Int type is %T \n", defaultInt)

	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("Default String type is %T \n", defaultString)

	// Implicit type
	var anotherInt = 3567890
	fmt.Println(anotherInt)
	fmt.Printf("Another Int type is %T \n", anotherInt)

	// no var
	noVar := 340000
	fmt.Println(noVar)
	fmt.Printf("No Var type is %T \n", noVar)

	fmt.Println(loginType)
	fmt.Printf("Login Type type is %T \n", loginType)
}

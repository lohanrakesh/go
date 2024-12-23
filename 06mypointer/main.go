package main

import "fmt"

func main() {
	fmt.Println("Wecome to the pointer")

	var ptr *int

	fmt.Println("Value of ptr is:- ", ptr)

	number := 23

	fmt.Println("Value of number is:- ", number)

	newPtr := &number

	fmt.Println("Value of newPtr is:- ", *newPtr)

	*newPtr = *newPtr * 2

	fmt.Println("Update Value of newPtr is:- ", *newPtr)
}

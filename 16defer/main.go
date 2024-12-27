package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer")

	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")

	fmt.Println("Hello")

	ownDefer()
}

func ownDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

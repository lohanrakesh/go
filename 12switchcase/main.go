package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch case")

	// rand.Seed(time.Now().UnixNano())

	number := rand.Intn(6) + 1

	switch number {
	case 1:
		fmt.Println("Dice value is 1")

	case 2:
		fmt.Println("Dice value is 2")

	case 3:
		fmt.Println("Dice value is 3")
		fallthrough

	case 4:
		fmt.Println("Dice value is 4")
		fallthrough

	case 5:
		fmt.Println("Dice value is 5")

	case 6:
		fmt.Println("Dice value is 6")
	}
}

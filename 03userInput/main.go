package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("welcome to user input")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Rating:- ")

	// comma ok || err err

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating:- ", input)
	fmt.Printf("Type of rating is %T ", input)

}

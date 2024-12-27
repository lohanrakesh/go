package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to go files")

	content := "My name is Rakesh Lohan"

	file, err := os.Create("./myname.txt")

	// if err != nil {
	// 	panic(err)
	// }

	checkError(err)

	length, err := io.WriteString(file, content)

	checkError(err)

	fmt.Println("file length is: ", length)

	defer file.Close()

	readFile("./myname.txt")
}

func readFile(file string) {
	content, err := os.ReadFile(file)

	checkError(err)

	fmt.Println("Content of file:- ", content)
	fmt.Println("Content of file:- ", string(content))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

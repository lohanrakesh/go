package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to go web request")

	response, err := http.Get(url)

	fmt.Printf("Type of response:- %T \n", response)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	databyte, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	content := string(databyte)

	fmt.Println(content)
}

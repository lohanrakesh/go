package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["SW"] = "Swift"
	languages["RB"] = "Ruby"

	fmt.Println("Languages are:- ", languages)

	for key, value := range languages {
		fmt.Printf("Languages Key:- %v, Value is:- %v \n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("Languages Value is:- %v \n", value)
	}

	delete(languages, "PY")

	fmt.Println("Languages are:- ", languages)
}

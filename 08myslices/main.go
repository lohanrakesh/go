package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitSlice = []string{"Apple", "Mango", "Peach"}

	fmt.Printf("Type of fruitSlice %T \n", fruitSlice)
	fmt.Println("fruitSlice is:- ", fruitSlice)

	fruitSlice = append(fruitSlice, "Banana", "Grapes")

	fmt.Println("fruitSlice is:- ", fruitSlice)

	fruitSlice = fruitSlice[1:]

	fmt.Println("fruitSlice is:- ", fruitSlice)

	fruitSlice = append(fruitSlice[1:])

	fmt.Println("fruitSlice is:- ", fruitSlice)

	fruitSlice = fruitSlice[1:2]

	fmt.Println("fruitSlice is:- ", fruitSlice)

	scores := make([]int, 4)

	scores[0] = 234
	scores[1] = 678
	scores[2] = 456
	scores[3] = 890

	fmt.Println("scores is:- ", scores)

	// scores[4] = 234

	scores = append(scores, 567, 340, 682)

	fmt.Println("scores is:- ", scores)

	fmt.Println(sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println("scores is:- ", scores)

	cources := []string{"Swift", "Ruby", "Java", "JavaScript", "Python"}

	fmt.Println(cources)

	index := 3
	cources = append(cources[:index], cources[index+1:]...)

	fmt.Println(cources)
}

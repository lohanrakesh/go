package main

import "fmt"

func main() {
	fmt.Println("Welome to array")

	var fruitList [4]string

	fmt.Println("Empty array:- ", fruitList)

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("fruit array:- ", fruitList)
	fmt.Println("Length of fruit array:- ", len(fruitList))

	var vegyList = [10]string{"Potata", "Tomato", "Onion"}

	fmt.Println("Veg array:- ", vegyList)
	fmt.Println("Length of veg array:- ", len(vegyList))

}

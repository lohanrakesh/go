package main

import "fmt"

func main() {
	fmt.Println("Welcome to go functions")

	result := adder(4, 6)

	fmt.Println("Result is:- ", result)

	result2 := multipleAdder(1, 2, 3, 4, 5)

	fmt.Println("Result 2 is:- ", result2)

	result3, message := multipleAdder2(1, 2, 3, 4, 5, 6)

	fmt.Println("Result 3 is:- ", result3)

	fmt.Println("Result 3 is:- ", message)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func multipleAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func multipleAdder2(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "good to find result"
}

package main

import "fmt"

func main() {
	fmt.Println("Loops")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursady"}

	for i := 0; i < len(days); i++ {
		fmt.Println("Day is:-", days[i])
	}

	for i := range days {
		fmt.Println("Day is:-", days[i])
	}

	for index, value := range days {
		fmt.Printf("Index is %v value is %v \n", index, value)
	}

	iterateValue := 0

	// for iterateValue < 10 {
	// 	fmt.Println(iterateValue)
	// 	iterateValue++
	// }

	for iterateValue < 10 {

		// if iterateValue == 4 {
		// 	break
		// }

		// if iterateValue == 4 {
		// 	iterateValue++
		// 	continue
		// }

		if iterateValue == 4 {
			goto Rl
		}

		fmt.Println(iterateValue)
		iterateValue++
	}

Rl:
	fmt.Println("You came Here")
}

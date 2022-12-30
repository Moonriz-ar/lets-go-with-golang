package main

import "fmt"

func main() {
	fmt.Println("loops")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for index, value := range days {
		fmt.Printf("index is %v and value is %v\n", index, value)
	}

	rougueValue := 1

	for rougueValue < 10 {
		// if rougueValue == 5 {
		// 	break // breaks from the loop
		// }

		if rougueValue == 6 {
			rougueValue++
			continue
		}

		fmt.Println("Value is:", rougueValue)
		rougueValue++
	}
}

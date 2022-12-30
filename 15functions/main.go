package main

import "fmt"

func main() {
	fmt.Println("functions")
	greet()

	result := add(3, 5)
	fmt.Println(result)

	resultAllNumbers := addAllNumbers(2, 2, 5, 5)
	fmt.Println(resultAllNumbers)
}

func greet() {
	fmt.Println("Hi from golang")
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func addAllNumbers(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

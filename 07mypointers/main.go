package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// & gets the pointer address
	// * retries value according to pointer address

	var number *int
	fmt.Println("Value of number is ", number) // points to nothing, pointer == nil

	myNumber := 23
	var pointerTwo = &myNumber // & reference, gets the pointer address of the variable after &

	fmt.Println("Value of pointer is ", pointerTwo)
	fmt.Println("Value is ", *pointerTwo)

	*pointerTwo = *pointerTwo + 2
	fmt.Println("New value is: ", myNumber)

}
package main

import (
	"fmt"
)

// public variable
// to make a variable or function public, it needs to start with capital letter
// to make a variable or function private, it needs to start with lower letter
const LoginToken string = "fdsjafjdsadfj"

func main() {
	var username string = "mario"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.85438954983954
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default value
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	// implicit type
	var website = "www.facebook.com"
	println(website)

	// no var style
	numberOfUser := 300000
	println(numberOfUser)
}
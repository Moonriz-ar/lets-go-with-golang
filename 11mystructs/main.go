package main

import "fmt"

func main() {
	fmt.Println("structs")

	andrea := User{"Andrea", "andre93_lin@hotmail.com", true, 29}
	fmt.Println(andrea)
	fmt.Printf("Andrea's details are: %+v\n", andrea)
	fmt.Println(andrea.Name)
}

// User of an application
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

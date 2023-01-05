package main

import (
	"fmt"
)

// a struct can nest another struct (or pointer to a struct)

type Address struct {
	Province string
	City     string
}

type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	user1 := User{
		Name:   "Mario",
		Gender: "Male",
		Address: Address{
			Province: "Mountain",
			City:     "Buenos Aires",
		},
	}
	fmt.Printf("user1=%#v\n", user1)
}

package main

import "fmt"

func main() {
	fmt.Println("structs")

	andrea := User{"Andrea", "andre93_lin@hotmail.com", true, 29}
	fmt.Println(andrea)
	fmt.Printf("Andrea's details are: %+v\n", andrea)
	fmt.Println(andrea.Name)
	andrea.GetStatus()
	andrea.NewEmail()
	fmt.Printf("Andrea's details are: %+v\n", andrea)
}

// User of an application
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// GetStatus prints the current user status
func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewEmail() {
	// a copy of the object is passed along, assigning a value to a property does not change original object
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}

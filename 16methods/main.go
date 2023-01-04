package main

import "fmt"

// a method is a function affiliated with a type
// a method is a function with an implicit first argument, called a receiver
// syntax of method: func (r ReceiverType) funcName(parameters) (results)
func main() {
	fmt.Println("methods")

	andrea := User{"Andrea", "a@hotmail.com", true, 29}
	fmt.Println(andrea)
	fmt.Printf("Andrea's details are: %+v\n", andrea)
	fmt.Println(andrea.Name)
	andrea.GetStatus()
	andrea.SetEmail("b@gmail.com")
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

// 指针类型的接收者

// SetEmail sets the user's email
func (u *User) SetEmail(newEmail string) {
	u.Email = newEmail
	fmt.Println("New email of this user is:", u.Email)
}

// a copy of the object is passed along, assigning a value to a property does not change original object
// func (u *User) SetEmail(newEmail string) {
// 	u.Email = newEmail
// 	fmt.Println("New email of this user is:", u.Email)
// }

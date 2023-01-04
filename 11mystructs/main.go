package main

import "fmt"

func main() {
	fmt.Println("structs")

	andrea := User{"Andrea", "andre93_lin@hotmail.com", true, 29}
	fmt.Println(andrea)
	fmt.Printf("Andrea's details are: %+v\n", andrea)
	fmt.Println(andrea.Name)

	// instantiate and initialize a student
	mark := Student{Human{"Mark", 14, 70}, "Computer Science"}

	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His specialty is ", mark.specialty)

	// modify mark's specialty
	mark.specialty = "Business Administration"
	fmt.Println("Mark changed his specialty")
	fmt.Println("His specialty is ", mark.specialty)
}

// User of an application
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// embedded fields in struct

// Human represents a person with name, age and weight
type Human struct {
	name   string
	age    int
	weight int
}

// Student represents a Human with specialty
type Student struct {
	Human     // embedded field, it means Student struct includes all fields that Human has
	specialty string
}

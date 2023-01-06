package main

import "fmt"

// difference between interface with variable and pointer

type mover interface {
	move()
}

type sayer interface {
	say()
}

// nest interface
type animal interface {
	mover
	sayer
}

type person struct {
	name string
	age  int8
}

// using variable in receiver: variable and pointer can both be saved a copy in the argument
// func (p person) move() {
// 	fmt.Printf("%s is running...\n", p.name)
// }

// using pointer in receiver:
func (p *person) move() {
	fmt.Printf("%s is running...\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s is speaking...\n", p.name)
}

func main() {
	var m mover // declare a variable with type mover
	// p1 := person{ // a variable type person
	// 	name: "mario",
	// 	age:  18,
	// }
	p2 := &person{ // a pointer to a variable type person
		name: "luigi",
		age:  17,
	}
	// m = p1 // cannot use p1 (variable of type person) as mover value in assignment: person does not implement mover (method move has pointer receiver)
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer // declare a variable with type sayer
	s = p2
	s.say()
	fmt.Println(s)

	var a animal // declare a variable with animal type
	a = p2
	a.move()
	a.say()
	fmt.Println(a)
}

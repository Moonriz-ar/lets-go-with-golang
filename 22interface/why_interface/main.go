package main

import "fmt"

// why is interface needed
// 在Go语言中接口（interface）是一种类型，一种抽象的类型。
// interface is a type, an abstract type

type dog struct{}

func (d dog) say() {
	fmt.Println("Woof woof")
}

type cat struct{}

func (c cat) say() {
	fmt.Println("Miau miau")
}

type crocodile struct{}

// interface does not care which type you are, only care about which method you execute
// define an abstract method, any type that implements say() method can be considered a sayer() type
type sayer interface {
	say()
}

func pat(arg sayer) {
	arg.say()
}

func main() {
	// garfield := cat{}
	// pat(garfield)
	// kuma := dog{}
	// pat(kuma)
	// croco := crocodile{}
	// pat(croco) // cannot use croco (variable of type crocodile) as sayer value in argument to pat: crocodile does not implement sayer (missing method say)

	var s sayer
	c2 := cat{}
	s = c2 // c2 (cat type) can be assigned to s (sayer type) because it implements say()
	fmt.Println(s)
}

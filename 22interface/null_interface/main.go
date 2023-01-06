package main

import "fmt"

// empty interface
// when inside an interface there isnt any defined method
// any type implements empty interface => it means that inside a variable that is an empty interface, you can put anything inside it!

// empty interface usually does not need to be defined beforehand
type xxx interface {
}

// some uses for empty interface
// 1. in functions, empty interface can defined as parameter to accept any type as argument
// for example, fmt.Println(x) accepts any type.
// in the go source code, it is func Println(a ...interface{}) (n int, err error) {...}
// 2. in map, empty interface can be the value, to accep any type as value

func main() {
	// just create variable and type it as <interface{}>
	var x interface{}
	x = "hello"
	x = 100
	x = false
	x = 0.2
	fmt.Println(x)

	// example of 2 use case
	// var m = make(map[string]interface{}, 16)
	// m["name"] = "mario"
	// m["age"] = 18
	// m["hobby"] = "rescue peach"

	// 類型斷言
	result, ok := x.(float64)
	if !ok {
		fmt.Println("it is not float64")
	} else {
		fmt.Println("it is float64", result)
	}

	// also can use switch to determine the type from any
	switch v := x.(type) {
	case string:
		fmt.Println("it is type string, value:", v)
	case bool:
		fmt.Println("it is type boolean, value:", v)
	case int:
		fmt.Println("it is type int, value:", v)
	case float64:
		fmt.Println("it is type float64, value:", v)
	default:
		fmt.Println("I dont know what type it is:", v)
	}
}

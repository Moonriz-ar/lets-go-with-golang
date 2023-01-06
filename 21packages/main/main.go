package main

// go does not allow import a package and not using it
// go does not allow 循環引用包, because it would cause infinite import cycle

import (
	// this way of importing is called anonymous import, and it is to run the init function of imported package
	// _ "example/21packages/calc"
	"example/21packages/calc"
	"fmt"
)

func main() {
	fmt.Println("hi")
	fmt.Println(calc.Add(1, 2))
}

package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// & gets the pointer address
	// * retries value according to pointer address

	var number *int
	fmt.Println("Value of number is ", number) // points to nothing, pointer == nil

	myNumber := 23
	var pointerTwo = &myNumber // & reference, gets the pointer address of the variable after &

	fmt.Println("Value of pointer is ", pointerTwo)
	fmt.Println("Value is ", *pointerTwo)

	*pointerTwo = *pointerTwo + 2
	fmt.Println("New value is: ", myNumber)

	//LIWENZHOU///////////////////////////////////////////
	a := 10                      // int
	b := &a                      // *int
	fmt.Printf("%v %p\n", a, &a) // 10 0xc00001c0c0
	fmt.Printf("%v %p\n", b, b)  // 0xc00001c0c0 0xc00001c0c0
	// the b variable is a pointer with int type (*int), saves the address of a
	c := *b        // 根據内存地址去取值
	fmt.Println(c) // 10

	d := 1
	modify1(d)
	fmt.Println("d value after calling modify1 func: ", d)
	modify2(&d)
	fmt.Println("d value after calling modify2 func:", d)

	// new and make are for internal memory assignments

	// panics because of invalid memory address
	/*
		var e *int
		*e = 100
		fmt.Println(*e)
	*/

	var e *int
	e = new(int)
	fmt.Println("e value", *e) // 0
	*e = 100
	fmt.Println("e value", *e) // 100

	// panic: assignment to entry in nil map
	/*
		var f map[string]int
		f["mario"] = 100
		fmt.Println(f)
	*/

	var f map[string]int
	f = make(map[string]int, 10)
	f["mario"] = 100
	fmt.Println(f)
}

func modify1(x int) {
	x = 100
}

func modify2(y *int) {
	*y = 100
}

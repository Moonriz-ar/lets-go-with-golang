package calc

import "fmt"

// Add requires two int input, and returns the sum as int
func Add(x, y int) int {
	return x + y
}

// init function is the first function to execute when a package is imported
// init function does not have parameter and return value
func init() {
	fmt.Println("execute calc.init()")
}

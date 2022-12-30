package main

import "fmt"

func main() {
	fmt.Println("array")

	var fruitList [4]string // array length is fixed and defined at initialization

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Kiwi"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("The length of fruit list is: ", len(fruitList)) // returns the fixed length of the array

	var vegList = [3]string{"potato", "beans", "brocoli"}
	fmt.Println("Veggie list is: ", vegList)
}
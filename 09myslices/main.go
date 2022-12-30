package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices")

	var fruitList = []string{"Apple", "Peach", "Orange"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Kiwi")
	fmt.Println(fruitList)
	fruitList = append(fruitList, "Grape", "Tomato") // can add more than one item with append
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3]) // assign the sub slice [1:3] (fron index zero to index two) to fruitList 
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465	
	highScores[3] = 2865
	// causes runtime error index out of range
	// highScores[4] = 4382843 

	// reallocates memory and allows all values to be accomodated
	highScores = append(highScores, 555, 356, 432)
	fmt.Println(highScores)
	
	fmt.Println("Is the slice sorted? ", sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println("Is the slice sorted? ", sort.IntsAreSorted(highScores))
}
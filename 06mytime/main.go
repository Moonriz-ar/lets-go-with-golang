package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// format takes in string and specific numbers/string should be input for the format
	// checkout doc https://pkg.go.dev/time#Time.Format
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 19, 23, 15, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 Monday"))
}
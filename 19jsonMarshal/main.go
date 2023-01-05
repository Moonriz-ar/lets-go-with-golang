package main

import (
	"encoding/json"
	"fmt"
)

// Student has id, gender and name
type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	// Age int // if no tag, json marshal would default use the field name, as is
	// documentNumber int // if there is private field, json marshal can't access it
}

// Class has title and a slice of students
type Class struct {
	Title    string    `json:"title"`
	Students []Student `json:"students"`
}

func main() {
	// create a class
	c := &Class{
		Title:    "golang-101",
		Students: make([]Student, 0, 200),
	}

	// add ten students
	for i := 0; i < 10; i++ {
		student := Student{
			Name:   fmt.Sprintf("student%02d", i),
			Gender: "female",
			ID:     i,
		}
		c.Students = append(c.Students, student)
	}

	// encoding json: struct => string with json format
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("encoded json type: %T\n", data) // []uint8
	fmt.Printf("encoded json:%s\n", data)

	// decoding json: string with json format => struct
	str := `{"title":"101","students":[{"id":0,"gender":"男","name":"stu00"},{"id":1,"Gender":"男","name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{} // instanciate a Class struct
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed")
		return
	}
	fmt.Printf("decoded json: %#v\n", c1)
}

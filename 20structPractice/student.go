package main

import "fmt"

type student struct {
	id    int // 學號是唯一的
	name  string
	class string
}

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, 0, 100),
	}
}

// 1- add student
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)
}

// 2- edit student
func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudents {
		if newStu.id == v.id { // when student id matches, means that found the student that needs to be modified
			s.allStudents[i] = newStu // copy the new student value
			return
		}
	}

	// if reaches here, means that could not find the student
	fmt.Printf("The is an error with the data input, there isn't any student with student id #%d", newStu.id)
}

// 3- show all students
func (s *studentMgr) getAllStudents() {
	for _, v := range s.allStudents {
		fmt.Printf("Student ID: %v | Name: %v | Class: %v\n", v.id, v.name, v.class)
	}
}

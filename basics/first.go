package main

import (
	"fmt"
)

type Person interface {
	GetName() string
}
type Student struct {
	Rollno int64
	Name   string
	Marks  int64
}

func (s *Student) GetName() string {
	return s.Name
}
func main() {
	s1 := Student{
		Rollno: 161,
		Name:   "Syed Muhammad Rizwan",
		Marks:  100,
	}
	s2 := Student{
		Rollno: 19,
		Name:   "Syed Abdullah Rizwan",
		Marks:  89,
	}
	xp := []Student{s1, s2}
	fmt.Printf("Students: %+v \n", xp)

	//Playing with pointer
	var s *Student = new(Student)
	s.Rollno = 1
	s.Name = "Jack"
	fmt.Println("Student: ", *s)

	xp = append(xp, *s)

	fmt.Println("Students: ", xp)

	//interface testing
	var p Person = &s2
	fmt.Println("Person: ", p.GetName())
}

package main

import (
	"fmt")

type Student struct {
	Rollno int64
	Name   string
	Marks  int64
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
}

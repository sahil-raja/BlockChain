package main

import (
	"fmt"
	"strings"
)


type Student struct {
	rollnumber 	int
	name 		string
	address 	string
}

func NewStudent(rollno int, name string, address string)*Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	return s
}

type StudentList struct {
	list []*Student
}

func(ls *StudentList) CreateStudent(rollno int, name string, address string)*Student {
	st := NewStudent(rollno, name, address)
	ls.list = append(ls.list, st)
	return st
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAA")
	student.CreateStudent(25, "Naveed", "BBBBBB")

	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25),0, strings.Repeat("=", 25))
	fmt.Printf("Student roll no \t %v \n", student.list[0].rollnumber)
	fmt.Printf("Student name \t\t %v \n", student.list[0].name)
	fmt.Printf("Student address \t %v \n", student.list[0].address)

	
	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25),1, strings.Repeat("=", 25))
	fmt.Printf("Student roll no \t %v \n", student.list[1].rollnumber)
	fmt.Printf("Student name \t\t %v \n", student.list[1].name)
	fmt.Printf("Student address \t %v \n", student.list[1].address)
	
}
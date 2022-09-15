package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)


type Student struct {
	rollnumber 	int
	name 		string
	address 	string
	subjects[]	string
}

func NewStudent(rollno int, name string, address string, sub[] string)*Student {
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.subjects = sub	
	return s
}

type StudentList struct {
	list []*Student
}

func(ls *StudentList) CreateStudent(rollno int, name string, address string, sub[] string)*Student {
	st := NewStudent(rollno, name, address, sub)
	ls.list = append(ls.list, st)
	return st
}

func main() {
	student := new(StudentList)
	student.CreateStudent(24, "Asim", "AAAAAA",  []string {"PF", "AP", "COAL"})
	student.CreateStudent(25, "Naveed", "BBBBBB",  []string {"Eng", "Cal", "IS"})

	for i:=0; i<2; i++{
	var packed_str string
	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25),i, strings.Repeat("=", 25))
	fmt.Printf("Student roll no \t %v \n", student.list[i].rollnumber)
	fmt.Printf("Student name \t\t %v \n", student.list[i].name)
	fmt.Printf("Student address \t %v \n", student.list[i].address) 
	packed_str = strconv.Itoa(student.list[i].rollnumber) + student.list[i].name + student.list[i].address

		for j:=0; j<len(student.list[i].subjects); j++{
			fmt.Printf("Subject  %v \t %v \n",j , student.list[i].subjects[j])
			packed_str += student.list[i].subjects[j]
		}
	fmt.Printf("Student Hash \t %x \n", sha256.Sum256([]byte(packed_str)))
	}

	
}
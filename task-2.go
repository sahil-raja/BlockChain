package main

import (
	"fmt"
)


type employee struct {
	name string
	salary int
	position string
}


type company struct {
	companyName string
	employees []employee
}

func main(){
	emp1 := employee{"Amir",80000, "Full-Stack Developer"}
	emp2 := employee{"Malik",50000, "Front-End Developer"}
	emp3 := employee{"Sahab",30000, "Back-End Developer"}

	emplys := [] employee{emp1, emp2, emp3}
	
	cmp1 := company{"Tetra",emplys}

	fmt.Printf("Company Name=%v, Employees=%v", cmp1.companyName, cmp1.employees)

}

package main

import (
	"fmt"
)

type student struct {
	name         string
	collage_name string
	course       string
}

/*
 displayname() method has student as the receiver type
*/
func (s student) displayname() {
	fmt.Printf(" student details of %s \n is %s \n %s", s.name, s.collage_name, s.course)
}

func main() {
	stud1 := student{
		name:         "Sangam biradar",
		collage_name: "alliance University Bengalore",
		course:       "btech",
	}
	stud1.displayname() //Calling displayname() method of student type
}

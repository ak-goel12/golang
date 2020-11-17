package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

func (p person) fullname() string {
	fmt.Printf("inside method: %p\n", &p)
	return p.fname + p.lname
}
func main() {
	p1 := person{"sangam", "biradar", 24}
	fmt.Printf(p1.fullname())
	fmt.Printf("\n inside method: %p \n", &p1)
	fmt.Printf(p1.fullname())

}

// p1 is reciever value for the call to fullname
// fullname is operating on value of p1

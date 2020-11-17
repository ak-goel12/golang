//using pointer
package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

func (p *person) changeage(newage int) {
	p.age = newage
}

func main() {
	p1 := person{"sangam", "biradar", 23}
	fmt.Println(p1.age)
	p1.changeage(24)
	fmt.Println(p1.age)

}

package main

import "fmt"

type xy struct {
	X, Y int
}

var (
	z1 = xy{1, 2}  // has type xy
	z2 = xy{X: 1}  // Y:0 is implicit
	z3 = xy{}      // X:0 and Y:0
	p  = &xy{1, 2} // has type *xy
)

func main() {
	fmt.Println(z1, p, z2, z3)
}

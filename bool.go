package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 7 // if we change this to 42, a == b will evaluate to true
	b := 42
	fmt.Println(a == b) // experiment with different operators: <=, >=, !=, >, <

}

package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = hotdog(a) // we can convert the value of a to a value of type hotdog
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

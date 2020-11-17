package main

import "fmt"

type xy struct {
	X int
	Y int
}

func main() {
	z := xy{1, 2}
	p := &z
	p.X = 20
	fmt.Println(z)
}

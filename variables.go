package main

import (
	"fmt"
	"strconv"
)

var i int = 43 // variable declaration
// if u declare a variable u must have to use it otherwise it will throw an error
func main() {
	fmt.Println(i)
	var i int = 27 //shadowing of a variable
	fmt.Println(i)
	j := 12 //another way to declare variable but the shortcoming is u don't about the datatype
	fmt.Printf("%v, %T\n", j, j)
	//group variable
	var (
		a int    = 5
		b string = "hi"
	)
	fmt.Println(a, b)
	var m float32 = 42.5
	var n int
	n = int(m) //type conversion
	fmt.Printf("%v, %T\n", n, n)
	var q int = 43
	var s string
	s = strconv.Itoa(q) // to convert int into string we have to use a strconv package
	fmt.Println(s)
	var HTTP string = "http://google.com" //we have to capatalize all acronyms
	fmt.Println(HTTP)
}

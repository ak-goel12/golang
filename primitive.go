package main

import (
	"fmt"
)

func main() {
	// boolean types
	var a bool = true
	fmt.Printf("%v, %T\n", a, a)
	m := 1 == 1
	n := 2 == 1
	fmt.Printf("%v, %T\n", m, m)
	fmt.Printf("%v, %T\n", n, n)
	//complex type
	var i complex64 = 1 + 2i
	var j complex64 = 2 + 5i
	fmt.Println(i + j)
	fmt.Println(i - j)
	fmt.Println(i * j)
	fmt.Println(i / j)
	var l complex64 = complex(1, 3)
	fmt.Printf("%v, %T", l, l)
	// string
	var str string = "this is string" // string is immutable //utf-8
	b := []byte(str)                  //can be converted to byte
	fmt.Printf("%v, %T\n", b, b)
	fmt.Println(string(str[2]))
	//rune
	r := 'a'                   // it uses single inverted comma
	fmt.Printf("%v, %T", r, r) //alias for int 32 //specials method normally required to process string.Reader#ReadRune
}

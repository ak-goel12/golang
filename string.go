package main

import (
	"fmt"
)

var string_literal string

func main() {
	s := "Hello, playground"
	string_literal = `"Hello,
	
	playground"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println(string_literal)
	fmt.Printf("%T\n", string_literal)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	fmt.Println("")

	/* Show each character from string s in hexadecimal */
	for i, v := range s {
		fmt.Printf("At index position %d we have hex %#X\n", i, v)
	}
}

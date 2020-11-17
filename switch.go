package main

import (
	"fmt"
)

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true 1")
		fallthrough //if fallthrough is not present then the switch will exit
		//after the condition evaluates to true
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
	}
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Dr No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("This is q")
	default:
		fmt.Println("this is default")
	}

	fmt.Println(unhex('E'))
	fmt.Println(shouldEscape('+'))
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	var pass string
	regex := "^[a-zA-Z0-9]{8}$" //regex for validating password

	//input password from user
	fmt.Println("Enter your password(must have atleast 8 character(only alphabets)): ")
	fmt.Scanln(&pass)
	//match password with regex
	match1, _ := regexp.MatchString(regex, pass)
	// Check whether the string match with regex or not
	if match1 {
		fmt.Println("Valid Password")
	} else {
		fmt.Println("Please enter password according to requirement")
	}
}

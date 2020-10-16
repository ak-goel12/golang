package main

import "fmt"

//unenumerated constant
const (
	isAdmin = 1 << iota //1 shifted right by 0 time
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeEurope
	canSeeAsia
)

func main() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles) // this way is used in different applications
	fmt.Printf("Is Admin? %v", isAdmin&roles == isAdmin)
}

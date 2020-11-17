package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	var a [3]int

	fmt.Println(copy(a[:2], s))
	fmt.Println(a)
	m := [...]int{1, 2, 3}
	n := make([]int, 3)

	fmt.Println(copy(n, m[:2]))
	fmt.Println(n)
}

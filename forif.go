package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 { // try changing the number to 3, 4, etc.
			fmt.Println(i)
		}
	}
}

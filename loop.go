// no while in go
package main

import "fmt"

func main() {

	i := 1       //init variable
	for i <= 3 { //condition
		fmt.Println(i)
		i = i + 1 //update
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	// nested loops
	for i := 0; i <= 10; i++ { //outer loop
		fmt.Printf("Outer loop: %d\n", i)
		for j := 0; j < 3; j++ { //inner loop
			fmt.Printf("\tInner loop: %d\n", j)
		}
	}
	// another way for loop
	x := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")
}

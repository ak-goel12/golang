package main

import (
	"fmt"
)

func main() {
	// two ways to declare 1-D array
	//1st way
	/*[...] basically it means that the size of array will
	be equal to the paramater passed to it*/
	grades := [...]int{97, 57, 64}
	fmt.Printf("Grades : %v,%T", grades, grades)
	b := grades
	b[1] = 25
	/* when u copy an array to other then it's does
	not point to the same array instead it makes the copy of an array and passed to */
	fmt.Printf("\nGrades : %v,%T", grades, grades)
	fmt.Printf("\nB : %v,%T", b, b)
	//2nd way
	var student [3]string //u can pass parameter too
	fmt.Printf("\nStudents : %v %T", student, student)
	student[0] = "akshit" // another way to assign value by index
	student[2] = "aman"
	student[1] = "rohit"
	fmt.Printf("\nStudents : %v %T", student, student)
	//len method give the length of array
	fmt.Printf("\nNo of Students : %v \n", len(student))
	//2-D array
	//1st way
	var matrix [3][3]int = [3][3]int{[3]int{1, 4, 3}, [3]int{33, 2, 6}, [3]int{2, 6, 3}}
	fmt.Println(matrix)
	//2nd way
	var mat [3][3]int
	mat[0] = [3]int{1, 4, 5}
	mat[1] = [3]int{6, 4, 3}
	mat[2] = [3]int{5, 3, 5}
	fmt.Println(mat)
}

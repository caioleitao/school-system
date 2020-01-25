package main

import "fmt"

func main() {
	var student string
	var n1 float64
	var n2 float64

	fmt.Println("Enter student name: ")
	fmt.Scan(&student)
	fmt.Println("Enter the first semester grade: ")
	fmt.Scanf("%f", &n1)
	fmt.Println("Enter the second semester grade: ")
	fmt.Scanf("%f", &n2)
	note := (n1 + n2) / 2

	if note < 7 {
		fmt.Println("The final grade of ", student, "was:", note, "(DISAPPROVED)")
	} else {
		fmt.Println("The final grade of ", student, "was:", note, "(APPROVED)")
	}


}

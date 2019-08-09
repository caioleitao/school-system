package main

import "fmt"

func main() {
	var student string
	var n1 float64
	var n2 float64

	fmt.Println("Digite o nome do aluno: ")
	fmt.Scan(&student)
	fmt.Println("Digite a nota do primeiro semestre: ")
	fmt.Scanf("%f", &n1)
	fmt.Println("Digite a nota do segundo semestre: ")
	fmt.Scanf("%f", &n2)
	note := (n1 + n2) / 2

	if note < 7 {
		fmt.Println("A nota final de", student, "foi:", note, "[REPROVADO]")
	} else {
		fmt.Println("A nota final de", student, "foi:", note, "[APROVADO]")
	}


}
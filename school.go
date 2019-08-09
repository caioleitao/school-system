package main

import "fmt"

func main() {
	var n1 float64
	var n2 float64


	fmt.Println("Digite a nota do primeiro semestre: ")
	fmt.Scanf("%f", &n1)
	fmt.Println("Digite a nota do segundo semestre: ")
	fmt.Scanf("%f", &n2)
	note := (n1 + n2) / 2
	fmt.Println("A nota final do aluno foi:", note)


}
package main

import "fmt"

func main() {
	//Declaração de variaveis

	A, B, C := 0.0, 0.0, 0.0

	//Leitura de dados
	fmt.Printf("Informe o lado A:")
	fmt.Scan(&A)
	fmt.Printf("Informe o lado B:")
	fmt.Scan(&B)
	fmt.Printf("Informe o lado C:")
	fmt.Scan(&C)

	//Teste e exibição dos resultados
	if !(A < (B+C) && B < (A+C) && C < (A+B)) {
		fmt.Printf("Lados Invalidos")
	} else {
		switch {
		case A == B && B == C:
			fmt.Printf("Triangulo equilatero")
		case A == B || A == C || B == C:
			fmt.Printf("Triangulo isosceles")
		default:
			fmt.Printf("Triangulo escaleno")
		}
	}
}

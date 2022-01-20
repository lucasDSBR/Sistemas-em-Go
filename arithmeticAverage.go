package main

import (
	"fmt"
)

func main() {
	//Declaração de variaveis
	var note1 = 0.0
	var note2 = 0.0
	var note3 = 0.0

	//Requisitando valores
	fmt.Println("Informe o valor da nota 1")
	fmt.Scanln(&note1)
	fmt.Println("Informe o valor da nota 2")
	fmt.Scanln(&note2)
	fmt.Println("Informe o valor da nota 3")
	fmt.Scanln(&note3)

	var media = (note1 + note2 + note3) / 3

	fmt.Println("A média geral das notas é: ", media)

}

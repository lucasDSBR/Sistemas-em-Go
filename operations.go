package main

import "fmt"

func main() {
	var valorUm float64
	var valorDois int
	fmt.Println("Entre com um valor: ")
	fmt.Scanln(&valorUm)
	fmt.Println("Entre com outro valor: ")
	fmt.Scanln(&valorDois)

	var soma float64 = valorUm + float64(valorDois)

	fmt.Println("O resultado da Soma é: ", soma)

}

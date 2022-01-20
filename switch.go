package main

import "fmt"

func main() {
	numero := -1
	switch {
	case numero < 0:
		fmt.Printf("O número %d é negativo", numero)
	case numero > 0:
		fmt.Printf("O número %d é positivo", numero)
	default:
		fmt.Printf("Erro")
	}
}

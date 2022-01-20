package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Por favor, digite seu nome:")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Olá, " + name + "! Você escreveu seu codigo em Go")
	name = strings.TrimSpace(name)
	fmt.Println(name)
}

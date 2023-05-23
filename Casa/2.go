package main

import (
	"fmt"
)

func verificarParImpar(num *int) {
	if *num%2 == 0 {
		*num = 0 // atualiza para 0 se for par
	} else {
		*num = 1 // atualiza para 1 se for ímpar
	}
}

func main() {
	numero := 7

	fmt.Println("Número original:", numero)

	verificarParImpar(&numero)

	fmt.Println("Número atualizado:", numero)
}

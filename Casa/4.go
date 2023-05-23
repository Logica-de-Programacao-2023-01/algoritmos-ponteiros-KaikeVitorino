package main

import (
	"fmt"
)

func atualizarUltimosDigitos(num *int) {
	// Obter os dois últimos dígitos do número
	digito1 := *num % 10
	digito2 := (*num / 10) % 10

	// Calcular a soma dos dois últimos dígitos
	soma := digito1 + digito2

	// Atualizar o valor da variável para a soma dos dois últimos dígitos
	*num = soma
}

func main() {
	numero := 1234

	fmt.Println("Número original:", numero)

	atualizarUltimosDigitos(&numero)

	fmt.Println("Número atualizado:", numero)
}

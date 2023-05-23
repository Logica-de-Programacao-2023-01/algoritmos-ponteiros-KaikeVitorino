package main

import (
	"fmt"
)

func modificarValor(ponteiro *int) {
	// Modifica o valor apontado pelo ponteiro
	*ponteiro = 42
}

func main() {
	// Cria uma variável inteira e atribui um valor inicial
	var numero int = 10

	// Cria um ponteiro para a variável
	ponteiro := &numero

	fmt.Println("Valor antes da modificação:", numero)

	// Chama a função para modificar o valor
	modificarValor(ponteiro)

	fmt.Println("Valor depois da modificação:", numero)
}

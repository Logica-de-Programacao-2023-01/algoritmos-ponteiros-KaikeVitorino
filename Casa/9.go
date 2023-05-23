package main

import (
	"fmt"
)

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(l *Livro, desconto float64) {
	l.Preco *= 1 - desconto/100
}

func main() {
	livro := Livro{Titulo: "O Senhor dos Anéis", Autor: "J. R. R. Tolkien", Preco: 50.0}
	fmt.Printf("Antes: Preço = %.2f\n", livro.Preco)
	aplicarDesconto(&livro, 10.0)
	fmt.Printf("Depois: Preço = %.2f\n", livro.Preco)
}

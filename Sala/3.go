package main

import (
	"fmt"
)

type Produto struct {
	nome    string
	preco   float64
	estoque int
}

func atualizarPreco(p *Produto, novoPreco float64) {
	p.preco = novoPreco
}

func main() {
	produto := Produto{
		nome:    "Camiseta",
		preco:   29.99,
		estoque: 10,
	}

	fmt.Println("Preço atual:", produto.preco)

	novoPreco := 39.99
	atualizarPreco(&produto, novoPreco)

	fmt.Println("Preço atualizado:", produto.preco)
}

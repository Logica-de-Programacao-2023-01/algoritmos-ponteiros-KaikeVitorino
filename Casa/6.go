package main

import (
	"fmt"
)

type Livro struct {
	titulo string
	autor  string
}

func alterarTituloSeAutorAnonimo(livro *Livro) {
	if livro.autor == "Anônimo" {
		livro.titulo = "Desconhecido"
	}
}

func main() {
	livro1 := Livro{
		titulo: "Livro 1",
		autor:  "Autor 1",
	}

	livro2 := Livro{
		titulo: "Livro 2",
		autor:  "Anônimo",
	}

	fmt.Println("Livro 1 antes:", livro1)
	fmt.Println("Livro 2 antes:", livro2)

	alterarTituloSeAutorAnonimo(&livro1)
	alterarTituloSeAutorAnonimo(&livro2)

	fmt.Println("Livro 1 depois:", livro1)
	fmt.Println("Livro 2 depois:", livro2)
}

package main

import (
	"fmt"
)

func inverterBooleano(b *bool) {
	*b = !*b
}

func main() {
	valor := true
	fmt.Println("Valor inicial:", valor)

	inverterBooleano(&valor)
	fmt.Println("Valor invertido:", valor)
}

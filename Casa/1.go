package main

import (
	"fmt"
)

func somaPrimeirosNaturais(n *int, valor int) {
	soma := 0
	for i := 1; i <= valor; i++ {
		soma += i
	}
	*n = soma
}

func main() {
	x := 0
	fmt.Printf("Antes: x = %d\n", x)
	somaPrimeirosNaturais(&x, 5)
	fmt.Printf("Depois: x = %d\n", x)
}

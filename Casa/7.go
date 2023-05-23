package main

import (
	"fmt"
)

type Conta struct {
	Saldo   float64
	Titular string
}

func depositar(c *Conta, valor float64) {
	c.Saldo += valor
}

func main() {
	conta := Conta{Saldo: 100.0, Titular: "João"}
	fmt.Printf("Antes: Saldo = %.2f\n", conta.Saldo)
	depositar(&conta, 50.0)
	fmt.Printf("Depois: Saldo = %.2f\n", conta.Saldo)
}

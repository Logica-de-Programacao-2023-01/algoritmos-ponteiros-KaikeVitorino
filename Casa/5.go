package main

import (
	"fmt"
	"math"
)

func mediaComPi(x *float64) {
	*x = (*x + math.Pi) / 2
}

func main() {
	x := 3.0
	fmt.Printf("Antes: x = %.2f\n", x)
	mediaComPi(&x)
	fmt.Printf("Depois: x = %.2f\n", x)
}

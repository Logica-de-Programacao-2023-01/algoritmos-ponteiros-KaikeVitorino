package main

import (
	"fmt"
)

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 1
	y := 2
	fmt.Printf("Antes: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Depois: x = %d, y = %d\n", x, y)
}

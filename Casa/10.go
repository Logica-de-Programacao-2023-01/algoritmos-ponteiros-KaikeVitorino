package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func fillPrimes(slice *[]int, size int) {
	count := 0
	num := 2

	for count < size {
		if isPrime(num) {
			*slice = append(*slice, num)
			count++
		}
		num++
	}
}

func main() {
	var primes []int
	size := 10

	fillPrimes(&primes, size)

	fmt.Println("Primeiros", size, "números primos:", primes)
}

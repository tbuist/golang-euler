package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	max := int(math.Sqrt(float64(x)))
	for i := 2; i <= max; i++ {
		if x % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0

	for i := 2; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	fmt.Println("Sum of primes below 2 million =", sum)
}
package main

import (
	"fmt"
)

// version 1
/*
func isPrime(x int) bool {
	for i := int(math.Sqrt(float64(x))); i > 1; i-- {
		if x % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	count, recent := 0, 0
	for i := 1; count < 100000; i++ {
		if isPrime(i) {
			count++
			recent = i
		}
	}

	fmt.Println("100,000th prime number:", recent)
}
*/

// version 2



func isPrime(primes []int, x int) bool {
	if func() bool {
		for _, v := range primes {
			if x % v == 0 {
				return false
			}
		}
		return true
	} {
		primes = append(primes, x)
		return true
	} else {
		return false
	}
}

func main() {
	primes := make([]int, 5)
	count, recent := 0

	for i := 1; count < 100000; i++ {
		if isPrime(primes, i) {
			count++
			recent = i
		}
	}

	fmt.Println("100,000th prime number:", recent)
}
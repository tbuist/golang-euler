// find largest prime factor
package main

import (
	"fmt"
	"math"
)

func isPrime(x int64) bool {
	sqrt := int64(math.Sqrt(float64(x)))
	for i := int64(2); i <= sqrt; i++ {
		if x % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int64
	fmt.Print("Find largest prime: ")
	fmt.Scanln(&num)

	for i := int64(math.Sqrt(float64(num))); i > 1; i-- {
		if num % i == 0 && isPrime(i) {
			fmt.Println("Largest prime =", i)
			return
		}
	}
}
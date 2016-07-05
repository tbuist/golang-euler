// sum square difference
package main

import (
	"fmt"
)

func sumOfSquares(x int) int {
	sum := 0
	for i := 1; i <= x; i++ {
		sum += i * i
	}
	return sum
}

func squareOfSums(x int) int {
	sum := 0
	for i := 1; i <= x; i++ {
		sum += i
	}
	return sum * sum
}

func main() {

	sum := sumOfSquares(100)
	square := squareOfSums(100)

	fmt.Println("Sum square difference of first 100 numbers: ", sum, square, square - sum)
}
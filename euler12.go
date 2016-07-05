// highly divisible triangular number
package main

import (
	"fmt"
	"math"
)

func numDivs(x int) int {
	if x == 1 {
		return 1
	}
	num := 0
	max := int(math.Sqrt(float64(x)))
	for i := 1; i <= max; i++ {
		if x % i == 0 {
			num += 2
		}
	}
	if max * max == x {
		num--
	}

	return num
}

func main() {
	curTriNum := 1
	curTriIdx := 1

	for ; numDivs(curTriNum) < 501; curTriNum += curTriIdx {
		curTriIdx++
	}

	fmt.Println("First tri num with > 500 divisors =", curTriNum)
}
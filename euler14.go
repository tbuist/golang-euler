// longest collatz sequence
package main

import (
	"fmt"
)

func numTerms(x int) int {
	count := 1
	for ; x != 1 ; {
		if x % 2 == 0 {
			x /= 2
		} else {
			x = 3*x + 1
		}
		count++
	}

	return count
}

func main() {
	longChain := 0
	longNum := 0

	for i := 2; i < 1000000; i++ {
		tmp := numTerms(i)
		if tmp > longChain {
			longChain = tmp
			longNum = i
		}
	}

	fmt.Println("starting with 13:", numTerms(13))

	fmt.Println("Longest chain:", longChain)
	fmt.Println("Starting num:", longNum)

}
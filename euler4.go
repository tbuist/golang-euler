// largest palindrome product
package main

import (
	"fmt"
	"strconv"
)

func isPal(x int) bool {
	num := strconv.Itoa(x)
	left, right := 0, len(num) - 1
	for ; left < right; {
		if num[left] != num[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	largest := 1
	prod1 := 1
	for i := 999; i > 99; i-- {
		for j:= 999; j > 99; j-- {
			if i * j > largest && isPal(i * j) {
				largest = i * j
				prod1 = i
			}
		}
	}

	fmt.Println("Largest pal prod: ", prod1, largest / prod1, largest)
}
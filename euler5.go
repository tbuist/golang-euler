// smallest multiple
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Smallest number divisible by numbers 1 through 20\n\n")

	for i := 11; ; i++ {
		j := 2
		for ; j < 21; j++ {
			if i % j != 0 {
				break
			}
		}
		if j >= 21 {
			fmt.Println("Smallest number =", i)
			return
		}
	}
}
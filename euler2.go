package main

import (
	"fmt"
)

func main() {
	num1, num2, sum := 1, 2, 0
	for ; num2 < 4000000; {
		if num2 % 2 == 0 {
			sum += num2
		}
		tmp := num2
		num2 = num1 + num2
		num1 = tmp
	}
	fmt.Println(sum)
}
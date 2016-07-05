// power digit sum
package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	num := math.Pow(2, 1000)

	str := strconv.FormatFloat(num, 'f', 'G', 64)
	fmt.Println(str)


	sum := 0

	for i := 0; i < len(str); i++ {
		
		tmp, _ := strconv.Atoi(string(str[i]))
		sum += tmp
	}

	fmt.Println("Sum of digits:", sum)
}
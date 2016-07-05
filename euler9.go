package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			for c := 1; c < 1000; c++ {
				if a+b+c == 1000 && math.Pow(float64(a),2) + math.Pow(float64(b),2) == math.Pow(float64(c),2) {
					fmt.Println("Combo is", a,b,c, "with product", a*b*c)
					return
				}
			}
		}
	}
}